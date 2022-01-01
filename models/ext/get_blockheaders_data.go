// Package ext external data
//
// @author: xwc1125
package ext

import (
	"fmt"
	"io"

	"github.com/chain5j/chain5j-pkg/codec/rlp"
	"github.com/chain5j/chain5j-pkg/types"
)

type GetBlockHeadersData struct {
	Origin  HashOrNumber // 区块检索头
	Amount  uint64       // 检索最大的个数
	Reverse bool         // 查询方向，false：往最新区块方向查询，true：往创世块方向查询
}

// ==============header=============

// HashOrNumber hash或数字对象
type HashOrNumber struct {
	Hash   types.Hash // 区块Hash
	Number uint64     // 区块高度
}

func (hn *HashOrNumber) EncodeRLP(w io.Writer) error {
	if hn.Hash == (types.Hash{}) {
		return rlp.Encode(w, hn.Number)
	}
	if hn.Number != 0 {
		return fmt.Errorf("both origin hash (%x) and number (%d) provided", hn.Hash, hn.Number)
	}
	return rlp.Encode(w, hn.Hash)
}
func (hn *HashOrNumber) DecodeRLP(s *rlp.Stream) error {
	_, size, _ := s.Kind()
	origin, err := s.Raw()
	if err == nil {
		switch {
		case size == 32:
			err = rlp.DecodeBytes(origin, &hn.Hash)
		case size <= 8:
			err = rlp.DecodeBytes(origin, &hn.Number)
		default:
			err = fmt.Errorf("invalid input size %d for origin", size)
		}
	}
	return err
}
