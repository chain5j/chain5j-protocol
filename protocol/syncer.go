// Package protocol
//
// @author: xwc1125
package protocol

import (
	"github.com/chain5j/chain5j-pkg/types"
	"github.com/chain5j/chain5j-protocol/models"
	"github.com/chain5j/chain5j-protocol/models/ext"
)

// Syncer 同步
type Syncer interface {
	Start() error
	Stop() error
}

// RequestSync ...
type RequestSync interface {
	RequestOneHeader(peerId models.P2PID, hash types.Hash) error
	RequestHeadersByHash(peerId models.P2PID, origin types.Hash, amount int, skip int, reverse bool) error
	RequestHeadersByNumber(peerId models.P2PID, origin uint64, amount int, skip int, reverse bool) error
	HandleBlockHeadersMsg(peerId models.P2PID, headers []*models.Header) error
	HandleBlockBodiesMsg(peerId models.P2PID, request []*models.Body)
}

// ResponseSync ...
type ResponseSync interface {
	SendBlockHeaders(peerId models.P2PID, query ext.GetBlockHeadersData)
	SendBlockBodies(peerId models.P2PID, hashes []types.Hash)
}
