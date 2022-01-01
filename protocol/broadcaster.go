// Package protocol
//
// @author: xwc1125
package protocol

import (
	"github.com/chain5j/chain5j-pkg/event"
	"github.com/chain5j/chain5j-protocol/models"
)

// Broadcaster 广播
type Broadcaster interface {
	Start() error // Start 启动
	Stop() error  // Stop 停止

	SubscribeMsg(msgType uint, ch chan<- *models.P2PMessage) event.Subscription // 订阅消息
	SubscribeNewPeer(newPeerCh chan models.P2PID) event.Subscription            // 订阅新接入
	SubscribeDropPeer(dropPeerCh chan models.P2PID) event.Subscription          // 订阅移除peer

	Broadcast(peers []models.P2PID, mType uint, payload []byte) error          // Broadcast 广播数据
	BroadcastTxs(peerId *models.P2PID, txs []models.Transaction, isForce bool) // BroadcastTxs 广播交易

	RegisterTrustPeer(peerID models.P2PID) error   // 注册trust peer
	DeregisterTrustPeer(peerID models.P2PID) error // 注销trust peer
}
