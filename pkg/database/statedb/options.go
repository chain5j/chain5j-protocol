// Package statedb
//
// @author: xwc1125
package statedb

import "fmt"

type option func(f *StateDB) error

func apply(f *StateDB, opts ...option) error {
	for _, opt := range opts {
		if opt == nil {
			continue
		}
		if err := opt(f); err != nil {
			return fmt.Errorf("option apply err:%v", err)
		}
	}
	return nil
}

func WithConfig(config *Config) option {
	return func(f *StateDB) error {
		if config != nil {
			f.config = config
		}
		return nil
	}
}
