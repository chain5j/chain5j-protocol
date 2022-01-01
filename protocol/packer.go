// Package protocol
//
// @author: xwc1125
package protocol

// Packer 打包器
type Packer interface {
	Start() error
	Stop() error
}
