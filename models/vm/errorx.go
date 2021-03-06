// Package vm
//
// @author: xwc1125
package vm

import "errors"

var (
	ErrOutOfGas                 = errors.New("out of gas")
	ErrCodeStoreOutOfGas        = errors.New("contract creation code storage out of gas")
	ErrDepth                    = errors.New("max call depth exceeded")
	ErrTraceLimitReached        = errors.New("the number of logs reached the specified limit")
	ErrInsufficientBalance      = errors.New("insufficient balance for transfer")
	ErrContractAddressCollision = errors.New("contract address collision")
	ErrNoCompatibleInterpreter  = errors.New("no compatible interpreter")
	ErrGasLimitReached          = errors.New("gas limit reached")
	ErrExecutionReverted        = errors.New("wavm: execution reverted")
	ErrMaxCodeSizeExceeded      = errors.New("wavm: max code size exceeded")
	ErrExecutionAssert          = errors.New("wavm: execution assert")
	ErrMagicNumberMismatch      = errors.New("magic number mismatch")
)
