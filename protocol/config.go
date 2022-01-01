// Package protocol
//
// @author: xwc1125
package protocol

import (
	"github.com/chain5j/chain5j-pkg/types"
	"github.com/chain5j/chain5j-protocol/models"
)

// Config 配置信息
type Config interface {
	SetDatabase(db DatabaseReader) error

	ChainConfig() models.ChainConfig //链配置
	GenesisBlock() *models.Block     //获取创世区块

	LocalConfig() *models.LocalConfig                 //本地所有配置
	TxSizeLimit() types.StorageSize                   //交易的size
	DatabaseConfig() models.DatabaseConfig            //获取数据库本地配置
	BlockchainConfig() models.BlockchainLocalConfig   //获取blockchain本地配置
	TxPoolConfig() models.TxPoolLocalConfig           //获取交易池本地配置
	NodeKeyConfig() models.NodeKeyLocalConfig         //获取nodeKey本地配置
	PackerConfig() models.PackerLocalConfig           //获取packer本地配置
	EnablePacker() bool                               //是否启动打包
	BroadcasterConfig() models.BroadcasterLocalConfig //获取Broadcaster本地配置
	P2PConfig() models.P2PConfig                      //获取p2p本地配置
	ConsensusConfig() models.ConsensusLocalConfig     //获取consensus本地配置
}
