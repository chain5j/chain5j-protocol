// Package protocol
//
// @author: xwc1125
package protocol

import (
	"crypto"
	"github.com/chain5j/chain5j-pkg/crypto/signature"
	"github.com/chain5j/chain5j-protocol/models"
	"hash"
)

// KeyType 密钥类型
type KeyType int32

// NodeKey nodeKey接口
type NodeKey interface {
	ID() (models.NodeID, error)                                                     // 获取节点的ID
	IdFromPub(pub crypto.PublicKey) (models.NodeID, error)                          // 通过公钥获取ID
	PubKey(pubKey crypto.PublicKey) (PubKey, error)                                 // 原生公钥转公钥
	Sign(data []byte) (*signature.SignResult, error)                                // 使用节点私钥签名数据
	Verify(data []byte, signResult *signature.SignResult) (bool, error)             // 验证签名，sig是data对应Hash的签名
	RecoverId(data []byte, signResult *signature.SignResult) (models.NodeID, error) // 从签名中获取地址
	RecoverPub(data []byte, signResult *signature.SignResult) (PubKey, error)       // 从签名中公钥
}

// Key 加密密钥
type Key interface {
	Marshal() ([]byte, error)     // key序列化
	Unmarshal(input []byte) error // key反序列化
	Equals(Key) bool              // 判断key是否一致
	Raw() ([]byte, error)         // 返回x509格式
	Type() KeyType                // 密钥类型
	Hash() func() hash.Hash       // hash算法
}

// PrivKey 私钥密钥串
type PrivKey interface {
	Key
	Sign(data []byte) (*signature.SignResult, error) // 对bytes进行签名
	GetPublic() PubKey                               // 获取私钥对应的公钥
}

// PubKey 公钥串，用于验证签名
type PubKey interface {
	Key
	Verify(data []byte, signResult *signature.SignResult) (bool, error) // 验证签名，sig是data对应Hash的签名
}
