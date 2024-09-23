package test

import (
	"fmt"
	logger_contetx "github.com/go-tech-lab/framework/logger/src/log_context"
	"github.com/go-tech-lab/framework/logger/src/log_filter"
	_ "github.com/go-tech-lab/framework/logger/src/logger/impl"
	"github.com/go-tech-lab/framework_common/src/context"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestLogger_Warn(t *testing.T) {
	asserts := require.New(t)
	logger := mockZapLogger()
	logger.Warn("I am testing credit logger.Warn framework. ", "key1:", "value1", "key2:", "value2", "key3:", "value3")
	logContext := logger_contetx.GetLogContext()
	funcContext := context.GetFuncContext(1)
	asserts.Equal(logContext.FuncContext.FileName, funcContext.FileName)
	asserts.Equal(logContext.FuncContext.FuncName, funcContext.FuncName)
	asserts.Equal(logContext.FuncContext.Line+2, funcContext.Line)

	//mustFalseFilter
	mustFalseFilter := log_filter.WrapLogFilterFunc(func(logContext *logger_contetx.LogContext) bool {
		return false
	})
	logger.WithFilter(mustFalseFilter)
	logger.Warn("I am testing credit logger.Warn framework. ", "key1:", "value1", "key2:", "value2", "key3:", "value3")
	logContext = logger_contetx.GetLogContext()
	funcContext = context.GetFuncContext(1)
	asserts.Equal(logContext.FuncContext.FileName, funcContext.FileName)
	asserts.Equal(logContext.FuncContext.FuncName, funcContext.FuncName)
	asserts.Equal(logContext.FuncContext.Line+2, funcContext.Line)

	//not filter, no monitor
	logger.WithFilter(nil)
	logger.WithMonitor(nil)
	logger.Warn("I am testing credit logger.Warn framework. ", "key1:", "value1", "key2:", "value2", "key3:", "value3")
	logContext = logger_contetx.GetLogContext()
	funcContext = context.GetFuncContext(1)
	asserts.Equal(logContext.FuncContext.FileName, funcContext.FileName)
	asserts.Equal(logContext.FuncContext.FuncName, funcContext.FuncName)
	asserts.Equal(logContext.FuncContext.Line+2, funcContext.Line)

	//filter false , monitor not nil
	mustTureFilter := log_filter.WrapLogFilterFunc(func(logContext *logger_contetx.LogContext) bool {
		return true
	})
	logger.WithFilter(mustTureFilter)
	logger.WithMonitor(func(logContext *logger_contetx.LogContext) {
		fmt.Println("Test logger.Monitor")
	})
	logger.Warn("I am testing credit logger.Warn framework. ", "key1:", "value1", "key2:", "value2", "key3:", "value3")
	logContext = logger_contetx.GetLogContext()
	funcContext = context.GetFuncContext(1)
	asserts.Equal(logContext.FuncContext.FileName, funcContext.FileName)
	asserts.Equal(logContext.FuncContext.FuncName, funcContext.FuncName)
	asserts.Equal(logContext.FuncContext.Line+2, funcContext.Line)
}

func TestLogger_Warnf(t *testing.T) {
	asserts := require.New(t)
	logger := mockZapLogger()
	logger.Warnf("I am testing credit logger.Warnf framework. key1:%s, key2:%s, key3:%s .", "value1", "value2", "value3")
	logContext := logger_contetx.GetLogContext()
	funcContext := context.GetFuncContext(1)
	asserts.Equal(logContext.FuncContext.FileName, funcContext.FileName)
	asserts.Equal(logContext.FuncContext.FuncName, funcContext.FuncName)
	asserts.Equal(logContext.FuncContext.Line+2, funcContext.Line)

	//mustFalseFilter
	mustFalseFilter := log_filter.WrapLogFilterFunc(func(logContext *logger_contetx.LogContext) bool {
		return false
	})
	logger.WithFilter(mustFalseFilter)
	logger.Warnf("I am testing credit logger.Warnf framework. key1:%s, key2:%s, key3:%s .", "value1", "value2", "value3")
	logContext = logger_contetx.GetLogContext()
	funcContext = context.GetFuncContext(1)
	asserts.Equal(logContext.FuncContext.FileName, funcContext.FileName)
	asserts.Equal(logContext.FuncContext.FuncName, funcContext.FuncName)
	asserts.Equal(logContext.FuncContext.Line+2, funcContext.Line)

	//filter false , monitor not nil
	mustTureFilter := log_filter.WrapLogFilterFunc(func(logContext *logger_contetx.LogContext) bool {
		return true
	})
	logger.WithFilter(mustTureFilter)
	logger.WithMonitor(func(logContext *logger_contetx.LogContext) {
		fmt.Println("Test logger.Monitor")
	})
	logger.Warnf("I am testing credit logger.Warnf framework. key1:%s, key2:%s, key3:%s .", "value1", "value2", "value3")
	logContext = logger_contetx.GetLogContext()
	funcContext = context.GetFuncContext(1)
	asserts.Equal(logContext.FuncContext.FileName, funcContext.FileName)
	asserts.Equal(logContext.FuncContext.FuncName, funcContext.FuncName)
	asserts.Equal(logContext.FuncContext.Line+2, funcContext.Line)
}

func TestLogger_Warnw(t *testing.T) {
	asserts := require.New(t)
	logger := mockZapLogger()
	logger.Warnw("I am testing credit logger.Warnw framework. ", map[string]interface{}{
		"key1": "value1",
		"key2": "value2",
		"key3": "value3",
	})
	logContext := logger_contetx.GetLogContext()
	funcContext := context.GetFuncContext(1)
	asserts.Equal(logContext.FuncContext.FileName, funcContext.FileName)
	asserts.Equal(logContext.FuncContext.FuncName, funcContext.FuncName)
	asserts.Equal(logContext.FuncContext.Line+6, funcContext.Line)

	//mustFalseFilter
	mustFalseFilter := log_filter.WrapLogFilterFunc(func(logContext *logger_contetx.LogContext) bool {
		return false
	})
	logger.WithFilter(mustFalseFilter)
	logger.Warnw("I am testing credit logger.Warnw framework. ", map[string]interface{}{
		"key1": "value1",
		"key2": "value2",
		"key3": "value3",
	})
	logContext = logger_contetx.GetLogContext()
	funcContext = context.GetFuncContext(1)
	asserts.Equal(logContext.FuncContext.FileName, funcContext.FileName)
	asserts.Equal(logContext.FuncContext.FuncName, funcContext.FuncName)
	asserts.Equal(logContext.FuncContext.Line+6, funcContext.Line)

	//filter false , monitor not nil
	mustTureFilter := log_filter.WrapLogFilterFunc(func(logContext *logger_contetx.LogContext) bool {
		return true
	})
	logger.WithFilter(mustTureFilter)
	logger.WithMonitor(func(logContext *logger_contetx.LogContext) {
		fmt.Println("Test logger.Monitor")
	})
	logger.Warnw("I am testing credit logger.Warnw framework. ", map[string]interface{}{
		"key1": "value1",
		"key2": "value2",
		"key3": "value3",
	})
	logContext = logger_contetx.GetLogContext()
	funcContext = context.GetFuncContext(1)
	asserts.Equal(logContext.FuncContext.FileName, funcContext.FileName)
	asserts.Equal(logContext.FuncContext.FuncName, funcContext.FuncName)
	asserts.Equal(logContext.FuncContext.Line+6, funcContext.Line)
}
