package logger

import (
	logger_contetx "github.com/go-tech-lab/framework/logger/src/log_context"
	. "github.com/go-tech-lab/framework/logger/src/log_filter"
)

type AfterWriteFunc func(p []byte, n int, err error)

//LogFields 日志field
type LogFields map[string]interface{}

//ILoggerCore 暴露给外部调用的实际接口
type ILoggerCore interface {
	Debugf(format string, params ...interface{})

	Infof(format string, params ...interface{})

	Warnf(format string, params ...interface{})

	Errorf(format string, params ...interface{})

	Debug(v ...interface{})

	Info(v ...interface{})

	Warn(v ...interface{})

	Error(v ...interface{})

	Debugw(msg string, logFields LogFields)

	Infow(msg string, logFields LogFields)

	Warnw(msg string, logFields LogFields)

	Errorw(msg string, logFields LogFields)
}

type ILoggerOps interface {
	//Filter 获取日志过滤器
	Filter() ILogFilter

	//WithFilter 设置过滤器
	WithFilter(filter ILogFilter) ILogger

	//WithModuleTag 设置RoutineLocal的ModuleTag
	WithModuleTag(moduleTag string) ILogger

	//SetDefaultModuleTag 设置全局默认的ModuleTag
	SetDefaultModuleTag(defaultModuleTag string) ILogger

	//WithMonitor 设置监控函数
	WithMonitor(monitorFunc func(logContext *logger_contetx.LogContext)) ILogger

	//LogWrapper 获取日志内容包装器
	LogWrapper() ILogWrapper

	//WithLogWrapper 设置日志包装器
	WithLogWrapper(logDecorator ILogWrapper) ILogger

	//LogFilterDecider 获取日志过滤器Decider
	LogFilterDecider() LogFilterDecider

	//WithLogFilterDecider 设置日志过滤器Decider
	WithLogFilterDecider(logFilterDecider LogFilterDecider) ILogger

	//LogContextLocal RoutineLocal日志上下文信息
	LogContextLocal() logger_contetx.ILogContextLocal

	//RawLogger 原生的logger组件（zap是zap.SugaredLogger）
	RawLogger() interface{}

	//SetEnableLogLevel 设置允许输出的日志级别
	SetEnableLogLevel(enableLogLevel string) ILogger
}

//ILogger 日志组件核心接口
type ILogger interface {
	ILoggerCore

	ILoggerOps
}
