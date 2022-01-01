// Package models
//
// @author: xwc1125
package models

import (
	"testing"

	"github.com/chain5j/chain5j-pkg/codec"
	crypto2 "github.com/chain5j/chain5j-pkg/crypto/signature"
	"github.com/chain5j/chain5j-pkg/types"
	"github.com/chain5j/chain5j-pkg/util/hexutil"
)

func TestBlockRlP(t *testing.T) {
	h := Header{
		ParentHash:  types.HexToHash("0xceb7cc96b472895cb3f8f608c3455c407ebcf58245e845d1b090549cd6620e3d"),
		Height:      100,
		StateRoots:  types.HexToHash("0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347").Bytes(),
		TxsRoot:     []types.Hash{types.HexToHash("0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347")},
		Timestamp:   1000000,
		GasUsed:     210000,
		GasLimit:    3000000,
		ArchiveHash: nil,
		Consensus: &Consensus{
			Name:      "pbft",
			Consensus: []byte{0x01, 0x02, 0x03},
		},
		Extra: []byte{0x04, 0x05, 0x06},
		Signature: &crypto2.SignResult{
			Name:      "P-256",
			PubKey:    []byte{0x04, 0x05, 0x06},
			Signature: []byte{0x04, 0x05, 0x06},
		},
	}

	block := NewBlock(&h, nil, nil)
	enc, err := codec.Coder().Encode(block)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("block rlp bytes: %s", hexutil.Encode(enc))
}
