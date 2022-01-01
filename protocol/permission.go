// Package protocol
//
// @author: xwc1125
package protocol

import (
	"github.com/chain5j/chain5j-pkg/event"
	"github.com/chain5j/chain5j-protocol/models"
	"github.com/chain5j/chain5j-protocol/models/permission"
)

// Handshake ...
type Handshake interface {
	Start() error
	Stop() error
	RequestHandshake(id models.P2PID) error // 对p2p进行握手协议处理
	SubscribeHandshake(msg chan *models.HandshakeMsg) event.Subscription
}

type Permission interface {
	IsAdmin(key string, height uint64) bool
	IsPeer(peerId models.P2PID, height uint64) bool

	AddSupervisor(key string, info permission.MemberInfo) error
	DelSupervisor(key string) error

	AddPermission(peerId models.P2PID, key string, info permission.MemberInfo, r permission.RoleType) error
	DelPermission(peerId models.P2PID, key string, r permission.RoleType) error
}
