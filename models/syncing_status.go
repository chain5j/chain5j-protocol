// Package models
//
// @author: xwc1125
package models

import "github.com/chain5j/chain5j-pkg/util/hexutil"

// SyncingStatus 同步状态
type SyncingStatus struct {
	IsSync        bool         `json:"is_sync"`        // true：同步中，false：未同步
	StartingBlock *hexutil.Big `json:"starting_block"` // 开始块
	CurrentBlock  *hexutil.Big `json:"current_block"`  // 当前块
	HighestBlock  *hexutil.Big `json:"highest_block"`  // 目标最高块
}
