// Package crypto
//
// @author: xwc1125
package crypto

import (
	"github.com/chain5j/chain5j-pkg/codec/rlp"
	"github.com/chain5j/chain5j-pkg/crypto/hashalg/sha3"
	"github.com/chain5j/chain5j-pkg/types"
)

// CreateAddress 根据给定的地址及nonce生成新地址
func CreateAddress(b types.Address, nonce uint64) types.Address {
	data, _ := rlp.EncodeToBytes([]interface{}{b, nonce})
	return types.BytesToAddress(sha3.Keccak256(data)[12:])
}

// CreateAddress2 根据给定地址，salt和code生成新的地址
func CreateAddress2(b types.Address, salt [32]byte, code []byte) types.Address {
	return types.BytesToAddress(sha3.Keccak256([]byte{0xff}, b.Bytes(), salt[:], sha3.Keccak256(code))[12:])
}
