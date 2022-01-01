// Package models
//
// @author: xwc1125
package models

import (
	"github.com/chain5j/chain5j-pkg/codec"
	"github.com/chain5j/chain5j-pkg/collection/maps/hashmap"
)

// ChainConfig 链配置
type ChainConfig struct {
	ChainID     uint64 `json:"chain_id" mapstructure:"chain_id"`         // 链ID
	ChainName   string `json:"chain_name" mapstructure:"chain_name"`     // 链名称,Chain5j
	VersionName string `json:"version_name" mapstructure:"version_name"` // 版本名称,v1.0.0
	VersionCode uint64 `json:"version_code" mapstructure:"version_code"` // 版本code,1

	GenesisHeight uint64 `json:"genesis_height" mapstructure:"genesis_height"` // 创世高度
	TxSizeLimit   uint64 `json:"tx_size_limit" mapstructure:"tx_size_limit"`   // 单笔交易最大限制 单位 KB

	Packer *PackerConfig `json:"packer" mapstructure:"packer" rlp:"nil"` // 打包逻辑处理
	// TODO 【xwc1125】需要将共识配置设置成map的方式。在每个共识内部自己去解析
	Consensus *ConsensusConfig `json:"consensus,omitempty" mapstructure:"consensus" rlp:"nil"`
	StateApp  *StateAppConfig  `json:"state_app,omitempty" mapstructure:"state_app" rlp:"nil"`
}

type WorkerType uint64

const (
	Timing             WorkerType = iota // 定时出块
	TransactionTrigger                   // 交易触发，无空块
	BlockStack                           // 定量空块
)

// PackerConfig 打包器的配置
type PackerConfig struct {
	WorkerType           WorkerType `json:"worker_type" mapstructure:"worker_type"`                         // 出块类型
	BlockMaxTxsCapacity  uint64     `json:"block_max_txs_capacity" mapstructure:"block_max_txs_capacity"`   // 每个区块的交易最大个数
	BlockMaxSize         uint64     `json:"block_max_size" mapstructure:"block_max_size"`                   // 区块最大size(kb)
	BlockMaxIntervalTime uint64     `json:"block_max_interval_time" mapstructure:"block_max_interval_time"` // 最大出块间隔(ms)[=0时，采用交易触发]
	BlockGasLimit        uint64     `json:"block_gas_limit" mapstructure:"block_gas_limit"`                 // 区块的最大gas
	Period               uint64     `json:"period" mapstructure:"period"`                                   // 定时出块（=0时，采用交易触发）【毫秒】
	EmptyBlocks          uint64     `json:"empty_blocks" mapstructure:"empty_blocks"`                       // 空块的个数
	Timeout              uint64     `json:"timeout" mapstructure:"timeout"`                                 // 超时时间【毫秒】
	MatchTxsCapacity     bool       `json:"match_txs_capacity" yaml:"match_txs_capacity"`                   // 是否满足最大txs才开始打包
}

type ConsensusConfig struct {
	Name string           `json:"name,omitempty"`      // 共识名称
	Data *hashmap.HashMap `json:"consensus,omitempty"` // 共识具体配置
}

func (c *ConsensusConfig) ToConsensus() (*Consensus, error) {
	toBytes, err := codec.Coder().Encode(c.Data)
	if err != nil {
		return nil, err
	}
	return &Consensus{
		Name:      c.Name,
		Consensus: toBytes,
	}, nil
}

// StateAppConfig 状态应用配置
type StateAppConfig struct {
	UseEthereum bool `json:"use_eth"` // 是否使用eth
}
