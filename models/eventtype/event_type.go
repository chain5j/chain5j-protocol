// Package eventtype
//
// @author: xwc1125
package eventtype

import (
	"github.com/chain5j/chain5j-pkg/types"
	"github.com/chain5j/chain5j-protocol/models"
)

// ChainEvent 区块事件
type ChainEvent struct {
	Block *models.Block
	Hash  types.Hash
}

// ChainSideEvent 侧链事件
type ChainSideEvent struct{ Block *models.Block }

// ChainHeadEvent 链head事件
type ChainHeadEvent struct{ Block *models.Block }

// BlockBroadcastEvent 区块广播事件
type BlockBroadcastEvent struct{}

// ExecBlockEvent 执行区块的事件
type ExecBlockEvent struct {
	Block *models.Block
}

// ExecFinishEvent 执行结束事件
type ExecFinishEvent struct {
	Res bool // 如果执行成功，返回true
}

// BlockReadyEvent 区块池收到后进行入库处理
type BlockReadyEvent struct {
	Block *models.Block
}

// ProposeBlockEvent 提议区块事件
type ProposeBlockEvent struct {
	Block *models.Block
}

// CommitBlockEvent 提交区块的事件
type CommitBlockEvent struct {
	Block *models.Block
}

// CommitCompleteEvent 提交完成的事件
type CommitCompleteEvent struct {
	Block *models.Block
}

// ErrOccurEvent 执行中出现错误时的事件
type ErrOccurEvent struct {
	Err error
}

// NewTxEvent 新交易事件
type NewTxEvent struct {
	Tx models.Transaction
}

// NewTxsEvent 新交易集事件
type NewTxsEvent struct {
	Txs models.Transactions
}

// ExecPendingTxEvent 执行pending交易事件
type ExecPendingTxEvent struct {
	Txs models.Transactions
}

// TxBroadcastEvent 交易广播事件
type TxBroadcastEvent struct{}

// RollbackEvent 回滚事件
type RollbackEvent struct{}
