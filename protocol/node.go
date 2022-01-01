// Package protocol
//
// @author: xwc1125
package protocol

import (
	"github.com/chain5j/chain5j-pkg/database/kvstore"
	"github.com/chain5j/chain5j-pkg/types"
)

type Node interface {
	Init() error  // 初始化模块
	Start() error // 启动
	Stop() error  // 停止
	Wait()        // 等待

	Database() Database
	KVDatabase() kvstore.Database
	Config() Config
	BlockReadWriter() BlockReadWriter
	NodeKey() NodeKey
	Apps() Apps
	APIs() APIs
	AddTxPool(txType types.TxType, txPool TxPool)

	SetConsensus(consensus Consensus) error
}
