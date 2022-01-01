// Package crypto
//
// @author: xwc1125
package crypto

import (
	"crypto"
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/rsa"
	"crypto/subtle"
	"crypto/tls"
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"
	"errors"
	"fmt"
	"strings"

	"github.com/btcsuite/btcd/btcec"
	"github.com/chain5j/chain5j-pkg/crypto/base/base58"
	"github.com/chain5j/chain5j-pkg/crypto/hashalg"
	"github.com/chain5j/chain5j-pkg/crypto/hashalg/sha3"
	"github.com/chain5j/chain5j-pkg/crypto/signature"
	"github.com/chain5j/chain5j-pkg/crypto/signature/gmsm"
	"github.com/chain5j/chain5j-pkg/types"
	"github.com/chain5j/chain5j-protocol/models"
	"github.com/chain5j/chain5j-protocol/protocol"
	"github.com/tjfoc/gmsm/sm2"
	x509SM2 "github.com/tjfoc/gmsm/x509"
)

var (
	FormatUnsupported = "unsupported cure crypto. type:%d, name:%s"

	ErrPrivate = errors.New("unsupported private key type")
	ErrPublic  = errors.New("unsupported public key type")

	errKeyBytesNil = errors.New("key bytes is nil")
	errKeyBytes    = errors.New("key bytes err")
	errPassword    = errors.New("password err")
)

// KeyPair 密钥对
type KeyPair struct {
	CryptoType CryptoType
	Prv        crypto.PrivateKey
	Pub        crypto.PublicKey
}

// ==================geneKeyPair========================

// GenerateKeyPair 生成p2p对应的公私钥
func GenerateKeyPair(cryptoType CryptoType) (*KeyPair, error) {
	var (
		prv crypto.PrivateKey
		pub crypto.PublicKey
	)

	switch cryptoType {
	case RSA:
		// rsa
		priv, err := rsa.GenerateKey(rand.Reader, 2048)
		if err != nil {
			return nil, err
		}
		prv = priv
		pub = priv.Public()
	case P256, P384, P521, S256, SM2P256:
		// ecdsa
		priv, err := signature.GenerateKeyWithECDSA(cryptoType.KeyName)
		if err != nil {
			return nil, err
		}

		prv = priv
		pub = priv.Public()
	default:
		return nil, fmt.Errorf(FormatUnsupported, cryptoType.KeyType, cryptoType.KeyName)
	}

	return &KeyPair{
		CryptoType: cryptoType,
		Prv:        prv,
		Pub:        pub,
	}, nil
}

// ==================diff key convert========================

// ToPrivateKey 将原生的PrivateKey转换为protocol.PrivateKey
func ToPrivateKey(prvKey crypto.PrivateKey) (*KeyPair, error) {
	var (
		cryptoType CryptoType
		prv        crypto.PrivateKey
		pub        crypto.PublicKey
		err        error
	)

	switch p := prvKey.(type) {
	case *KeyPair:
		return p, nil
	case *rsa.PrivateKey:
		skBytes := x509.MarshalPKCS1PrivateKey(p)
		sk, err := x509.ParsePKCS1PrivateKey(skBytes)
		if err != nil {
			return nil, err
		}
		if sk.N.BitLen() < MinRsaKeyBits {
			return nil, ErrRsaKeyTooSmall
		}

		cryptoType = RSA
		prv = sk
		pub = sk.PublicKey
	case *ecdsa.PrivateKey:
		curveName := signature.CurveName(p.Curve)
		cryptoType, err = ParseKeyName(curveName)
		if err != nil {
			return nil, err
		}
		prv = p
		pub = p.Public()
	case *btcec.PrivateKey:
		cryptoType = S256
		prv = p
		pub = p.PubKey()
	case *sm2.PrivateKey:
		cryptoType = SM2P256
		prv = p
		pub = p.Public()
	default:
		return nil, fmt.Errorf(FormatUnsupported, cryptoType.KeyType, cryptoType.KeyName)
	}
	return &KeyPair{
		CryptoType: cryptoType,
		Prv:        prv,
		Pub:        pub,
	}, nil
}

// ToPublicKey 将原生的PublicKey转换为protocol.PublicKey
func ToPublicKey(pubKey crypto.PublicKey) (*KeyPair, error) {
	var (
		cryptoType CryptoType
		pub        crypto.PublicKey
	)

	switch p := pubKey.(type) {
	case *KeyPair:
		return p, nil
	case *rsa.PublicKey, rsa.PublicKey:
		pub = p
		cryptoType = RSA
	case *ecdsa.PublicKey:
		pub = p
		cryptoType, _ = ParseKeyName(p.Curve.Params().Name)
	case *btcec.PublicKey:
		pub = p
		cryptoType = S256
	case *sm2.PublicKey:
		pub = p
		cryptoType = SM2P256
	default:
		return nil, fmt.Errorf(FormatUnsupported, cryptoType.KeyType, cryptoType.KeyName)
	}
	return &KeyPair{
		CryptoType: cryptoType,
		Pub:        pub,
	}, nil
}

// ==================to JsonKey========================

// MarshalPrivateKey 私钥转换为JsonKey,data尽可能x509格式
func MarshalPrivateKey(prvKey crypto.PrivateKey) (*JsonKey, error) {
	var (
		err     error
		jsonPrv JsonKey
	)

	switch p := prvKey.(type) {
	case *KeyPair:
		return MarshalPrivateKey(p.Prv)
	case *rsa.PrivateKey:
		jsonPrv.Type = RSA.KeyType
		jsonPrv.Data = x509.MarshalPKCS1PrivateKey(p)
	case *ecdsa.PrivateKey:
		cryptoType, err := ParseKeyName(signature.CurveName(p.Curve))
		if err != nil {
			return nil, err
		}
		jsonPrv.Type = cryptoType.KeyType
		jsonPrv.Data, err = signature.MarshalPrvkeyWithECDSA(p)
		if err != nil {
			return nil, err
		}
	case *btcec.PrivateKey:
		jsonPrv.Type = S256.KeyType
		jsonPrv.Data, err = signature.MarshalPrvkeyWithECDSA(p.ToECDSA())
		if err != nil {
			return nil, err
		}
	case *sm2.PrivateKey:
		jsonPrv.Type = SM2P256.KeyType
		jsonPrv.Data, err = signature.MarshalPrvkeyWithECDSA(gmsm.ToECDSA(p))
		if err != nil {
			return nil, err
		}
	default:
		return nil, ErrPrivate
	}
	return &jsonPrv, err
}

// UnmarshalPrivateKey 解析JsonKeyBytes成对象，将x509格式转换为p2p私钥
func UnmarshalPrivateKey(jsonPrvData []byte) (*KeyPair, error) {
	var (
		err     error
		jsonPrv JsonKey
	)
	err = jsonPrv.Deserialize(jsonPrvData)
	if err == nil {
		return ParsePrivateKeyJsonKey(jsonPrv)
	}
	return nil, err
}

// ParsePrivateKeyJsonKey 解析jsonKey为对象
func ParsePrivateKeyJsonKey(jsonPrv JsonKey) (*KeyPair, error) {
	switch jsonPrv.Type {
	case RSA.KeyType:
		sk, err := x509.ParsePKCS1PrivateKey(jsonPrv.Data)
		if err != nil {
			return nil, err
		}
		if sk.N.BitLen() < MinRsaKeyBits {
			return nil, ErrRsaKeyTooSmall
		}
		return ToPrivateKey(sk)
	case P256.KeyType, P384.KeyType, P521.KeyType, S256.KeyType, SM2P256.KeyType:
		cryptoType, err := ParseKeyType(jsonPrv.Type)
		sk, err := signature.UnMarshalPrvkeyWithECDSA(cryptoType.KeyName, jsonPrv.Data)
		if err != nil {
			return nil, err
		}
		return ToPrivateKey(sk)
	}
	return nil, fmt.Errorf(FormatUnsupported, jsonPrv.Type, "Unkown")
}

func MarshalPrivateKeyX509(prvKey crypto.PrivateKey) (*JsonKey, error) {
	var (
		err     error
		jsonPrv JsonKey
	)

	switch p := prvKey.(type) {
	case *KeyPair:
		return MarshalPrivateKeyX509(p.Prv)
	case *rsa.PrivateKey:
		jsonPrv.Type = RSA.KeyType
		jsonPrv.Data = x509.MarshalPKCS1PrivateKey(p)
	case *ecdsa.PrivateKey:
		curveName := signature.CurveName(p.Curve)
		cryptoType, err := ParseKeyName(curveName)
		if err != nil {
			return nil, err
		}
		jsonPrv.Type = cryptoType.KeyType
		jsonPrv.Data, err = signature.MarshalPrvkeyWithECDSAX509(p)
		if err != nil {
			return nil, err
		}
	case *btcec.PrivateKey:
		jsonPrv.Type = S256.KeyType
		jsonPrv.Data, err = signature.MarshalPrvkeyWithECDSAX509(p.ToECDSA())
	case *sm2.PrivateKey:
		jsonPrv.Type = SM2P256.KeyType
		jsonPrv.Data, err = signature.MarshalPrvkeyWithECDSAX509(gmsm.ToECDSA(p))
		if err != nil {
			return nil, err
		}
	default:
		return nil, ErrPrivate
	}
	return &jsonPrv, err
}

// MarshalPublicKey 将公钥转换为jsonKey格式,jsonKey.data=x509格式
func MarshalPublicKey(ePub crypto.PublicKey) (*JsonKey, error) {
	var (
		err     error
		jsonPrv JsonKey
	)
	switch p := ePub.(type) {
	case *KeyPair:
		return MarshalPublicKey(p.Pub)
	case *rsa.PublicKey:
		jsonPrv.Type = RSA.KeyType
		jsonPrv.Data = x509.MarshalPKCS1PublicKey(p)
	case rsa.PublicKey:
		jsonPrv.Type = RSA.KeyType
		jsonPrv.Data = x509.MarshalPKCS1PublicKey(&p)
	case *ecdsa.PublicKey:
		cryptoType, err := ParseKeyName(signature.CurveName(p.Curve))
		if err != nil {
			return nil, err
		}
		jsonPrv.Type = cryptoType.KeyType
		jsonPrv.Data, err = signature.MarshalPubkeyWithECDSA(p)
	case *btcec.PublicKey:
		jsonPrv.Type = S256.KeyType
		jsonPrv.Data, err = signature.MarshalPubkeyWithECDSA(p.ToECDSA())
	case *sm2.PublicKey:
		jsonPrv.Type = SM2P256.KeyType
		jsonPrv.Data, err = signature.MarshalPubkeyWithECDSA(gmsm.ToECDSAPubKey(p))
	default:
		return nil, ErrPublic
	}
	return &jsonPrv, err
}

// UnmarshalPublicKey 解析jsonKeyBytes成对象，将x509格式转换为p2p公钥
func UnmarshalPublicKey(jsonKeyBytes []byte) (*KeyPair, error) {
	var (
		err     error
		jsonPrv JsonKey
	)
	err = jsonPrv.Deserialize(jsonKeyBytes)
	if err == nil {
		return ParsePublicKeyJsonKey(jsonPrv)
	}
	return nil, err
}

// ParsePublicKeyJsonKey 解析jsonKey为对象
func ParsePublicKeyJsonKey(jsonKey JsonKey) (*KeyPair, error) {
	switch jsonKey.Type {
	case RSA.KeyType:
		pub, err := x509.ParsePKCS1PublicKey(jsonKey.Data)
		if err != nil {
			return nil, err
		}
		return ToPublicKey(pub)
	case P256.KeyType, P384.KeyType, P521.KeyType, S256.KeyType, SM2P256.KeyType:
		cryptoType, err := ParseKeyType(jsonKey.Type)
		if err != nil {
			return nil, err
		}
		pub, err := signature.UnmarshalPubkeyWithECDSA(cryptoType.KeyName, jsonKey.Data)
		if err != nil {
			return nil, err
		}
		return ToPublicKey(pub)
	}
	return nil, ErrPublic
}

func MarshalPublicKeyX509(ePub crypto.PublicKey) (*JsonKey, error) {
	var (
		err     error
		jsonPrv JsonKey
	)
	switch p := ePub.(type) {
	case *KeyPair:
		return MarshalPublicKey(p.Pub)
	case *rsa.PublicKey:
		jsonPrv.Type = RSA.KeyType
		jsonPrv.Data = x509.MarshalPKCS1PublicKey(p)
	case rsa.PublicKey:
		jsonPrv.Type = RSA.KeyType
		jsonPrv.Data = x509.MarshalPKCS1PublicKey(&p)
	case *ecdsa.PublicKey:
		cryptoType, err := ParseKeyName(signature.CurveName(p.Curve))
		if err != nil {
			return nil, err
		}
		jsonPrv.Type = cryptoType.KeyType
		jsonPrv.Data, err = signature.MarshalPubkeyWithECDSAX509(p)
	case *btcec.PublicKey:
		jsonPrv.Type = S256.KeyType
		jsonPrv.Data, err = signature.MarshalPubkeyWithECDSAX509(p.ToECDSA())
	case *sm2.PublicKey:
		jsonPrv.Type = SM2P256.KeyType
		jsonPrv.Data, err = signature.MarshalPubkeyWithECDSAX509(gmsm.ToECDSAPubKey(p))
	default:
		return nil, ErrPublic
	}
	return &jsonPrv, err
}

// ==================sign&verify========================

// Sign 签名数据[节点之间的签名]
func Sign(data []byte, prv crypto.PrivateKey) (sig *signature.SignResult, err error) {
	switch p := prv.(type) {
	case *KeyPair:
		return Sign(data, p.Prv)
	case *rsa.PrivateKey:
		hashed := hashalg.Sha256(data)
		sig1, err := rsa.SignPKCS1v15(rand.Reader, p, crypto.SHA256, hashed[:])
		if err != nil {
			return nil, err
		}
		jsonKey, err := MarshalPublicKey(p.Public())
		if err != nil {
			return nil, err
		}
		return &signature.SignResult{
			Name:      RSA.KeyName,
			PubKey:    jsonKey.SerializeUnsafe(),
			Signature: sig1,
		}, nil
	case *ecdsa.PrivateKey:
		return signature.SignWithECDSA(p, data[:])
	case *btcec.PrivateKey:
		return signature.SignWithECDSA(p.ToECDSA(), data[:])
	case *sm2.PrivateKey:
		return signature.SignWithECDSA(gmsm.ToECDSA(p), data[:])
	default:
		return nil, ErrPrivate
	}

}

// Verify 验证签名内容
// sigBytes 是signResult的bytes值
func Verify(data []byte, signResult *signature.SignResult) (bool, error) {
	cryptoType, err := ParseKeyName(signResult.Name)
	if err != nil {
		return false, err
	}
	switch cryptoType.KeyType {
	case RSA.KeyType:
		hashed := hashalg.Sha256(data)
		pub, err := UnmarshalPublicKey(signResult.PubKey)
		if err != nil {
			return false, err
		}
		err = rsa.VerifyPKCS1v15(pub.Pub.(*rsa.PublicKey), crypto.SHA256, hashed[:], signResult.Signature)
		if err != nil {
			return false, err
		}
		return true, nil
	case P256.KeyType, P384.KeyType, P521.KeyType, S256.KeyType, SM2P256.KeyType:
		return signature.VerifyWithECDSA(signResult, data[:]), nil
	}
	return false, fmt.Errorf(FormatUnsupported, cryptoType.KeyType, cryptoType.KeyName)
}

// RecoverPubKey 从签名中恢复公钥
func RecoverPubKey(data []byte, signResult *signature.SignResult) (crypto.PublicKey, error) {
	cryptoType, err := ParseKeyName(signResult.Name)
	if err != nil {
		return nil, err
	}

	var (
		pub crypto.PublicKey
	)

	switch cryptoType {
	case RSA:
		pub, err = UnmarshalPublicKey(signResult.PubKey)
	case P256, P384, P521, S256, SM2P256:
		hash := signature.HashMsg(cryptoType.KeyName, data)
		pub1, err := signature.SigToPub(hash[:], signResult)
		if err != nil {
			return nil, err
		}
		pub = &KeyPair{
			CryptoType: cryptoType,
			Pub:        pub1,
		}
	case Ed25519:
		pub, err = UnmarshalPublicKey(signResult.PubKey)
	}

	if err != nil {
		return nil, err
	}
	return pub, nil
}

// ==================pem or der========================

// ParsePrivateKeyPem parse key pem to privateKey
func ParsePrivateKeyPem(keyPemBytes, certPemBytes []byte, pwd []byte) (privateKey *KeyPair, err error) {
	if keyPemBytes == nil || len(keyPemBytes) <= 0 {
		return nil, errKeyBytesNil
	}

	if !strings.Contains(string(keyPemBytes), "-----BEGIN") {
		keyBytes, err := hex.DecodeString(string(keyPemBytes))
		if err != nil {
			return nil, fmt.Errorf("fail to hex decode pem: [%v]", err)
		}
		return PrivKeyWithDer(keyBytes)
	}

	// tls
	if certPemBytes != nil && len(certPemBytes) > 0 {
		if sk, err := tls.X509KeyPair(certPemBytes, keyPemBytes); err == nil {
			return ToPrivateKey(sk)
		}
	}

	block, _ := pem.Decode(keyPemBytes)
	if block == nil {
		return PrivKeyWithDer(keyPemBytes)
	}

	keyBytes := block.Bytes
	if x509.IsEncryptedPEMBlock(block) {
		if len(pwd) <= 0 {
			return nil, errPassword
		}

		keyBytes, err = x509.DecryptPEMBlock(block, pwd)
		if err != nil {
			return nil, fmt.Errorf("fail to decrypt pem: [%s]", err)
		}
	}

	return PrivKeyWithDer(keyBytes)

}

// ParsePublicKeyPem parse key pem to publicKey
func ParsePublicKeyPem(keyPemBytes []byte) (*KeyPair, error) {
	if keyPemBytes == nil || len(keyPemBytes) <= 0 {
		return nil, errKeyBytesNil
	}

	if !strings.Contains(string(keyPemBytes), "-----BEGIN") {
		keyBytes, err := hex.DecodeString(string(keyPemBytes))
		if err != nil {
			return nil, fmt.Errorf("fail to hex decode public pem: [%v]", err)
		}
		return PubKeyWithDer(keyBytes)
	}

	block, _ := pem.Decode(keyPemBytes)
	if block == nil {
		return PubKeyWithDer(keyPemBytes)
	}

	return PubKeyWithDer(block.Bytes)
}

// PrivKeyWithDer 将der格式转换为crypto.PrivateKey
func PrivKeyWithDer(der []byte) (*KeyPair, error) {
	// rsa
	if sk, err := x509.ParsePKCS1PrivateKey(der); err == nil {
		if sk.N.BitLen() < MinRsaKeyBits {
			return nil, ErrRsaKeyTooSmall
		}
		return ToPrivateKey(sk)
	}

	if sk, err := x509.ParsePKCS8PrivateKey(der); err == nil {
		return ToPrivateKey(sk)
	}

	// ecdsa
	if sk, err := x509.ParseECPrivateKey(der); err == nil {
		return ToPrivateKey(sk)
	}

	// sm2
	if sk, err := x509SM2.ParsePKCS8UnecryptedPrivateKey(der); err == nil {
		return ToPrivateKey(sk)
	}

	// chain5j
	if key, err := UnmarshalPrivateKey(der); err == nil {
		return key, nil
	}

	return nil, errKeyBytes
}

// PubKeyWithDer 将der格式转换为crypto.PublicKey
func PubKeyWithDer(der []byte) (*KeyPair, error) {
	// rsa
	if sk, err := x509.ParsePKCS1PublicKey(der); err == nil {
		return ToPublicKey(sk)
	}

	if sk, err := x509.ParsePKIXPublicKey(der); err == nil {
		return ToPublicKey(sk)
	}

	// sm2
	if sk, err := x509SM2.ParseSm2PublicKey(der); err == nil {
		return ToPublicKey(sk)
	}

	// chain5j
	if key, err := UnmarshalPublicKey(der); err == nil {
		return key, nil
	}

	return nil, errKeyBytes
}

// CryptoLabel 获取
func CryptoLabel(cryptoType CryptoType) string {
	switch cryptoType {
	case RSA:
		return "RSA"
	case P256, P384, P521, S256, SM2P256:
		return "EC"
	case Ed25519:
		return "Ed25519"
	default:
		return ""
	}
}

// ==================tool========================

// BasicEquals 比较两个key是否一样。实际是比较其bytes
func BasicEquals(k1, k2 protocol.Key) bool {
	if k1.Type() != k2.Type() {
		return false
	}

	a, err := k1.Raw()
	if err != nil {
		return false
	}
	b, err := k2.Raw()
	if err != nil {
		return false
	}
	return subtle.ConstantTimeCompare(a, b) == 1
}

// IDFromPrivateKey 通过私钥获取PeerID
func IDFromPrivateKey(sk protocol.PrivKey) (models.NodeID, error) {
	return IDFromPublicKey(sk.GetPublic())
}

// IDFromPublicKey 根据公钥生成PeerID
func IDFromPublicKey(pk protocol.PubKey) (models.NodeID, error) {
	b, err := pk.Marshal()
	if err != nil {
		return "", err
	}
	alg := pk.Hash()
	h := alg()
	h.Write(b)
	bytes := h.Sum(nil)
	return models.NodeID(base58.Encode(bytes)), nil
}

// PubkeyToAddress 根据公钥生成地址
func PubkeyToAddress(pk crypto.PublicKey) (types.Address, error) {
	jsonKey, err := MarshalPublicKey(pk)
	if err != nil {
		return types.EmptyAddress, err
	}
	return types.BytesToAddress(sha3.Keccak256(jsonKey.Data[1:])[12:]), nil
}
