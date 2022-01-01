// Package models
//
// @author: xwc1125
package models

import (
	"io"

	"github.com/chain5j/chain5j-pkg/codec/rlp"
)

// P2PID p2p的ID(用于p2p)
type P2PID string

// IDFromString id为pretty
func IDFromString(base58Id string) (P2PID, error) {
	// TODO validate 检查格式
	return P2PID(base58Id), nil
}

// String 节点ID的string
func (id P2PID) String() string {
	return string(id)
}

func (id *P2PID) EncodeRLP(w io.Writer) error {
	s := string(*id)
	if err := rlp.Encode(w, &s); err != nil {
		return err
	}
	return nil
}
func (id *P2PID) DecodeRLP(s *rlp.Stream) error {
	var dec string
	if err := s.Decode(&dec); err != nil {
		return err
	}

	*id = P2PID(dec)
	return nil
}

// P2PInfo 节点信息
type P2PInfo struct {
	Id        P2PID  `json:"id"`        // 节点ID
	NetUrl    string `json:"net_url"`   // 网络url
	Connected bool   `json:"connected"` // 是否连接
}
