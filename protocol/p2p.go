// Package protocol
//
// @author: xwc1125
package protocol

import (
	"github.com/chain5j/chain5j-pkg/event"
	"github.com/chain5j/chain5j-protocol/models"
)

// P2PService P2P 服务
type P2PService interface {
	Start() error     // 启动p2p服务
	Stop() error      // 停止p2p服务
	Id() models.P2PID // 本节点peerId

	NetURL() string                       // 本节点网络标识
	RemotePeers() []models.P2PID          // 已经连接的节点列表
	P2PInfo() map[string]*models.P2PInfo  // 已经连接的节点URL
	HandshakeSuccess(peerId models.P2PID) //  握手成功

	Send(peerId models.P2PID, msg *models.P2PMessage) error // 发送p2p消息

	AddPeer(peerUrl string) error       // 添加节点
	DropPeer(peerId models.P2PID) error // 删除节点

	SubscribeMsg(msgType uint, ch chan<- *models.P2PMessage) event.Subscription // 订阅节点信息
	SubscribeHandshakePeer(ch chan<- models.P2PID) event.Subscription           // 新节点事件
	SubscribeNewPeer(ch chan<- models.P2PID) event.Subscription                 // 新节点事件
	SubscribeDropPeer(ch chan<- models.P2PID) event.Subscription                // 节点断开事件
}
