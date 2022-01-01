// Package models
//
// @author: xwc1125
package models

import "github.com/chain5j/chain5j-pkg/types"

type AppsStatus []TxsStatus

// TxsStatus 交易集状态
type TxsStatus struct {
	TxType     types.TxType          `json:"tx_type"`
	StateRoots []byte                `json:"state_roots"`
	GasUsed    uint64                `json:"gas_used"`
	OkTxs      TransactionSortedList `json:"ok_txs"`
	ErrTxs     TransactionSortedList `json:"err_txs"`
}

// TxStatus 交易状态
type TxStatus uint32

const (
	TxStatus_Unkown    TxStatus = 0 // 未知
	TxStatus_Failed    TxStatus = 1 // 失败
	TxStatus_Succeeded TxStatus = 2 // 成功
	TxStatus_Waiting   TxStatus = 3 // 等待中
	TxStatus_Pending   TxStatus = 4 // 提交中
)

var TxStatus_name = map[uint32]string{
	0: "Unkown",
	1: "Failed",
	2: "Succeeded",
	3: "Waiting",
	4: "Pending",
}

var TxStatus_value = map[string]uint32{
	"Unkown":    0,
	"Failed":    1,
	"Succeeded": 2,
	"Waiting":   3,
	"Pending":   4,
}
