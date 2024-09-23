package log_filter

import (
	logger_contetx "github.com/go-tech-lab/framework/logger/src/log_context"
)

// ILogFilter 日志过滤器接口
type ILogFilter interface {
	//Filter false: 有过滤规则，true 无规则打印日志
	Filter(logContext *logger_contetx.LogContext) bool
}

//LogFilterDecider 日志过滤决策-是启用日志过滤器
// return false: 忽略日志过滤器，强制输出日志，比如对一些特定的流量强制输出日志
// return true:  正常执行日志过滤器
type LogFilterDecider func() bool
