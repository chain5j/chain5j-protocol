// Package protocol
//
// @author: xwc1125
package protocol

import (
	"github.com/chain5j/chain5j-pkg/event"
	"github.com/chain5j/chain5j-pkg/types"
	"github.com/chain5j/chain5j-protocol/models"
	"github.com/chain5j/chain5j-protocol/models/eventtype"
)

// BlockReadWriter 区块读写接口
type BlockReadWriter interface {
	BlockWriter
	BlockReader
}

// BlockWriter 区块写接口
type BlockWriter interface {
	// InsertBlock 写入区块, 不包含状态处理, 不包含校验
	InsertBlock(block *models.Block, propagate bool) (err error)
	// ProcessBlock 区块处理，校验，状态处理(同步或接收到广播的区块信息)
	ProcessBlock(block *models.Block, propagate bool) (err error)
}

// BlockReader 区块读接口
type BlockReader interface {
	// Start 启动
	Start() error
	// Stop 停止
	Stop() error
	// IsRunning 判断blockChain是否正在执行
	IsRunning() bool

	// CurrentHeader 获取当前的header
	CurrentHeader() *models.Header
	// GetHeader 根据hash及区块高度获取区块头
	GetHeader(hash types.Hash, number uint64) *models.Header
	// GetHeaderByHash 根据区块hash获取区块头
	GetHeaderByHash(hash types.Hash) *models.Header
	// GetHeaderByNumber 根据区块高度获取区块头
	GetHeaderByNumber(number uint64) *models.Header
	// HasHeader 根据hash及区块高度判断区块是否存在
	HasHeader(hash types.Hash, number uint64) bool

	// CurrentBlock 获取当前区块
	CurrentBlock() *models.Block
	// GetBlock 根据hash及区块高度获取区块
	GetBlock(hash types.Hash, number uint64) *models.Block
	// GetBlockByHash 根据hash获取区块
	GetBlockByHash(hash types.Hash) *models.Block
	// GetBlockByNumber 根据区块高度获取区块
	GetBlockByNumber(number uint64) *models.Block
	// HasBlock 根据hash及区块高度判断区块是否存在
	HasBlock(hash types.Hash, number uint64) bool

	// GetBlockHashesFromHash 从指定hash开始获取一系列区块hash, 降序
	GetBlockHashesFromHash(hash types.Hash, max uint64) []types.Hash
	// GetBody 根据hash获取区块的交易内容
	GetBody(hash types.Hash) *models.Body
	// ValidateBody 验证区块的body
	ValidateBody(block *models.Block) error

	// GetAncestor 根据hash及区块高度，祖先高度获取区块hash及高度
	GetAncestor(hash types.Hash, number, ancestor uint64, maxNonCanonical *uint64) (types.Hash, uint64)
	// SubscribeChainHeadEvent 链订阅
	SubscribeChainHeadEvent(ch chan<- eventtype.ChainHeadEvent) event.Subscription
}
