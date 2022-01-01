// Package dispatch
//
// @author: xwc1125
package dispatch

import (
	"context"
	"errors"
	"sync"
	"time"

	linkedHashMap "github.com/chain5j/chain5j-pkg/collection/maps/linked_hashmap"
	"github.com/chain5j/chain5j-pkg/pool/pool"
	"github.com/chain5j/chain5j-pkg/types"
	"github.com/chain5j/chain5j-pkg/util/dateutil"
	"github.com/chain5j/chain5j-protocol/models"
	"github.com/chain5j/chain5j-protocol/protocol"
	"github.com/chain5j/logger"
	"go.uber.org/atomic"
)

var (
	_ protocol.TxPool = new(txPool)
)

var (
	errTxType = errors.New("unsupported the txType")
)

type txPool struct {
	log      logger.Logger
	txSet    sync.Map      // tx集合,hash==>models.Transaction,缓存所有未入库的交易
	txSetLen atomic.Uint64 // tx集合长度

	queue       *linkedHashMap.LinkedHashMap // 等待的交易：hash==>hash
	pendingMap  map[types.Hash]types.Hash    // pending交易：hash==>hash
	pendingLock sync.RWMutex                 // 锁

	txType      types.TxType
	config      protocol.Config
	application protocol.Application

	appContext protocol.AppContext // context
	pool       *pool.Pool
}

func NewDefaultTxPool(txType types.TxType, config protocol.Config, app protocol.Application, appContext protocol.AppContext) (protocol.TxPool, error) {
	newPool, err := pool.NewPool(1000)
	if err != nil {
		return nil, err
	}

	return &txPool{
		log:        logger.New("txPool_" + txType.Value()),
		queue:      linkedHashMap.NewLinkedHashMap(),
		pendingMap: make(map[types.Hash]types.Hash, 0),

		txType:      txType,
		config:      config,
		application: app,
		appContext:  appContext,
		pool:        newPool,
	}, nil
}

func (tl *txPool) Start() error {
	return nil
}
func (tl *txPool) Stop() error {
	return nil
}

// Add 添加交易
func (tl *txPool) Add(peerId *models.P2PID, tx models.Transaction) error {
	if tx.TxType() != tl.txType {
		return errTxType
	}
	t1 := time.Now()
	hash := tx.Hash()
	if _, ok := tl.txSet.Load(hash); !ok {
		tl.txSet.Store(tx.Hash(), tx)
		tl.txSetLen.Add(1)

		tl.queue.Add(hash, hash)
	}
	if tl.metrics(2) {
		tl.log.Debug("txPool add end", "elapsed", dateutil.PrettyDuration(time.Since(t1)))
	}
	return nil
}

// Exist 判断key是否存在
func (tl *txPool) Exist(hash types.Hash) bool {
	_, ok := tl.txSet.Load(hash)
	return ok
}

// Get 通过hash获取交易
func (tl *txPool) Get(hash types.Hash) (models.Transaction, models.TxStatus) {
	txStatus := models.TxStatus_Unkown
	value, ok := tl.txSet.Load(hash)
	if !ok {
		return nil, txStatus
	}

	tl.pendingLock.RLock()
	_, ok = tl.pendingMap[hash]
	tl.pendingLock.RUnlock()
	if ok {
		txStatus = models.TxStatus_Pending
	}
	if !ok {
		if _, ok = tl.queue.Get(hash.String()); ok {
			txStatus = models.TxStatus_Waiting
		}
	}

	return value.(models.Transaction), txStatus
}

func (tl *txPool) GetTxs(txsLimit uint64) []models.Transaction {
	txs := make([]models.Transaction, 0, txsLimit)
	tl.txSet.Range(func(key, value interface{}) bool {
		txs = append(txs, value.(models.Transaction))
		return true
	})
	return txs
}

// GetFromQueue 通过hash获取queue中交易
func (tl *txPool) GetFromQueue(hash types.Hash) models.Transaction {
	value, ok := tl.txSet.Load(hash)
	if !ok {
		return nil
	}

	if _, ok = tl.queue.Get(hash); ok {
		return value.(models.Transaction)
	}
	return nil
}

// Len 获取交易池的大小
func (tl *txPool) Len() uint64 {
	return tl.txSetLen.Load()
}

// PendingLen pending的大小
func (tl *txPool) PendingLen() int {
	tl.pendingLock.RLock()
	defer tl.pendingLock.RUnlock()
	return len(tl.pendingMap)
}

// QueueLen queue的大小
func (tl *txPool) QueueLen() int {
	return tl.queue.Len()
}

// Delete 删除hashes
func (tl *txPool) Delete(tx models.Transaction, noErr bool) error {
	tl.pendingLock.Lock()
	hash := tx.Hash()
	delete(tl.pendingMap, hash)
	tl.pendingLock.Unlock()
	tl.queue.Remove(hash)
	tl.txSet.Delete(hash)
	return nil
}

// Fallback 交易放回
func (tl *txPool) Fallback(txs []models.Transaction) error {
	if len(txs) == 0 {
		return nil
	}
	for _, tx := range txs {
		hash := tx.Hash()
		// 从pending中删除
		tl.pendingLock.Lock()
		delete(tl.pendingMap, hash)
		tl.pendingLock.Unlock()

		if _, ok := tl.txSet.Load(hash); !ok {
			// 交易不存在，那么直接添加
			tl.Add(nil, tx)
		} else {
			// 交易已经存在
			if _, ok := tl.queue.Get(hash); !ok {
				tl.queue.AddFront(hash, hash)
			}
		}
	}
	return nil
}

func (tl *txPool) FallbackHashes(txHashes []types.Hash) error {
	if len(txHashes) == 0 {
		return nil
	}
	for _, hash := range txHashes {
		// 从pending中删除
		tl.pendingLock.Lock()
		delete(tl.pendingMap, hash)
		tl.pendingLock.Unlock()

		if _, ok := tl.txSet.Load(hash); ok {
			// 交易已经存在
			if _, ok := tl.queue.Get(hash); !ok {
				tl.queue.AddFront(hash, hash)
			}
		}
	}
	return nil
}

// FetchTxs 拉取交易
func (tl *txPool) FetchTxs(txsLimit uint64, headerTimestamp uint64) []models.Transaction {
	var (
		okTxs   = make([]models.Transaction, 0, txsLimit)
		tempTxs = make([]types.Hash, 0, txsLimit) // 缓存需要被拉走的交易，防止异常数据丢失
	)
	defer func() {
		if p := recover(); p != nil {
			tl.log.Error("fetchTxs recover", "err", p)
			if len(tempTxs) > 0 {
				tl.FallbackHashes(tempTxs)
			}
		}
	}()

	qLen := tl.queue.Len()
	if qLen > 0 {
		var (
			errHashes = make([]types.Hash, 0) // 错误的hash
			isCh      = true                  // 是否选择多协程
			wg        sync.WaitGroup
		)
		t1 := time.Now()
		if isCh {
			tempTxs, okTxs, errHashes = tl.asyncFetchTxs(txsLimit, headerTimestamp)
		} else {
			tempTxs, okTxs, errHashes = tl.syncFetchTxs(txsLimit, headerTimestamp)
		}
		if tl.metrics(1) {
			tl.log.Info("t-0) parse txs end", "txsLen", len(okTxs), "errTxsLen", len(errHashes), "elapsed", dateutil.PrettyDuration(time.Since(t1)))
		}

		// 删除错误
		t2 := time.Now()
		tl.delErrHashes(errHashes, &wg)
		tl.addPending(okTxs, &wg)
		wg.Wait()
		if tl.metrics(1) {
			tl.log.Info("t-1) pending add and queue delete", "txsLen", len(okTxs), "elapsed", dateutil.PrettyDuration(time.Since(t2)))
		}

		if tl.metrics(1) {
			tl.log.Info("txPool fetchTxs end", "txsLimit", txsLimit, "txsLen", len(okTxs), "tempTxsLen", len(tempTxs), "elapsed", dateutil.PrettyDuration(time.Since(t1)))
		}
	}

	// todo 如果库里都没有数据了，查看缓存是否还有数据
	if len(okTxs) == 0 && (tl.Len() > 0 || tl.PendingLen() > 0) {
		tl.log.Error("数据不一致", "txSetLen", tl.Len(), "queueLen", tl.queue.Len(), "pendingLen", tl.PendingLen())
	}

	return okTxs
}

func (tl *txPool) addPending(okTxs []models.Transaction, wg *sync.WaitGroup) {
	if len(okTxs) > 0 {
		wg.Add(1)
		go func(okTxs []models.Transaction) {
			defer wg.Done()
			tl.pendingLock.Lock()
			defer tl.pendingLock.Unlock()
			t3 := time.Now()
			for _, tx := range okTxs {
				tl.pendingMap[tx.Hash()] = tx.Hash()
				tl.queue.Remove(tx.Hash())
			}
			if tl.metrics(1) {
				tl.log.Info("t-1_3) pending add cacheKVs", "okTxsLen", len(okTxs), "elapsed", dateutil.PrettyDuration(time.Since(t3)))
			}
		}(okTxs)
	}
}
func (tl *txPool) delErrHashes(errHashes []types.Hash, wg *sync.WaitGroup) {
	if len(errHashes) > 0 {
		wg.Add(1)
		go func(errHashes []types.Hash) {
			defer wg.Done()
			t3 := time.Now()
			for _, hash := range errHashes {
				tl.queue.Remove(hash)
				tl.txSet.Delete(hash)
			}
			if tl.metrics(1) {
				tl.log.Info("t-1_2) rm errHashes", "elapsed", dateutil.PrettyDuration(time.Since(t3)), "errHashes", errHashes)
			}
		}(errHashes)
	}
}

func (tl *txPool) asyncFetchTxs(txsLimit uint64, headerTimestamp uint64) ([]types.Hash, []models.Transaction, []types.Hash) {
	// 需要打包的交易存在
	list := tl.queue.GetLinkList()
	qLen := list.Len()

	if int(txsLimit) > qLen {
		// 如果所需的交易个数小于队列的长度，那么只取最小的个数
		txsLimit = uint64(qLen)
	}
	var (
		tempTxs   = make([]types.Hash, 0, txsLimit)         // 缓存需要被拉走的交易，防止异常数据丢失
		okTxs     = make([]models.Transaction, 0, txsLimit) // 正确的hash
		errHashes = make([]types.Hash, 0, txsLimit)         // 错误的hash

		okTxCh      = make(chan models.Transaction, txsLimit)
		errTxCh     = make(chan types.Hash, txsLimit)
		ctx, cancel = context.WithCancel(context.Background())
		wg          sync.WaitGroup
	)

	// 循环取数据
	{
		wg.Add(1)
		go func(txsLimit uint64) {
			defer wg.Done()
			forTimes := 0
			for head := list.Front(); head != nil; head = head.Next() {
				hash := head.Value.(types.Hash)
				tx := tl.GetFromQueue(hash)
				if tx != nil {
					tempTxs = append(tempTxs, hash)

					wg.Add(1)
					tl.pool.Submit(tl.asyncVerifyTx(hash, tx, headerTimestamp, okTxCh, errTxCh))

					forTimes++
					if forTimes >= int(txsLimit) {
						if tl.metrics(2) {
							tl.log.Debug("verify tx", "txsLimit", txsLimit, "forTimes", forTimes)
						}
						break
					}
				} else {
					errHashes = append(errHashes, hash)
				}
			}
		}(txsLimit)
	}
	// 接收chan数据
	{
		go func() {
			if r := recover(); r != nil {
				tl.log.Error("receive chan recover", "err", r)
				cancel()
				if len(tempTxs) > 0 {
					tl.FallbackHashes(tempTxs)
				}
				return
			}
			for {
				select {
				case okTx := <-okTxCh:
					okTxs = append(okTxs, okTx)
					wg.Done()
				case errTx := <-errTxCh:
					errHashes = append(errHashes, errTx)
					wg.Done()
				case <-ctx.Done():
					return
				}
			}
		}()
	}
	wg.Wait()
	cancel()
	close(okTxCh)
	close(errTxCh)
	return tempTxs, okTxs, errHashes
}
func (tl *txPool) asyncVerifyTx(hash types.Hash, tx models.Transaction, headerTimestamp uint64, okTxCh chan models.Transaction, errTxCh chan types.Hash) func() {
	return func() {
		// 验证交易是否过期,数据库是否存在等
		err := tl.application.ValidateTxSafe(tl.appContext, tx, headerTimestamp)
		if err == nil {
			okTxCh <- tx
		} else {
			tl.log.Error("validate err", "hash", hash, "err", err)
			errTxCh <- hash
		}
	}
}
func (tl *txPool) syncFetchTxs(txsLimit uint64, headerTimestamp uint64) ([]types.Hash, []models.Transaction, []types.Hash) {
	// 需要打包的交易存在
	list := tl.queue.GetLinkList()
	qLen := list.Len()
	if int(txsLimit) > qLen {
		// 如果所需的交易个数小于队列的长度，那么只取最小的个数
		txsLimit = uint64(qLen)
	}
	var (
		tempTxs   = make([]types.Hash, 0, txsLimit)         // 缓存需要被拉走的交易，防止异常数据丢失
		okTxs     = make([]models.Transaction, 0, txsLimit) // 正确的hash
		errHashes = make([]types.Hash, 0, txsLimit)         // 错误的hash
	)
	forTimes := 0
	// list中的数据出来
	for head := list.Front(); head != nil; head = head.Next() {
		hash := head.Value.(types.Hash)
		tx := tl.GetFromQueue(hash)
		if tx != nil {
			tempTxs = append(tempTxs, hash)

			err := tl.application.ValidateTxSafe(tl.appContext, tx, headerTimestamp)
			if err != nil {
				// todo 需要删除
				errHashes = append(errHashes, hash)
			} else {
				okTxs = append(okTxs, tx)
			}

			forTimes++
			if forTimes >= int(txsLimit) {
				if tl.metrics(2) {
					tl.log.Debug("verify tx", "txsLimit", txsLimit, "forTimes", forTimes)
				}
				break
			}
		} else {
			errHashes = append(errHashes, hash)
		}
	}
	return tempTxs, okTxs, errHashes
}

func (tl *txPool) metrics(_metricsLevel uint64) bool {
	return tl.config.TxPoolConfig().Metrics && _metricsLevel <= tl.config.TxPoolConfig().MetricsLevel
}
