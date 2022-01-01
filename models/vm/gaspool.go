// Package vm
//
// @author: xwc1125
package vm

import (
	"fmt"
	"math"

	"github.com/chain5j/logger"
)

// GasPool Gas池，在一个区块中执行交易时，跟踪gas的中消耗
type GasPool uint64

// AddGas 增加可用gas
func (gp *GasPool) AddGas(amount uint64) *GasPool {
	if uint64(*gp) > math.MaxUint64-amount {
		panic("gas pool pushed above uint64")
	}
	*(*uint64)(gp) += amount
	return gp
}

// SubGas 减少可用gas。如果gas不足，返回错误
func (gp *GasPool) SubGas(amount uint64) error {
	// 如果交易的gas大于gasLimit，那么直接返回错误
	if uint64(*gp) < amount {
		logger.Error("SubGas err", "gasPool", gp, "amount", amount, "err", ErrGasLimitReached)
		return ErrGasLimitReached
	}
	*(*uint64)(gp) -= amount
	return nil
}

// Gas 获取剩余的gas
func (gp *GasPool) Gas() uint64 {
	return uint64(*gp)
}

func (gp *GasPool) String() string {
	return fmt.Sprintf("%d", *gp)
}
