// Package models
//
// @author: xwc1125
package models

import (
	crypto2 "github.com/chain5j/chain5j-pkg/crypto/signature"
	"github.com/chain5j/chain5j-pkg/types"
	"github.com/chain5j/chain5j-pkg/util/hexutil"
	"github.com/chain5j/chain5j-protocol/models/accounts"
	"github.com/chain5j/chain5j-protocol/models/statetype"
	"math/big"
)

// Genesis 创世文件内容，此文件用于生成genesis.json
// 在进行sign_genesis时，需要将ConsensusConfig转换为chainConfig中的Consensus
// 将GenesisBlock转换为chainConfig中的GenesisBlock
// 将AccountAlloc，AddressAlloc转换为txs
type Genesis struct {
	ChainConfig     *ChainConfig        `json:"config" mapstructure:"config"`                               // 链配置
	ConsensusConfig *ConsensusConfig    `json:"consensus_config,omitempty" mapstructure:"consensus_config"` // 共识配置
	GenesisBlock    *GenesisBlock       `json:"genesis_block,omitempty" mapstructure:"genesis_block"`       // 创世区块
	AccountAlloc    GenesisAccountAlloc `json:"account_alloc,omitempty" mapstructure:"account_alloc"`       // 账户配置
	AddressAlloc    GenesisAddressAlloc `json:"address_alloc,omitempty" mapstructure:"address_alloc"`       // 地址配置
}

// GenesisBlock 创世文件中区块的配置
type GenesisBlock struct {
	ParentHash   types.Hash          `json:"parent_hash,omitempty"`
	Height       uint64              `json:"height,omitempty"`
	StateRoots   hexutil.Bytes       `json:"state_roots,omitempty"`
	TxsRoot      []types.Hash        `json:"txs_root,omitempty"`
	TxsCount     uint64              `json:"txs_count,omitempty"`
	Timestamp    uint64              `json:"timestamp,omitempty"`
	GasUsed      uint64              `json:"gas_used,omitempty"`
	GasLimit     uint64              `json:"gas_limit,omitempty"`
	Consensus    *Consensus          `json:"consensus,omitempty"`
	ArchiveHash  *types.Hash         `json:"archive_hash,omitempty"`
	ReceiptsRoot []types.Hash        `json:"receipts_root,omitempty"` // 交易收据的hash
	LogsBloom    *statetype.Bloom    `json:"logs_bloom,omitempty"`    // logs的bloom
	Extra        hexutil.Bytes       `json:"extra,omitempty"`
	Signature    *crypto2.SignResult `json:"signature,omitempty"`
}

// GenesisAddress 创世地址信息
type GenesisAddress struct {
	Code    *hexutil.Bytes            `json:"code,omitempty" mapstructure:"code"`       // code
	Storage map[types.Hash]types.Hash `json:"storage,omitempty" mapstructure:"storage"` // 存储
	Balance *big.Int                  `json:"balance,omitempty" mapstructure:"balance"` // 余额
	Nonce   uint64                    `json:"nonce,omitempty" mapstructure:"nonce"`     // nonce
}

// GenesisAddressAlloc address ==> GenesisAddress
type GenesisAddressAlloc map[types.Address]GenesisAddress

// GenesisAccountAlloc account ==> accounts.AccountStore
type GenesisAccountAlloc map[string]accounts.AccountStore
