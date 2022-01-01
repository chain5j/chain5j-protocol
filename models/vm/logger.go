// Package vm
//
// @author: xwc1125
package vm

import (
	"math/big"

	"github.com/chain5j/chain5j-pkg/math"
	"github.com/chain5j/chain5j-pkg/types"
	"github.com/chain5j/chain5j-pkg/util/hexutil"
	"github.com/chain5j/chain5j-protocol/protocol"
)

type Storage map[types.Hash]types.Hash

func (s Storage) Copy() Storage {
	cpy := make(Storage)
	for key, value := range s {
		cpy[key] = value
	}

	return cpy
}

// LogConfig 日志配置
type LogConfig struct {
	DisableMemory  bool // 是否内存捕获
	DisableStack   bool // 是否stack捕获
	DisableStorage bool // 是否storage捕获
	Debug          bool // 是否打印捕获日志
	Limit          int  // 最大的日志输出，0代表不限制
}

//go:generate gencodec -type StructLog -field-override structLogMarshaling -out gen_structlog.go

// StructLog is emitted to the VM each cycle and lists information about the current internal state
// prior to the execution of the statement.
type StructLog struct {
	Pc         uint64                    `json:"pc"`
	Op         protocol.OPCode           `json:"op"`
	Gas        uint64                    `json:"gas"`
	GasCost    uint64                    `json:"gas_cost"`
	Memory     []byte                    `json:"memory"`
	MemorySize int                       `json:"memory_size"`
	Stack      []*big.Int                `json:"stack"`
	Storage    map[types.Hash]types.Hash `json:"-"`
	Depth      int                       `json:"depth"`
	Err        error                     `json:"-"`
}

type DebugLog struct {
	PrintMsg string `json:"print_msg"`
}

// overrides for gencodec
type structLogMarshaling struct {
	Stack       []*math.HexOrDecimal256
	Gas         math.HexOrDecimal64
	GasCost     math.HexOrDecimal64
	Memory      hexutil.Bytes
	OpName      string `json:"op_name"` // adds call to OpName() in MarshalJSON
	ErrorString string `json:"error"`   // adds call to ErrorString() in MarshalJSON
}

func (s *StructLog) OpName() string {
	return s.Op.String()
}

func (s *StructLog) ErrorString() string {
	if s.Err != nil {
		return s.Err.Error()
	}
	return ""
}
