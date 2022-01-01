// Package protocol
//
// @author: xwc1125
package protocol

import (
	"github.com/chain5j/chain5j-pkg/event"
	"github.com/chain5j/chain5j-pkg/types"
	"github.com/chain5j/chain5j-protocol/models"
)

// TxPools 交易池中心
type TxPools interface {
	Start() error // 启动
	Stop() error  // 停止

	Register(txType types.TxType, txPool TxPool)
	TxPool(txType types.TxType) (TxPool, error)

	Add(peerId *models.P2PID, tx models.Transaction) error                          // 添加交易
	Get(txType types.TxType, hash types.Hash) (models.Transaction, models.TxStatus) // 通过hash获取交易
	GetTxs(txsLimit uint64) map[types.TxType][]models.Transaction                   // 获取指定数量的交易
	FetchTxs(txsLimit uint64, headerTimestamp uint64) models.Transactions           // 获取需打包的交易
	Fallback(txType types.TxType, txs []models.Transaction) error                   // 放回交易集
	Delete(txType types.TxType, txs []models.Transaction, noErr bool) error         // 删除交易
	Len() map[types.TxType]uint64                                                   // 交易池的数量

	Subscribe(ch chan []models.Transaction) event.Subscription // 订阅
}

// TxPool 交易池接口
type TxPool interface {
	Start() error // 启动
	Stop() error  // 停止

	Add(peerId *models.P2PID, tx models.Transaction) error                 // 添加交易
	Exist(hash types.Hash) bool                                            // 判断hash是否存在
	Get(hash types.Hash) (models.Transaction, models.TxStatus)             // 通过hash获取交易
	GetTxs(txsLimit uint64) []models.Transaction                           // 获取指定数量的交易
	FetchTxs(txsLimit uint64, headerTimestamp uint64) []models.Transaction // 获取需打包的交易
	Fallback(txs []models.Transaction) error                               // 放回交易集
	Delete(tx models.Transaction, noErr bool) error                        // 删除交易
	Len() uint64                                                           // 交易池的数量

	// Subscribe(ch chan models.Transactions) event.Subscription // 订阅
}
