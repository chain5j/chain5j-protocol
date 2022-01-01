// Package dispatch
//
// @author: xwc1125
package dispatch

import (
	"context"
	"fmt"
	"runtime/debug"
	"sort"
	"sync"

	"github.com/chain5j/chain5j-pkg/codec"
	"github.com/chain5j/chain5j-pkg/event"
	"github.com/chain5j/chain5j-pkg/types"
	"github.com/chain5j/chain5j-pkg/util/hexutil"
	"github.com/chain5j/chain5j-protocol/models"
	"github.com/chain5j/chain5j-protocol/models/eventtype"
	"github.com/chain5j/chain5j-protocol/protocol"
	"github.com/chain5j/logger"
)

var (
	_ protocol.TxPools = new(TxPools)
)

type TxPoolsOption func(f *TxPools) error

type TxPools struct {
	log    logger.Logger
	ctx    context.Context
	cancel context.CancelFunc

	config      protocol.Config      // 配置
	broadcaster protocol.Broadcaster // 广播
	blockReader protocol.BlockReader // 数据库读
	apps        protocol.Apps        // apps
	appContexts protocol.AppContexts // context

	// appContextLock sync.RWMutex         // context lock
	txPools *sync.Map // types.TxType-->protocol.TxPool

	txsEvent event.Feed              // subscribe
	scope    event.SubscriptionScope // Scope
}

func NewTxPools(rootCtx context.Context, opts ...TxPoolsOption) (protocol.TxPools, error) {
	ctx, cancel := context.WithCancel(rootCtx)
	p := &TxPools{
		log:     logger.New("txPools"),
		ctx:     ctx,
		cancel:  cancel,
		txPools: new(sync.Map),
	}
	if err := p.apply(opts...); err != nil {
		p.log.Error("apply is error", "err", err)
		return nil, err
	}
	return p, nil
}

func (p *TxPools) apply(opts ...TxPoolsOption) error {
	for _, opt := range opts {
		if opt == nil {
			continue
		}
		if err := opt(p); err != nil {
			return err
		}
	}
	return nil
}
func WithConfig(config protocol.Config) TxPoolsOption {
	return func(f *TxPools) error {
		f.config = config
		return nil
	}
}
func WithBroadcaster(broadcaster protocol.Broadcaster) TxPoolsOption {
	return func(f *TxPools) error {
		f.broadcaster = broadcaster
		return nil
	}
}
func WithBlockReader(blockReader protocol.BlockReader) TxPoolsOption {
	return func(f *TxPools) error {
		f.blockReader = blockReader
		return nil
	}
}
func WithApps(apps protocol.Apps) TxPoolsOption {
	return func(f *TxPools) error {
		f.apps = apps
		return nil
	}
}

func (p *TxPools) Register(txType types.TxType, txPool protocol.TxPool) {
	p.txPools.Store(txType, txPool)
}

func (p *TxPools) TxPool(txType types.TxType) (protocol.TxPool, error) {
	txPool, ok := p.txPools.Load(txType)
	if ok {
		return txPool.(protocol.TxPool), nil
	}
	// 返回默认的txPool
	app, err := p.apps.App(txType)
	if err != nil {
		return nil, err
	}
	if pool, err := app.TxPool(p.config, p.apps, p.blockReader, p.broadcaster); err != nil {
		return nil, err
	} else if pool != nil {
		p.txPools.Store(txType, pool)
		return pool, nil
	} else {
		appContext := p.appContexts.Ctx(txType)

		defaultTxPool, err := NewDefaultTxPool(txType, p.config, app, appContext)
		if err != nil {
			return nil, err
		}
		p.txPools.Store(txType, defaultTxPool)
		return defaultTxPool, nil
	}
}

func (p *TxPools) Start() (err error) {
	// 根据当前的block stateRoot获取ctx
	p.log.Debug("block stateRoot", "state_roots", hexutil.Encode(p.blockReader.CurrentBlock().StateRoots()))
	appContexts, err := p.apps.NewAppContexts("txPool", p.blockReader.CurrentBlock().StateRoots())
	if err != nil {
		p.log.Crit("new appContexts error", "error", err)
		return err
	}
	p.appContexts = appContexts

	p.txPools.Range(func(key, value interface{}) bool {
		txPool := value.(protocol.TxPool)
		err = txPool.Start()
		if err != nil {
			return false
		}
		return true
	})
	return
}

func (p *TxPools) listen() {
	// 订阅新peer加入
	newPeerCh := make(chan models.P2PID)
	newPeerSub := p.broadcaster.SubscribeNewPeer(newPeerCh)
	// 订阅p2p交易消息
	p2pTxsMsg := make(chan *models.P2PMessage, 100)
	p2pTxsMsgSub := p.broadcaster.SubscribeMsg(models.TransactionSend, p2pTxsMsg)
	// 订阅区块
	chainHeadEventCh := make(chan eventtype.ChainHeadEvent)
	chainHeadEventSub := p.blockReader.SubscribeChainHeadEvent(chainHeadEventCh)

	for {
		select {
		case id := <-newPeerCh:
			if p.isMetrics(2) {
				p.log.Debug("new peer connect", "peer", id)
			}
			txs := p.GetTxs(10000)
			if len(txs) > 0 {
				for _, transactions := range txs {
					p.broadcaster.BroadcastTxs(&id, transactions, true)
				}
			}
			if p.isMetrics(2) {
				p.log.Info("new peer connect BroadcastTxs")
			}
		case err := <-newPeerSub.Err():
			p.log.Error("newPeerSub.Err", "err", err)
		case msg := <-p2pTxsMsg:
			// 接收到p2p的消息
			var txs models.TransactionSortedList
			err := codec.Coder().Decode(msg.Data, &txs)
			if err != nil {
				p.log.Error("decode p2p txs msg err", "err", err)
				break
			}
			if p.isMetrics(3) {
				p.log.Debug("receive txs from p2p", "hashes", txs.Hashes())
			}
			sort.Sort(txs)
			for _, tx := range txs {
				if err := p.Add(&msg.Peer, tx); err != nil {
					if p.isMetrics(3) {
						p.log.Trace("add p2p txs err", "err", err)
					}
				}
			}
		case err := <-p2pTxsMsgSub.Err():
			p.log.Error("p2p txsMsg sub err", "err", err)
			break
		case ch := <-chainHeadEventCh:
			// 接收到区块入库的事件后，需要修改AppContexts
			// var err error
			// p.appContextLock.Lock()
			// if p.appContexts, err = p.apps.NewAppContexts("txPool", p.blockReader.CurrentBlock().StateRoots()); err != nil {
			//	p.log.Crit("chan to new appContexts error", "error", err)
			// }
			// p.appContextLock.Unlock()
			// 删除区块中已经存在的交易
			for _, txList := range ch.Block.Transactions().Data() {
				if txList.Len() > 0 {
					for _, tx := range txList {
						p.DeleteOne(tx, true)
					}
				}
			}
			break
		case err := <-chainHeadEventSub.Err():
			if err != nil {
				p.log.Error("chain event err", "err", err)
			}
			return
		case <-p.ctx.Done():
			// 停止服务
			return
		}
	}
}

func (p *TxPools) Stop() (err error) {
	p.cancel()
	p.scope.Close()

	p.txPools.Range(func(key, value interface{}) bool {
		txPool := value.(protocol.TxPool)
		err = txPool.Stop()
		if err != nil {
			return false
		}
		return true
	})
	return
}

func (p *TxPools) Add(peerId *models.P2PID, tx models.Transaction) (err error) {
	defer func() {
		if r := recover(); r != nil {
			p.log.Error("txPools Add recover", "err", r)
			err = fmt.Errorf("%v", r)
			debug.PrintStack()
		}
	}()
	txType := tx.TxType()
	value, ok := p.txPools.Load(txType)
	if ok {
		if err := value.(protocol.TxPool).Add(peerId, tx); err != nil {
			return err
		}
	} else {
		// 使用默认的交易池
		app, err := p.apps.App(txType)
		if err != nil {
			return err
		}
		txPool, err := app.TxPool(p.config, p.apps, p.blockReader, p.broadcaster)
		if txPool == nil {
			appContext := p.appContexts.Ctx(txType)
			txPool, err = NewDefaultTxPool(txType, p.config, app, appContext)
			if err != nil {
				return err
			}
		}
		txPool.Add(peerId, tx)
		p.txPools.Store(txType, txPool)
	}
	// 通知订阅者
	txArray := []models.Transaction{tx}
	go p.txsEvent.Send(txArray)
	// 广播交易
	p.broadcaster.BroadcastTxs(peerId, txArray, false)

	return nil
}

func (p *TxPools) Get(txType types.TxType, hash types.Hash) (models.Transaction, models.TxStatus) {
	value, ok := p.txPools.Load(txType)
	if !ok {
		return nil, models.TxStatus_Unkown
	}
	return value.(protocol.TxPool).Get(hash)
}

func (p *TxPools) GetTxs(txsLimit uint64) map[types.TxType][]models.Transaction {
	txs := make(map[types.TxType][]models.Transaction)
	p.txPools.Range(func(txTypeObj, txListObj interface{}) bool {
		list := txListObj.(protocol.TxPool)
		txsSub := list.GetTxs(txsLimit)
		txs[txTypeObj.(types.TxType)] = txsSub
		return true
	})
	return txs
}

func (p *TxPools) FetchTxs(txsLimit uint64, headerTimestamp uint64) models.Transactions {
	txs := make(models.Transactions, 0)
	p.txPools.Range(func(txTypeObj, txListObj interface{}) bool {
		list := txListObj.(protocol.TxPool)
		txsSub := list.FetchTxs(txsLimit, headerTimestamp)
		if len(txsSub) > 0 {
			txs = append(txs, txsSub)
		}
		return true
	})
	sort.Sort(txs)
	return txs
}

func (p *TxPools) Fallback(txType types.TxType, txs []models.Transaction) error {
	txListObj, ok := p.txPools.Load(txType)
	if !ok {
		return errTxType
	}
	return txListObj.(protocol.TxPool).Fallback(txs)
}

func (p *TxPools) Delete(txType types.TxType, txs []models.Transaction, noErr bool) error {
	txListObj, ok := p.txPools.Load(txType)
	if !ok {
		return errTxType
	}
	for _, tx := range txs {
		return txListObj.(protocol.TxPool).Delete(tx, noErr)
	}
	return nil
}
func (p *TxPools) DeleteOne(tx models.Transaction, noErr bool) error {
	txListObj, ok := p.txPools.Load(tx.TxType())
	if !ok {
		return errTxType
	}
	return txListObj.(protocol.TxPool).Delete(tx, noErr)
}

func (p *TxPools) Len() map[types.TxType]uint64 {
	result := make(map[types.TxType]uint64, 0)
	p.txPools.Range(func(key, value interface{}) bool {
		result[key.(types.TxType)] = value.(protocol.TxPool).Len()
		return true
	})
	return result
}

func (p *TxPools) isMetrics(metrics uint64) bool {
	return p.config.TxPoolConfig().IsMetrics(metrics)
}

// func (p *TxPools) setAppContexts() error {
//	appContexts, err := p.apps.NewAppContexts("txPool", p.blockReader.CurrentBlock().StateRoots())
//	if err != nil {
//		p.log.Crit("new appContexts error", "error", err)
//		return err
//	}
//	p.appContexts = appContexts
//	return nil
// }

// ===========订阅===========

// Subscribe 订阅
func (p *TxPools) Subscribe(ch chan []models.Transaction) event.Subscription {
	return p.scope.Track(p.txsEvent.Subscribe(ch))
}
