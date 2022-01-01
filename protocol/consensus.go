// Package protocol
//
// @author: xwc1125
package protocol

import (
	"context"

	"github.com/chain5j/chain5j-protocol/models"
)

// Consensus 共识引擎
type Consensus interface {
	Start() error // 启动引擎
	Stop() error  // 停止引擎

	Begin() error
	// Commit()error

	// VerifyHeader 检查Header是否符合引擎的一致规则。
	// 可在此选择验证seal，或通过VerifySeal方法明确验证。
	// 同步或接收到广播的区块信息时调用此方法。此处header已被签名并包含共识内容
	VerifyHeader(blockReader BlockReader, header *models.Header) error
	// VerifyHeaders 批量验证区块头
	VerifyHeaders(blockReader BlockReader, headers []*models.Header, seals []bool) (chan<- struct{}, <-chan error)

	// Prepare 根据规则初始化header的字段。以内联方式执行
	// 如初始化header.Consensus，header.Timestamp，此时的header还未被签名
	Prepare(blockReader BlockReader, header *models.Header) error

	// Finalize 运行交易后状态修改（例如区块奖励）并组装最终区块。
	// 注意：header和stateDb可能会更新，以反映最终确定时发生的任何共识规则（例如区块奖励）。
	// 组装header及txs形成block，此时的block还未被签名
	Finalize(blockReader BlockReader, header *models.Header, txs models.Transactions) (*models.Block, error)

	// Seal 根据block进行密封处理（区块签名），并且放入chan中
	// 注意：该方法立即返回，并将异步发送结果。根据一致性算法，还可能返回多个结果。
	// 使用nodeKey对区块进行签名处理，并真实地开启共识核心处理
	Seal(ctx context.Context, blockReader BlockReader, block *models.Block, results chan<- *models.Block) error
}
