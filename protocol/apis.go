// Package protocol
//
// @author: xwc1125
package protocol

import (
	"context"

	"github.com/chain5j/chain5j-pkg/types"
	"github.com/chain5j/chain5j-pkg/util/hexutil"
	"github.com/chain5j/chain5j-protocol/models"
	"github.com/chain5j/chain5j-protocol/models/statetype"
)

// API 对外的API
type API struct {
	Namespace string      // 对外的命名空间
	Version   string      // api版本
	Service   interface{} // 方法集合
	Public    bool        // 可供公众使用
}

// APIs 对外提供链服务
type APIs interface {
	// Syncing 同步状态
	Syncing(ctx context.Context) (*models.SyncingStatus, error)

	// SendRawTransaction 发送交易
	SendRawTransaction(ctx context.Context, txType types.TxType, rawTx hexutil.Bytes) (types.Hash, error)
	// GetTransaction 根据Hash获取交易
	GetTransaction(ctx context.Context, hash types.Hash) models.Transaction
	// GetTransactionReceipt 根据Hash获取交易收据
	GetTransactionReceipt(ctx context.Context, hash types.Hash) (models.Transaction, error)
	// GetTransactionLogs 根据Hash获取交易Logs
	GetTransactionLogs(ctx context.Context, hash types.Hash) ([]*statetype.Log, error)

	// BlockHeight 获取最新的区块高度
	BlockHeight(ctx context.Context) (*hexutil.Uint64, error)
	// GetBlockByHash 根据区块Hash获取区块信息
	GetBlockByHash(ctx context.Context, blockHash types.Hash) (*models.Block, error)
	// GetBlockByHeight 根据区块高度获取区块信息
	GetBlockByHeight(ctx context.Context, blockHeight hexutil.Uint64) (*models.Block, error)
	// GetBlockTransactionCountByHash 根据区块Hash获取区块交易个数
	GetBlockTransactionCountByHash(ctx context.Context, blockHash types.Hash) (*hexutil.Uint64, error)
	// GetBlockTransactionCountByHeight 根据区块高度获取区块交易个数
	GetBlockTransactionCountByHeight(ctx context.Context, blockHeight hexutil.Uint64) (*hexutil.Uint64, error)
	// GetTransactionByBlockHashAndIndex 查询指定块内具有指定索引序号的交易
	GetTransactionByBlockHashAndIndex(ctx context.Context, blockHash types.Hash, txIndex hexutil.Uint64) (models.Transaction, error)
	// GetTransactionByBlockHeightAndIndex 查询指定块内具有指定索引序号的交易
	GetTransactionByBlockHeightAndIndex(ctx context.Context, blockHeight hexutil.Uint64, txIndex hexutil.Uint64) (models.Transaction, error)

	// GetCode 回去合约对应的代码
	GetCode(ctx context.Context, contract types.Address, blockHeight hexutil.Uint64) (*hexutil.Bytes, error)
	// Call 调用合约，无需在区块链上创建交易
	Call(ctx context.Context, hash models.VmMessage) (*hexutil.Bytes, error)
	// EstimateGas 估算交易的gas
	EstimateGas(ctx context.Context, transaction models.Transaction) (*hexutil.Uint64, error)
	// CompileContract 编译合约
	CompileContract(ctx context.Context, compileType CompileType, contract hexutil.Bytes) (*hexutil.Bytes, error)

	// FilterSubscribeHeaders 过滤订阅header
	FilterSubscribeHeaders(ctx context.Context) (*models.Header, error)

	// APIs api集合
	APIs() []API
	RegisterAPI(apis []API)
}
