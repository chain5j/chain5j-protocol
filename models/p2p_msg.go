// Package models
//
// @author: xwc1125
package models

type P2PType = uint

const (
	TransactionSend  P2PType  = iota // 交易发送
	BlockMsg                         // 区块信息
	ConsensusMsg                     // 共识信息
	MsgTypeHandshake = 0x9999        // 握手消息
)

// P2PMessage P2P消息
type P2PMessage struct {
	Type uint   `json:"type"` // 类型
	Peer P2PID  `json:"peer"` // 节点ID
	Data []byte `json:"data"` // 数据
}

// MsgReader 消息读
type MsgReader interface {
	ReadMsg() (*P2PMessage, error)
}

// MsgWriter 消息写
type MsgWriter interface {
	// WriteMsg 发送消息是堵塞的，必须对方已消费。[消息只能发送一次]
	WriteMsg(msg *P2PMessage) error
}

// MsgReadWriter 消息读写
type MsgReadWriter interface {
	MsgReader
	MsgWriter
}
