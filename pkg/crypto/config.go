// Package crypto
//
// @author: xwc1125
package crypto

import (
	"errors"
	"fmt"
	"github.com/chain5j/chain5j-pkg/codec/json"
	"github.com/chain5j/logger"
)

const (
	KeyType_RSA     int32 = iota
	KeyType_Ed25519 int32 = 1
	KeyType_P256    int32 = 3
	KeyType_S256    int32 = 5
	KeyType_SM2     int32 = 6
	KeyType_P384    int32 = 7
	KeyType_P521    int32 = 8
)

var (
	UNKNOWN = CryptoType{"UNKNOWN", -1}
	RSA     = CryptoType{"RSA", KeyType_RSA}
	Ed25519 = CryptoType{"Ed25519", KeyType_Ed25519}
	P256    = CryptoType{"P-256", KeyType_P256}
	P384    = CryptoType{"P-384", KeyType_P384}
	P521    = CryptoType{"P-521", KeyType_P521}
	S256    = CryptoType{"S-256", KeyType_S256}
	SM2P256 = CryptoType{"SM2-P-256", KeyType_SM2}
)

var (
	MinRsaKeyBits = 1024
)

var (
	ErrRsaKeyTooSmall = errors.New("rsa key too small")

	FormatKeyType = "unsupported the key type:%d"
	FormatKeyName = "unsupported the key name:%s"
)

// CryptoType crypto type
type CryptoType struct {
	KeyName string
	KeyType int32
}

// ParseKeyType 将int32转换为对应的CryptoType
func ParseKeyType(keyType int32) (CryptoType, error) {
	switch keyType {
	case RSA.KeyType:
		return RSA, nil
	case Ed25519.KeyType:
		return Ed25519, nil
	case P256.KeyType:
		return P256, nil
	case S256.KeyType:
		return S256, nil
	case SM2P256.KeyType:
		return SM2P256, nil
	case P384.KeyType:
		return P384, nil
	case P521.KeyType:
		return P521, nil
	default:
		return RSA, fmt.Errorf(FormatKeyType, keyType)
	}
}

// ParseKeyName 将keyName转换为CryptoType
func ParseKeyName(keyName string) (CryptoType, error) {
	switch keyName {
	case RSA.KeyName:
		return RSA, nil
	case Ed25519.KeyName:
		return Ed25519, nil
	case P256.KeyName:
		return P256, nil
	case S256.KeyName:
		return S256, nil
	case SM2P256.KeyName:
		return SM2P256, nil
	case P384.KeyName:
		return P384, nil
	case P521.KeyName:
		return P521, nil
	default:
		return RSA, fmt.Errorf(FormatKeyName, keyName)
	}
}

// JsonKey key对象
type JsonKey struct {
	Type int32  `json:"type" mapstructure:"type"`
	Data []byte `json:"data" mapstructure:"data"`
}

func (j *JsonKey) Serialize() ([]byte, error) {
	return json.Marshal(j)
}

func (j *JsonKey) SerializeUnsafe() []byte {
	bytes, err := json.Marshal(j)
	if err != nil {
		logger.Error("jsonKey SerializeUnsafe err", "err", err)
	}
	return bytes
}

func (j *JsonKey) Deserialize(data []byte) error {
	return json.Unmarshal(data, &j)
}
