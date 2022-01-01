// Package model
//
// @author: xwc1125
package model

import (
	"fmt"
	"github.com/chain5j/chain5j-pkg/types"
)

type Storage map[types.Hash]types.Hash

func (s Storage) String() (str string) {
	for key, value := range s {
		str += fmt.Sprintf("%X : %X\n", key, value)
	}

	return
}

func (s Storage) Copy() Storage {
	cpy := make(Storage)
	for key, value := range s {
		cpy[key] = value
	}

	return cpy
}
