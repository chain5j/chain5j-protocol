// Package model
//
// @author: xwc1125
package model

import (
	"math/big"

	"github.com/chain5j/chain5j-pkg/types"
)

type Account struct {
	Nonce    uint64     `json:"nonce"`
	Balance  *big.Int   `json:"balance"`
	Root     types.Hash `json:"root"`
	CodeHash []byte     `json:"code_hash"`
}
