// Package models
//
// @author: xwc1125
package models

import (
	"github.com/chain5j/chain5j-pkg/codec/json"
	"github.com/chain5j/chain5j-pkg/util/hexutil"
)

// Consensus 共识信息
type Consensus struct {
	Name      string `json:"name,omitempty"`                 // 共识名称
	Consensus []byte `json:"consensus,omitempty" rlp:"tail"` // 共识内容
}

// Copy 拷贝
func (c *Consensus) Copy() *Consensus {
	cpy := &Consensus{
		Name:      c.Name,
		Consensus: nil,
	}
	cpy.Consensus = make([]byte, len(c.Consensus))
	copy(cpy.Consensus, c.Consensus)
	return cpy
}
func (c *Consensus) MarshalJSON() ([]byte, error) {
	cJson := &extConsensus{
		Name:      c.Name,
		Consensus: c.Consensus,
	}
	return json.Marshal(cJson)
}
func (c *Consensus) UnmarshalJSON(input []byte) error {
	var cJson extConsensus
	err := json.Unmarshal(input, &cJson)
	if err != nil {
		return err
	}
	c.Name = cJson.Name
	c.Consensus = cJson.Consensus
	return nil
}

// extConsensus 共识对应的external对象
type extConsensus struct {
	Name      string        `json:"name,omitempty"`      // 共识名称
	Consensus hexutil.Bytes `json:"consensus,omitempty"` // 共识内容
}

type ConsensusHandler struct {
	Method  string
	Message []byte
	Block   *Block
}
