// Package crypto
//
// @author: xwc1125
package crypto

import (
	"crypto/ecdsa"
	"fmt"
	"github.com/btcsuite/btcd/btcec"
	"github.com/chain5j/chain5j-pkg/codec/json"
	"github.com/chain5j/chain5j-pkg/crypto/hashalg/sha3"
	"github.com/chain5j/chain5j-pkg/crypto/signature"
	"github.com/chain5j/chain5j-pkg/types"
	"github.com/chain5j/chain5j-pkg/util/hexutil"
	"github.com/chain5j/logger"
	"github.com/chain5j/logger/zap"
	"testing"
)

func init() {
	zap.InitWithConfig(&logger.LogConfig{
		Console: logger.ConsoleLogConfig{
			Level:    4,
			Modules:  "*",
			ShowPath: false,
			Format:   "",
			UseColor: true,
			Console:  true,
		},
		File: logger.FileLogConfig{},
	})
}

func TestSignS256(t *testing.T) {
	testAddrHex := "970e8128ab834e8eac17ab8e3812f010678cf791"
	testPrivHex := "289c2857d4598e37fb9647507e47a309d6133539bf21a8b9cb6df88fd5232032"
	key, _ := signature.HexToECDSA(signature.S256, testPrivHex)
	addr := types.HexToAddress(testAddrHex)
	t.Log("addr", addr.Hex())

	data := []byte("foo")
	{
		msg := sha3.Keccak256(data)
		t.Log("msg", hexutil.Encode(msg))
		sig2, err := btcec.SignCompact(btcec.S256(), &btcec.PrivateKey{
			PublicKey: ecdsa.PublicKey{
				Curve: btcec.S256(),
				X:     key.X,
				Y:     key.Y,
			},
			D: key.D,
		}, msg, true)
		if err != nil {
			t.Fatal(err)
		}
		t.Log("sign2", hexutil.Encode(sig2))
	}
	var recoveredAddr types.Address
	{
		// 【注意】底层实现已经将数据进行了hash处理
		sig, err := Sign(data, key)
		if err != nil {
			t.Errorf("Sign error: %s", err)
		}
		t.Log("pub", hexutil.Encode(sig.PubKey))
		t.Log("sign", hexutil.Encode(sig.Signature))
		sig.PubKey = nil
		recoveredPub, err := RecoverPubKey(data, sig)
		if err != nil {
			t.Errorf("ECRecover error: %s", err)
		}
		recoveredAddr, err = PubkeyToAddress(recoveredPub)
		t.Log("recoveredAddr", recoveredAddr.Hex())
	}

	if addr != recoveredAddr {
		t.Errorf("Address mismatch: want: %x have: %x", addr, recoveredAddr)
	}
}

func TestGenerateKeyPair(t *testing.T) {
	// rsa
	privKey, err := GenerateKeyPair(RSA)
	if err != nil {
		t.Fatal(err)
	}
	testKeyPair(privKey)
	// p256
	privKey, err = GenerateKeyPair(P256)
	if err != nil {
		panic(err)
	}
	testKeyPair(privKey)
	// p384
	privKey, err = GenerateKeyPair(P384)
	if err != nil {
		panic(err)
	}
	testKeyPair(privKey)
	// P521
	privKey, err = GenerateKeyPair(P521)
	if err != nil {
		panic(err)
	}
	testKeyPair(privKey)
	// s256
	privKey, err = GenerateKeyPair(S256)
	if err != nil {
		panic(err)
	}
	testKeyPair(privKey)
	// sm2
	privKey, err = GenerateKeyPair(SM2P256)
	if err != nil {
		panic(err)
	}
	testKeyPair(privKey)
}
func testKeyPair(privKey *KeyPair) {
	p2pPrivateKey, err := ToPrivateKey(privKey.Prv)
	if err != nil {
		panic(err)
	}
	_ = p2pPrivateKey

	p2pPublicKey, err := ToPublicKey(privKey.Pub)
	if err != nil {
		panic(err)
	}
	_ = p2pPublicKey

	signResult, err := Sign([]byte("123"), p2pPrivateKey)
	if err != nil {
		panic(err)
	}
	marshal, _ := json.Marshal(signResult)
	fmt.Println(string(marshal))
	verify, err := Verify([]byte("123"), signResult)
	if err != nil {
		panic(err)
	}
	fmt.Println(verify)
}

func TestMarshalPrivateKey(t *testing.T) {
	// rsa
	{
		privKey, err := GenerateKeyPair(RSA)
		if err != nil {
			panic(err)
		}
		testMarshal(privKey)
	}

	// p256
	{
		privKey, err := GenerateKeyPair(P256)
		if err != nil {
			panic(err)
		}
		testMarshal(privKey)
	}
	// P384
	{
		privKey, err := GenerateKeyPair(P384)
		if err != nil {
			panic(err)
		}
		testMarshal(privKey)
	}
	// P521
	{
		privKey, err := GenerateKeyPair(P521)
		if err != nil {
			panic(err)
		}
		testMarshal(privKey)
	}
	// s256
	{
		privKey, err := GenerateKeyPair(S256)
		if err != nil {
			panic(err)
		}
		testMarshal(privKey)
	}

	// sm2
	{
		privKey, err := GenerateKeyPair(SM2P256)
		if err != nil {
			panic(err)
		}
		testMarshal(privKey)
	}

}
func testMarshal(privKey *KeyPair) {
	privateKeyJsonKey, err := MarshalPrivateKey(privKey.Prv)
	if err != nil {
		panic(err)
	}
	fmt.Println("privateKeyJsonKey", hexutil.Encode(privateKeyJsonKey.SerializeUnsafe()))
	privateKey, err := UnmarshalPrivateKey(privateKeyJsonKey.SerializeUnsafe())
	if err != nil {
		panic(err)
	}
	privateKeyJsonKey, err = MarshalPrivateKey(privateKey.Prv)
	if err != nil {
		panic(err)
	}
	fmt.Println("privateKeyJsonKey2", hexutil.Encode(privateKeyJsonKey.SerializeUnsafe()))

	publicKeyJsonKey, err := MarshalPublicKey(privKey.Pub)
	if err != nil {
		panic(err)
	}
	fmt.Println("publicKeyJsonKey", hexutil.Encode(publicKeyJsonKey.SerializeUnsafe()))
	publicKey, err := UnmarshalPublicKey(publicKeyJsonKey.SerializeUnsafe())
	if err != nil {
		panic(err)
	}
	publicKeyJsonKey, err = MarshalPublicKey(publicKey.Pub)
	if err != nil {
		panic(err)
	}
	fmt.Println("publicKeyJsonKey2", hexutil.Encode(publicKeyJsonKey.SerializeUnsafe()))
}

func TestSign(t *testing.T) {
	// rsa
	{
		privKey, err := GenerateKeyPair(RSA)
		if err != nil {
			panic(err)
		}
		testSign(privKey)
	}

	// p256
	{
		privKey, err := GenerateKeyPair(P256)
		if err != nil {
			panic(err)
		}
		testSign(privKey)
	}
	// P384
	{
		privKey, err := GenerateKeyPair(P384)
		if err != nil {
			panic(err)
		}
		testSign(privKey)
	}
	// P521
	{
		privKey, err := GenerateKeyPair(P521)
		if err != nil {
			panic(err)
		}
		testSign(privKey)
	}
	// s256
	{
		privKey, err := GenerateKeyPair(S256)
		if err != nil {
			panic(err)
		}
		testSign(privKey)
	}

	// sm2
	{
		privKey, err := GenerateKeyPair(SM2P256)
		if err != nil {
			panic(err)
		}
		testSign(privKey)
	}
}
func testSign(privKey *KeyPair) {
	msg := []byte("Hello world!123456789")
	signResult, err := Sign(msg, privKey.Prv)
	if err != nil {
		panic(err)
	}
	marshal, _ := json.Marshal(signResult)
	fmt.Println("signResult", string(marshal))

	verify, err := Verify(msg, signResult)
	if err != nil {
		panic(err)
	}
	fmt.Println("verify", verify)
	if !verify {
		panic("verify err")
	}
	pub, err := RecoverPubKey(msg, signResult)
	if err != nil {
		panic(err)
	}
	fmt.Println("ID", pub)
}
