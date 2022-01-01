// Package protocol
//
// @author: xwc1125
package protocol

import (
	"github.com/chain5j/chain5j-pkg/types"
	"github.com/chain5j/chain5j-protocol/models"
	"github.com/chain5j/chain5j-protocol/models/statetype"
)

// Database 数据库，包含读写
type Database interface {
	DatabaseReader
	DatabaseWriter
}

// DatabaseReader 数据库读
type DatabaseReader interface {
	Start() error // 启动
	Stop() error  // 停止

	ChainConfig() (*models.ChainConfig, error)                                  // 获取最新的链配置
	GetChainConfig(hash types.Hash, height uint64) (*models.ChainConfig, error) // 根据区块hash和高度获取最新的链配置
	GetChainConfigByHash(hash types.Hash) (*models.ChainConfig, error)          // 根据区块hash获取最新的链配置
	GetChainConfigByHeight(height uint64) (*models.ChainConfig, error)          // 根据区块高度获取最新的链配置

	LatestHeader() (*models.Header, error)                            // 获取数据库当前的header
	GetHeader(hash types.Hash, height uint64) (*models.Header, error) // 根据hash及高度获取header
	GetHeaderByHash(hash types.Hash) (*models.Header, error)          // 根据hash获取header
	GetHeaderByHeight(height uint64) (*models.Header, error)          // 根据高度获取header
	GetHeaderHeight(hash types.Hash) (*uint64, error)                 // 根据hash获取其对应的高度
	HasHeader(hash types.Hash, height uint64) (bool, error)           // 根据hash及高度判断header是否存在

	CurrentBlock() (*models.Block, error)                           // 当前区块
	GetBlock(hash types.Hash, height uint64) (*models.Block, error) // 根据hash及高度获取区块
	GetBlockByHash(hash types.Hash) (*models.Block, error)          // 根据hash获取区块
	GetBlockByHeight(height uint64) (*models.Block, error)          // 根据高度获取区块
	HasBlock(hash types.Hash, number uint64) (bool, error)          // 根据hash及number判断区块是否存在

	GetCanonicalHash(height uint64) (bHash types.Hash, err error) // 获取标准hash
	LatestBlockHash() (bHash types.Hash, err error)               // 获取最新的区块hash
	LatestHeaderHash() (bHash types.Hash, err error)              // 获取最新的header头hash

	GetBody(hash types.Hash, height uint64) (*models.Body, error) // 获取body

	GetTransaction(hash types.Hash) (tx models.Transaction, blockHash types.Hash, blockHeight uint64, txIndex uint64, err error) // 获取交易内容
	GetReceipts(bHash types.Hash, height uint64) (statetype.Receipts, error)                                                     // 获取区块所有的receipt
}

// DatabaseWriter 数据库写
type DatabaseWriter interface {
	WriteBlock(block *models.Block) (err error)                                              // 写入区块
	WriteHeader(header *models.Header) (err error)                                           // 写入header
	WriteChainConfig(bHash types.Hash, height uint64, chainConfig *models.ChainConfig) error // 写入链配置
	WriteLatestBlockHash(bHash types.Hash) error                                             // 写入最新的区块hash
	WriteLatestHeaderHash(bHash types.Hash) error                                            // 写入最新的header头hash
	WriteCanonicalHash(bHash types.Hash, height uint64) error                                // 写入标准的hash
	WriteTxsLookup(block *models.Block) error                                                // 写入交易的索引
	WriteReceipts(bHash types.Hash, height uint64, receipts statetype.Receipts) error        // 写入receipt

	DeleteBlock(blockAbs []models.BlockAbstract, currentHeight, desHeight uint64) error // 删除区块
}

// KVStore kv数据库
type KVStore interface {
	KVStoreReader
	KVStoreWriter
	KVStoreDeleter
}

// KVStoreReader kv读
type KVStoreReader interface {
	Has(key []byte) (bool, error)
	Get(key []byte) ([]byte, error)
}

// KVStoreWriter kv写
type KVStoreWriter interface {
	Put(key []byte, value []byte) error
}

// KVStoreDeleter kv删
type KVStoreDeleter interface {
	Delete(key []byte) error
}
