// Package models
//
// @author: xwc1125
package models

import (
	"github.com/chain5j/logger"
)

// LocalConfig 节点启动本地配置
type LocalConfig struct {
	Log         logger.LogConfig       `json:"log" mapstructure:"log"`                 // 日志配置
	Database    DatabaseConfig         `json:"database" mapstructure:"database"`       // database配置
	Blockchain  BlockchainLocalConfig  `json:"blockchain" mapstructure:"blockchain"`   // blockchain配置
	Packer      PackerLocalConfig      `json:"packer" mapstructure:"packer"`           // packer配置
	TxPool      TxPoolLocalConfig      `json:"tx_pool" mapstructure:"tx_pool"`         // 交易池配置
	NodeKey     NodeKeyLocalConfig     `json:"node_key" mapstructure:"node_key"`       // nodekey模块配置
	Broadcaster BroadcasterLocalConfig `json:"broadcaster" mapstructure:"broadcaster"` // broadcaster模块配置
	P2P         P2PConfig              `json:"p2p" mapstructure:"p2p"`                 // p2p模块配置
	Consensus   ConsensusLocalConfig   `json:"consensus" mapstructure:"consensus"`     // consensus模块配置
}

type DatabaseConfig struct {
	Driver       string `json:"driver" mapstructure:"driver"`               // 驱动类型
	Source       string `json:"source" mapstructure:"source"`               // 资源
	Username     string `json:"username" mapstructure:"username"`           // 用户名
	Password     string `json:"password" mapstructure:"password"`           // 密码
	Metrics      bool   `json:"metrics" mapstructure:"metrics"`             // 是否显示指标
	MetricsLevel uint64 `json:"metrics_level" mapstructure:"metrics_level"` // 指标级别
}

func (c DatabaseConfig) IsMetrics(_metricsLevel uint64) bool {
	return c.Metrics && _metricsLevel <= c.MetricsLevel
}

// BlockchainLocalConfig blockchain配置
type BlockchainLocalConfig struct {
	Metrics      bool   `json:"metrics" mapstructure:"metrics"`             // 是否显示指标
	MetricsLevel uint64 `json:"metrics_level" mapstructure:"metrics_level"` // 指标级别
}

func (c BlockchainLocalConfig) IsMetrics(_metricsLevel uint64) bool {
	return c.Metrics && _metricsLevel <= c.MetricsLevel
}

// TxPoolLocalConfig 交易池配置
type TxPoolLocalConfig struct {
	Capacity     uint64 `json:"capacity" mapstructure:"capacity"`           // 交易池的容量
	CacheDir     string `json:"cache_dir" mapstructure:"cache_dir"`         // 交易池的持久化目录
	Metrics      bool   `json:"metrics" mapstructure:"metrics"`             // 是否显示指标
	MetricsLevel uint64 `json:"metrics_level" mapstructure:"metrics_level"` // 指标级别
}

func (c TxPoolLocalConfig) IsMetrics(_metricsLevel uint64) bool {
	return c.Metrics && _metricsLevel <= c.MetricsLevel
}

// NodeKeyLocalConfig nodeKey模块参数配置
type NodeKeyLocalConfig struct {
	// 节点私钥模块参数属性定义
	FileType     string `json:"file_type" mapstructure:"file_type"`         // 文件类型
	PrvKeyFile   string `json:"prv_key_file" mapstructure:"prv_key_file"`   // 私钥文件
	PubKeyFile   string `json:"pub_key_file" mapstructure:"pub_key_file"`   // 公钥文件
	Password     string `json:"password" mapstructure:"password"`           // 密码
	Metrics      bool   `json:"metrics" mapstructure:"metrics"`             // 是否显示指标
	MetricsLevel uint64 `json:"metrics_level" mapstructure:"metrics_level"` // 指标级别
}

func (c NodeKeyLocalConfig) IsMetrics(_metricsLevel uint64) bool {
	return c.Metrics && _metricsLevel <= c.MetricsLevel
}

// PackerLocalConfig 打包器配置
type PackerLocalConfig struct {
	Metrics      bool   `json:"metrics" mapstructure:"metrics"`             // 是否显示指标
	MetricsLevel uint64 `json:"metrics_level" mapstructure:"metrics_level"` // 指标级别
}

func (c PackerLocalConfig) IsMetrics(_metricsLevel uint64) bool {
	return c.Metrics && _metricsLevel <= c.MetricsLevel
}

// BroadcasterLocalConfig 广播的配置
type BroadcasterLocalConfig struct {
	Metrics      bool   `json:"metrics" mapstructure:"metrics"`             // 是否显示指标
	MetricsLevel uint64 `json:"metrics_level" mapstructure:"metrics_level"` // 指标级别
}

func (c BroadcasterLocalConfig) IsMetrics(_metricsLevel uint64) bool {
	return c.Metrics && _metricsLevel <= c.MetricsLevel
}

// P2PConfig p2p配置
type P2PConfig struct {
	Host string `json:"host" mapstructure:"host"` // host
	Port int    `json:"port" mapstructure:"port"` // 端口号

	KeyPath  string `json:"key_path" mapstructure:"key_path"`   // 私钥路径
	CertPath string `json:"cert_path" mapstructure:"cert_path"` // 证书路径
	IsTls    bool   `json:"is_tls" mapstructure:"is_tls"`       // 安全P2P

	EnablePermission bool     `json:"enable_permission" mapstructure:"enable_permission"` // 权限校验
	MaxPeers         int32    `json:"max_peers" mapstructure:"max_peers"`                 // 最大peer个数
	StaticNodes      []string `json:"static_nodes" mapstructure:"static_nodes"`           // 静态节点.peerID-->p2p.ID
	CaRoots          []string `json:"ca_roots" mapstructure:"ca_roots"`                   // Ca证书

	Metrics      bool   `json:"metrics" mapstructure:"metrics"`             // 是否显示指标
	MetricsLevel uint64 `json:"metrics_level" mapstructure:"metrics_level"` // 指标级别
}

func (c P2PConfig) IsMetrics(_metricsLevel uint64) bool {
	return c.Metrics && _metricsLevel <= c.MetricsLevel
}

// ConsensusLocalConfig consensus配置
type ConsensusLocalConfig struct {
	Metrics      bool   `json:"metrics" mapstructure:"metrics"`             // 是否显示指标
	MetricsLevel uint64 `json:"metrics_level" mapstructure:"metrics_level"` // 指标级别
}

func (c ConsensusLocalConfig) IsMetrics(_metricsLevel uint64) bool {
	return c.Metrics && _metricsLevel <= c.MetricsLevel
}
