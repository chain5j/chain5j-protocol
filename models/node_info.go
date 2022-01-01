// Package models
//
// @author: xwc1125
package models

import (
	"github.com/chain5j/chain5j-pkg/codec/rlp"
	"io"
)

// NodeID 节点签名的地址(用于区块签名)
type NodeID string

func NodeIdFromString(base58Id string) (NodeID, error) {
	return NodeID(base58Id), nil
}
func (id NodeID) String() string {
	return string(id)
}
func (id *NodeID) EncodeRLP(w io.Writer) error {
	s := string(*id)
	if err := rlp.Encode(w, &s); err != nil {
		return err
	}
	return nil
}
func (id *NodeID) DecodeRLP(s *rlp.Stream) error {
	var dec string
	if err := s.Decode(&dec); err != nil {
		return err
	}

	*id = NodeID(dec)
	return nil
}
