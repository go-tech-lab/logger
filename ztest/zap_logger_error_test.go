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

func TestLogger_Error(t *testing.T) {
	asserts := require.New(t)
	logger := mockZapLogger()
	logger.Error("I am testing credit logger.Error framework. ", "key1:", "value1", "key2:", "value2", "key3:", "value3")
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
	logger.Error("I am testing credit logger.Error framework. ", "key1:", "value1", "key2:", "value2", "key3:", "value3")
	logContext = logger_contetx.GetLogContext()
	funcContext = context.GetFuncContext(1)
	asserts.Equal(logContext.FuncContext.FileName, funcContext.FileName)
	asserts.Equal(logContext.FuncContext.FuncName, funcContext.FuncName)
	asserts.Equal(logContext.FuncContext.Line+2, funcContext.Line)

	//not filter
	logger.WithFilter(nil)
	logger.WithMonitor(nil)
	logger.Error("I am testing credit logger.Error framework. ", "key1:", "value1", "key2:", "value2", "key3:", "value3")
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
	logger.Error("I am testing credit logger.Error framework. ", "key1:", "value1", "key2:", "value2", "key3:", "value3")
	logContext = logger_contetx.GetLogContext()
	funcContext = context.GetFuncContext(1)
	asserts.Equal(logContext.FuncContext.FileName, funcContext.FileName)
	asserts.Equal(logContext.FuncContext.FuncName, funcContext.FuncName)
	asserts.Equal(logContext.FuncContext.Line+2, funcContext.Line)
}

func TestLogger_Errorf(t *testing.T) {
	asserts := require.New(t)
	logger := mockZapLogger()
	logger.Errorf("I am testing credit logger.Errorf framework. key1:%s, key2:%s, key3:%s .", "value1", "value2", "value3")
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
	logger.Errorf("I am testing credit logger.Errorf framework. key1:%s, key2:%s, key3:%s .", "value1", "value2", "value3")
	logContext = logger_contetx.GetLogContext()
	funcContext = context.GetFuncContext(1)
	asserts.Equal(logContext.FuncContext.FileName, funcContext.FileName)
	asserts.Equal(logContext.FuncContext.FuncName, funcContext.FuncName)
	asserts.Equal(logContext.FuncContext.Line+2, funcContext.Line)

	//not filter, no monitor
	logger.WithFilter(nil)
	logger.WithMonitor(nil)
	logger.Errorf("I am testing credit logger.Errorf framework. key1:%s, key2:%s, key3:%s .", "value1", "value2", "value3")
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
	logger.Errorf("I am testing credit logger.Errorf framework. key1:%s, key2:%s, key3:%s .", "value1", "value2", "value3")
	logContext = logger_contetx.GetLogContext()
	funcContext = context.GetFuncContext(1)
	asserts.Equal(logContext.FuncContext.FileName, funcContext.FileName)
	asserts.Equal(logContext.FuncContext.FuncName, funcContext.FuncName)
	asserts.Equal(logContext.FuncContext.Line+2, funcContext.Line)
}

func TestLogger_Errorw(t *testing.T) {
	asserts := require.New(t)
	logger := mockZapLogger()
	logger.Errorw("I am testing credit logger.Errorw framework. ", map[string]interface{}{
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
	logger.Errorw("I am testing credit logger.Errorw framework. ", map[string]interface{}{
		"key1": "value1",
		"key2": "value2",
		"key3": "value3",
	})
	logContext = logger_contetx.GetLogContext()
	funcContext = context.GetFuncContext(1)
	asserts.Equal(logContext.FuncContext.FileName, funcContext.FileName)
	asserts.Equal(logContext.FuncContext.FuncName, funcContext.FuncName)
	asserts.Equal(logContext.FuncContext.Line+6, funcContext.Line)

	//not filter, no monitor
	logger.WithFilter(nil)
	logger.Errorw("I am testing credit logger.Errorw framework. ", map[string]interface{}{
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
	logger.Errorw("I am testing credit logger.Errorw framework. ", map[string]interface{}{
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
