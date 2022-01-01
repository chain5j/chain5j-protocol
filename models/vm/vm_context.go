// Package vm
//
// @author: xwc1125
package vm

import (
	"math/big"

	"github.com/chain5j/chain5j-pkg/types"
	"github.com/chain5j/chain5j-protocol/models"
	"github.com/chain5j/chain5j-protocol/protocol"
)

type (
	// CanTransferFunc is the signature of a transfer guard function
	CanTransferFunc func(protocol.StateDB, types.Address, *big.Int) bool
	// TransferFunc is the signature of a transfer function
	TransferFunc func(protocol.StateDB, types.Address, types.Address, *big.Int)
	// GetHashFunc returns the nth block hash in the blockchain
	GetHashFunc func(uint64) types.Hash
)

// Context vm的上下文
type Context struct {
	CanTransfer CanTransferFunc // 帐户是否能够转账
	Transfer    TransferFunc    // 平台币转账
	GetHash     GetHashFunc     // 获取hash

	// Message内容
	Origin   types.Address
	GasPrice *big.Int

	// 区块内容
	Coinbase    types.Address
	GasLimit    uint64
	BlockNumber *big.Int
	Time        *big.Int
	Difficulty  *big.Int
	BlockHeader *models.Header
}
