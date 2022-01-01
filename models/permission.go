// Package models
//
// @author: xwc1125
package models

import (
	"github.com/chain5j/chain5j-pkg/types"
)

// HandshakeMsg 握手消息
type HandshakeMsg struct {
	Peer               P2PID      `json:"peer" rlp:"-"`         // 节点ID
	ProtocolVersion    uint32     `json:"protocol_version"`     // 协议版本
	NetworkId          uint64     `json:"network_id"`           // 网络ID
	CurrentBlockHash   types.Hash `json:"current_block_hash"`   // 当前区块hash
	CurrentBlockHeight uint64     `json:"current_block_height"` // 当前区块高度
	GenesisBlockHash   types.Hash `json:"genesis_block_hash"`   // 创世块hash
}
