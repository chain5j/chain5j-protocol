// Package statetype
//
// @author: xwc1125
package statetype

import (
	"fmt"
	"github.com/chain5j/chain5j-pkg/codec"
	"github.com/chain5j/chain5j-pkg/crypto/hashalg/sha3"
	"github.com/chain5j/chain5j-pkg/types"
	"testing"
)

func TestRoot(t *testing.T) {
	root := NewRoots()
	root.Put("STATE", types.HexToHash("0xfd781323483264460439877771c4624e195c69e3092f1c8c9a2a0e4c218ae23f"))
	root.Put("UTXO", types.HexToHash("0x11781323483264460439877771c4624e195c69e3092f1c8c9a2a0e4c218ae23f"))
	fmt.Println(root)

	bytes, _ := codec.Coder().Encode(root)

	fmt.Println(bytes)

	toHash := sha3.Keccak256(bytes)
	fmt.Println(types.BytesToHash(toHash))
	fmt.Println(toHash)

	newRoot := NewRoots()
	codec.Coder().Decode(bytes, newRoot)
	fmt.Println(newRoot)
}
