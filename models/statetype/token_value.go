// Package statetype
//
// @author: xwc1125
package statetype

import (
	"github.com/chain5j/chain5j-pkg/types"
	"math/big"
)

type TokenValue struct {
	TokenAddress types.DomainAddress `json:"token_address"`
	Value        *big.Int            `json:"value"`
}

type TokenValues []TokenValue
