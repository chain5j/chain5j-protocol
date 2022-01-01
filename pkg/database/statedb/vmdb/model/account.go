// Package model
//
// @author: xwc1125
package model

import (
	"bytes"
	"io"

	"github.com/chain5j/chain5j-pkg/codec/rlp"
	"github.com/chain5j/chain5j-pkg/types"
	"github.com/chain5j/chain5j-protocol/models/accounts"
	log "github.com/chain5j/logger"
)

type Account struct {
	accounts.AccountStore
	// Nonce    uint64
	// Balance  *big.Int
	Root     types.Hash
	CodeHash []byte

	// Tokens  map[types.DomainAddress]*big.Int
}

func NewAccount(addr types.DomainAddress) *Account {
	accountStore := accounts.NewAccountStore(addr.DomainAddr, "")
	return &Account{
		AccountStore: *accountStore,
		Root:         types.EmptyRootHash,
		CodeHash:     types.EmptyCode,
	}
}

func (a Account) IsEmpty() bool {
	if a.Balance == nil {
		return a.Nonce == 0 && bytes.Equal(a.CodeHash, types.EmptyCode)
	} else {
		return a.Nonce == 0 && a.Balance.Sign() == 0 && bytes.Equal(a.CodeHash, types.EmptyCode)
	}
}

type rlpAccount struct {
	CN       string     `json:"cn"`     // 用户名称 common name
	Domain   string     `json:"domain"` // 所在域
	Root     types.Hash `json:"root"`
	CodeHash []byte     `json:"code_hash"`
}

func (a *Account) EncodeRLP(w io.Writer) error {
	r := rlpAccount{
		CN:       a.CN,
		Domain:   a.Domain,
		Root:     a.Root,
		CodeHash: a.CodeHash,
	}
	return rlp.Encode(w, &r)
}

func (a *Account) DecodeRLP(s *rlp.Stream) error {
	var r rlpAccount
	err := s.Decode(&r)
	if err != nil {
		log.Error("DecodeTxsRLP", "err", err)
		return err
	}
	accountStore := accounts.NewAccountStore(r.CN, r.Domain)
	a.AccountStore = *accountStore
	a.CodeHash = r.CodeHash
	a.Root = r.Root
	return nil
}
