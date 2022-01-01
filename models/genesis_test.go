// Package models
//
// @author: xwc1125
package models

import (
	"github.com/chain5j/chain5j-pkg/math"
	"github.com/chain5j/chain5j-pkg/types"
	"github.com/chain5j/chain5j-pkg/util/hexutil"
	"github.com/chain5j/chain5j-pkg/util/ioutil"
	"github.com/chain5j/logger"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
	"math/big"
	"reflect"
)

func GetGenesisConfig(genesisJson string) *Genesis {
	if !ioutil.PathExists(genesisJson) {
		logger.Crit("genesisJson is no exist")
	}
	genesisConfig := &Genesis{}
	v := viper.New()
	v.SetConfigFile(genesisJson)
	v.AutomaticEnv()
	if err := v.ReadInConfig(); err != nil {
		logger.Crit("viper.ReadInConfig()", "err", err)
	}

	optDecode := viper.DecodeHook(mapstructure.ComposeDecodeHookFunc(
		mapstructure.StringToTimeDurationHookFunc(),
		StringToByteSizesHookFunc(),
	))

	err := v.Unmarshal(&genesisConfig, optDecode)
	if err != nil {
		logger.Crit("getGenesisConfig is error", "err", err)
	}
	return genesisConfig
}

func StringToByteSizesHookFunc() mapstructure.DecodeHookFunc {
	return func(
		f reflect.Type,
		t reflect.Type,
		data interface{}) (interface{}, error) {
		if f.Kind() != reflect.String {
			return data, nil
		}

		switch t {
		case reflect.TypeOf(types.Address{}):
			// Convert it by parsing
			raw := data.(string)
			address := types.HexToAddress(raw)
			return address, nil
		case reflect.TypeOf(types.Hash{}):
			raw := data.(string)
			hash := types.HexToHash(raw)
			return hash, nil
		case reflect.TypeOf(math.HexOrDecimal256{}):
			raw := data.(string)
			result := new(math.HexOrDecimal256)
			result.UnmarshalText([]byte(raw))
			return result, nil
		case reflect.TypeOf(big.Int{}):
			raw := data.(string)
			result := new(big.Int)
			result.UnmarshalText([]byte(raw))
			return result, nil
		case reflect.TypeOf(hexutil.Bytes{}):
			raw := data.(string)
			result := hexutil.MustDecode(raw)
			return result, nil
		default:
			return data, nil
		}
	}
}
