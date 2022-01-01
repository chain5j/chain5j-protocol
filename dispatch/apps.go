// Package dispatch
//
// @author: xwc1125
package dispatch

import (
	"errors"
	"sync"

	"github.com/chain5j/chain5j-pkg/codec"
	"github.com/chain5j/chain5j-pkg/types"
	"github.com/chain5j/chain5j-pkg/util/hexutil"
	"github.com/chain5j/chain5j-protocol/models"
	"github.com/chain5j/chain5j-protocol/models/statetype"
	"github.com/chain5j/chain5j-protocol/protocol"
	"github.com/chain5j/logger"
)

var (
	_ protocol.Apps = new(apps)
)

var (
	errNoApplication = errors.New("no application is exist")
)

type apps struct {
	log     logger.Logger
	nodeKey protocol.NodeKey
	apps    *sync.Map // types.TxType-->protocol.Application
}

func NewApps(nodeKey protocol.NodeKey) protocol.Apps {
	return &apps{
		log:     logger.New("apps"),
		apps:    new(sync.Map),
		nodeKey: nodeKey,
	}
}

// Register 注册回调
// 传入txType进行注册是为了可以实现一个application被多次使用
func (a *apps) Register(txType types.TxType, app protocol.Application) {
	// 如果map里没有就new一个新的
	if _, ok := a.apps.Load(txType); !ok {
		a.apps.Store(txType, app)
	}
}

// App 通过txType获取Application
func (a *apps) App(txType types.TxType) (protocol.Application, error) {
	app, ok := a.apps.Load(txType)
	if ok {
		return app.(protocol.Application), nil
	}
	return nil, errNoApplication
}

// NewAppContexts 根据module创建一个新的appContexts
func (a *apps) NewAppContexts(module string, preRoot []byte) (appCtxs protocol.AppContexts, err error) {
	ctx := &appContexts{
		appCtxs: new(sync.Map),
	}
	a.log.Debug("new app contexts by preRoot", "preRoot", hexutil.Encode(preRoot))

	roots := statetype.NewRoots()
	err = codec.Coder().Decode(preRoot, roots)
	if err != nil {
		a.log.Error("decode roots err", "err", err)
		return nil, err
	}
	// 循环获取appCtx
	var appCtx protocol.AppContext
	a.apps.Range(func(key, value interface{}) bool {
		t := key.(types.TxType)
		appCtx, err = value.(protocol.Application).NewAppContexts(module, roots.GetObj(t))
		if err != nil {
			a.log.Error("new app contexts err", "txType", t.Value(), "err", err)
			return false
		}
		ctx.appCtxs.Store(t, appCtx)
		return true
	})

	return ctx, nil
}

func (a *apps) Prepare(ctx protocol.AppContexts, preRoot []byte, header *models.Header, txs models.Transactions, totalGas uint64) ([]byte, models.AppsStatus) {
	txsLen := txs.Len()
	var resultTxStatus models.AppsStatus
	if txsLen == 0 {
		return preRoot, resultTxStatus
	}

	// 创建新的roots
	currentRoots := statetype.NewRoots()

	for _, txList := range txs.Data() {
		if txList.Len() > 0 {
			for _, tx := range txList {
				txType := tx.TxType()
				app, err := a.App(txType)
				if err != nil {
					a.log.Error("get application by txType err", "err", err)
					break
				}
				// app预处理交易
				txsStatus := app.Prepare(ctx.Ctx(txType), header, txList, totalGas)
				if txsStatus != nil {
					currentRoots.Put(txType, types.BytesToHash(txsStatus.StateRoots))
					resultTxStatus = append(resultTxStatus, *txsStatus)
				} else {
					return preRoot, nil
				}
			}
		}
	}

	// 最终的stateRoots
	rootBytes, _ := codec.Coder().Encode(currentRoots)
	return rootBytes, resultTxStatus
}

// Commit 提交
func (a *apps) Commit(ctx protocol.AppContexts, header *models.Header) error {
	// 判断roots是否正确
	roots := statetype.NewRoots()
	err := codec.Coder().Decode(header.StateRoots, roots)
	if err != nil {
		return err
	}

	a.log.Debug("apps commit", "root", roots)

	kvs := roots.Sort()
	for _, v := range kvs {
		app, err := a.App(v.Key)
		if err != nil {
			a.log.Error("get application by txType err", "err", err)
			return err
		}
		// app进行提交
		err = app.Commit(ctx.Ctx(v.Key), header)
		if err != nil {
			return err
		}
	}

	return nil
}
