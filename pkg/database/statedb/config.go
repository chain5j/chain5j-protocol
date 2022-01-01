// Package statedb
//
// @author: xwc1125
package statedb

type Config struct {
	Metrics      bool   `json:"metrics" mapstructure:"metrics"`             // 是否显示指标
	MetricsLevel uint64 `json:"metrics_level" mapstructure:"metrics_level"` // 指标级别
}

func (c Config) IsMetrics(_metricsLevel uint64) bool {
	return c.Metrics && _metricsLevel <= c.MetricsLevel
}
