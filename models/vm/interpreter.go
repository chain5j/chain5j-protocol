// Package vm
//
// @author: xwc1125
package vm

import "github.com/chain5j/chain5j-protocol/protocol"

// Config 解释器的配置
type Config struct {
	Debug                   bool            // 是否debug
	Tracer                  protocol.Tracer // op code的logger
	NoRecursion             bool            // 禁用解释器的call, callcode, delegate call and create
	EnablePreimageRecording bool            // 是否能够镜像记录
}
