// Package model
//
// @author: xwc1125
package model

type Code []byte

func (c Code) String() string {
	return string(c)
}
