// Package accounts
//
// @author: xwc1125
package accounts

import (
	"encoding/json"
	"math/big"
	"testing"

	"github.com/chain5j/chain5j-pkg/codec"
	"github.com/chain5j/chain5j-pkg/types"
)

func TestAccountStoreRLP(t *testing.T) {
	address := types.HexToAddress("0x92c8cae42a94045670cbb0bfcf8f790d9f8097e7")
	store := AccountStore{
		Nonce:     2,
		Balance:   big.NewInt(100),
		Addresses: make(map[types.Address]*AddressStore),
		CN:        "user1",
		Domain:    "chain5j.com",
		Permissions: &Permissions{
			EnableRegisterUser:   true,
			EnableRegisterDomain: true,
		},
		XXX: make(map[string][]byte),
	}
	store.Addresses[address] = &AddressStore{
		map[string][]byte{
			"111": []byte("111"),
		},
	}
	store.XXX["Key"] = []byte{0x1}

	rlpBytes, err := codec.Coder().Encode(&store)
	if err != nil {
		t.Fatal(err)
	}

	var dec AccountStore
	err = codec.Coder().Decode(rlpBytes, &dec)
	if err != nil {
		t.Fatal(err)
	}

	bytes, err := json.Marshal(dec)
	t.Logf("dec: %v\n", string(bytes))
}

func TestNewAccountStore(t *testing.T) {
	address := types.HexToAddress("0x92c8cae42a94045670cbb0bfcf8f790d9f8097e7")
	accountStore := NewAccountStore("admin", "chain5j.com")
	accountStore.SetAddress(address, nil)
	t.Log("AccountName=", accountStore.AccountName())
	t.Log(accountStore.ContainAddress(types.HexToAddress("0x92c8cae42a94045670cbb0bfcf8f790d9f8097e8")))
	t.Log(accountStore.ContainAddress(address))
	t.Log("AuthorizedRegisterUser=", accountStore.AuthorizedRegisterUser())
	t.Log("AuthorizedRegisterDomain=", accountStore.AuthorizedRegisterDomain())
	t.Log("IsContract=", accountStore.IsContract())

	addressStore, err := accountStore.GetAddressStore(address)
	if err != nil {
		t.Error(err)
	} else {
		t.Log(addressStore)
	}

}
