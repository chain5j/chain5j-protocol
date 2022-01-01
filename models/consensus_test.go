// Package models
//
// @author: xwc1125
package models

import (
	"github.com/chain5j/chain5j-pkg/codec"
	"github.com/chain5j/chain5j-pkg/util/hexutil"
	"testing"
)

func TestConsensusRLP(t *testing.T) {
	c := &Consensus{
		Name:      "pbft",
		Consensus: hexutil.MustDecode("0xf8a580c0f83f9492c8cae42a94045670cbb0bfcf8f790d9f8097e7949254e62fbca63769dfd4cc8e23f630f0785610ce94353c02434de6c99f5587b62ae9d6da2bd776daa7f85eae516d584a394a6d376479526251543869704157774470687647504364786f4246346b4b616173476e31556f624e67ae516d6353627843656756656f72716f69663666594b4c694d67627637744a6b4a46514e44646d464b32473346396ac0c0"),
	}
	bytes, err := codec.Coder().Encode(c)
	if err != nil {
		t.Fatal(err)
	}
	var c1 Consensus
	if err := codec.Coder().Decode(bytes, &c1); err != nil {
		t.Fatal(err)
	}

	chainConfig := &ChainConfig{
		ChainID:       1,
		ChainName:     "Chain5j",
		VersionName:   "v0.0.1",
		VersionCode:   1,
		GenesisHeight: 10,
		TxSizeLimit:   128,
		Packer: &PackerConfig{
			WorkerType:           0,
			BlockMaxTxsCapacity:  50000,
			BlockMaxSize:         2048,
			BlockMaxIntervalTime: 1000,
			BlockGasLimit:        8000000,
			Period:               0,
			EmptyBlocks:          0,
			Timeout:              10000,
			MatchTxsCapacity:     false,
		},
		//Consensus:    c,
		StateApp: &StateAppConfig{
			false,
		},
	}

	cBytes, err := codec.Coder().Encode(chainConfig)
	if err != nil {
		t.Fatal(err)
	}

	var cc1 ChainConfig
	if err := codec.Coder().Decode(cBytes, &cc1); err != nil {
		t.Fatal(err)
	}
	t.Logf("%v", cc1)
}
