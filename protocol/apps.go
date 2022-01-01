// Package protocol
//
// @author: xwc1125
package protocol

import (
	"github.com/chain5j/chain5j-pkg/types"
	"github.com/chain5j/chain5j-protocol/models"
)

// AppContext app的上下文
type AppContext interface {
	Caller() string   // 调用模块名称
	App() Application // 从ctx 获取App实例
}

// AppContexts AppContext集
type AppContexts interface {
	Ctx(t types.TxType) AppContext
}

// TxValidator 交易验证器
type TxValidator interface {
	// ValidateTx 添加transaction时进行校验
	ValidateTx(ctx AppContext, txI models.Transaction) error
	// ValidateTxSafe 验证交易，打包等检测
	ValidateTxSafe(ctx AppContext, txI models.Transaction, headerTimestamp uint64) error
}

// Application 应用接口
type Application interface {
	// Start 启动app
	Start() error
	// Stop 停止app
	Stop() error
	// NewAppContexts 创建一个新的appContext
	NewAppContexts(module string, args ...interface{}) (AppContext, error)

	// TxPool 每个app可自我维护一个txPool，如果不自我维护，那么将会使用默认的
	// 如果返回nil,nil-->使用默认的txPool
	// 如果返回nil,err-->有错误返回，会报错
	// 如果返回!nil,nil-->使用app自定义的txPool
	TxPool(config Config, apps Apps, blockReader BlockReader, broadcaster Broadcaster) (TxPool, error)

	TxValidator

	// DeleteErrTx 删除错误的交易
	DeleteErrTx(txI models.Transaction) error
	// DeleteOkTx 删除OK的交易
	DeleteOkTx(txI models.Transaction) error

	GetCacheNonce(ctx AppContext, account string) (uint64, error)

	// Prepare 计算出stateRoot
	Prepare(ctx AppContext, header *models.Header, txs models.TransactionSortedList, totalGas uint64) *models.TxsStatus

	// Commit 提交txs,返回root【commit】提交到数据库
	Commit(ctx AppContext, header *models.Header) error

	// 处理日志收据
	//Post(ctx AppContext, block *models.Block, txI models.Transaction, txIndex int) (receipt interface{}, gasUsed uint64, err error)
	// 数据的版本
	//Snapshot(ctx AppContext) (revisionId int)
	// 回滚到指定的版本
	//RevertToSnapshot(ctx AppContext, versionId int)
}

// Apps app接口
type Apps interface {
	// Register 注册application
	Register(txType types.TxType, app Application)
	// App 根据txType获取application
	App(txType types.TxType) (Application, error)
	// NewAppContexts 创建新的context
	NewAppContexts(module string, preRoot []byte) (AppContexts, error)
	// Prepare 预处理交易
	Prepare(ctx AppContexts, preRoot []byte, header *models.Header, txs models.Transactions, totalGas uint64) ([]byte, models.AppsStatus)
	// Commit 提交
	Commit(ctx AppContexts, header *models.Header) error
}
