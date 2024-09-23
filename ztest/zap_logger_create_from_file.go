package test

import (
	"fmt"
	logger_contetx "github.com/go-tech-lab/framework/logger/src/log_context"
	. "github.com/go-tech-lab/framework/logger/src/log_filter"
	_ "github.com/go-tech-lab/framework/logger/src/logger/impl"
	"github.com/go-tech-lab/framework_common/src/context"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestLogger_Info(t *testing.T) {
	asserts := require.New(t)
	logger := mockZapLogger()
	logger.Info("I am testing credit logger.Info framework. ", "key1:", "value1", "key2:", "value2", "key3:", "value3")
	logContext := logger_contetx.GetLogContext()
	funcContext := context.GetFuncContext(1)
	asserts.Equal(logContext.FuncContext.FileName, funcContext.FileName)
	asserts.Equal(logContext.FuncContext.FuncName, funcContext.FuncName)
	asserts.Equal(logContext.FuncContext.Line+2, funcContext.Line)

	//mustFalseFilter
	mustFalseFilter := WrapLogFilterFunc(func(logContext *logger_contetx.LogContext) bool {
		return false
	})
	logger.WithFilter(mustFalseFilter)
	logger.Info("I am testing credit logger.Info framework. ", "key1:", "value1", "key2:", "value2", "key3:", "value3")
	logContext = logger_contetx.GetLogContext()
	funcContext = context.GetFuncContext(1)
	asserts.Equal(logContext.FuncContext.FileName, funcContext.FileName)
	asserts.Equal(logContext.FuncContext.FuncName, funcContext.FuncName)
	asserts.Equal(logContext.FuncContext.Line+2, funcContext.Line)

	//not filter
	logger.WithFilter(nil)
	logger.Info("I am testing credit logger.Info framework. ", "key1:", "value1", "key2:", "value2", "key3:", "value3")
	logContext = logger_contetx.GetLogContext()
	funcContext = context.GetFuncContext(1)
	asserts.Equal(logContext.FuncContext.FileName, funcContext.FileName)
	asserts.Equal(logContext.FuncContext.FuncName, funcContext.FuncName)
	asserts.Equal(logContext.FuncContext.Line+2, funcContext.Line)

	//filter false , monitor not nil
	mustTureFilter := WrapLogFilterFunc(func(logContext *logger_contetx.LogContext) bool {
		return true
	})
	logger.WithFilter(mustTureFilter)
	logger.WithMonitor(func(logContext *logger_contetx.LogContext) {
		fmt.Println("Test logger.Monitor")
	})
	logger.Info("I am testing credit logger.Info framework. ", "key1:", "value1", "key2:", "value2", "key3:", "value3")
	logContext = logger_contetx.GetLogContext()
	funcContext = context.GetFuncContext(1)
	asserts.Equal(logContext.FuncContext.FileName, funcContext.FileName)
	asserts.Equal(logContext.FuncContext.FuncName, funcContext.FuncName)
	asserts.Equal(logContext.FuncContext.Line+2, funcContext.Line)
}

func TestLogger_Infof(t *testing.T) {
	asserts := require.New(t)
	logger := mockZapLogger()
	logger.Infof("I am testing credit logger.Infof framework. key1:%s, key2:%s, key3:%s .", "value1", "value2", "value3")
	logContext := logger_contetx.GetLogContext()
	funcContext := context.GetFuncContext(1)
	asserts.Equal(logContext.FuncContext.FileName, funcContext.FileName)
	asserts.Equal(logContext.FuncContext.FuncName, funcContext.FuncName)
	asserts.Equal(logContext.FuncContext.Line+2, funcContext.Line)

	//mustFalseFilter
	mustFalseFilter := WrapLogFilterFunc(func(logContext *logger_contetx.LogContext) bool {
		return false
	})
	logger.WithFilter(mustFalseFilter)
	logger.Infof("I am testing credit logger.Infof framework. key1:%s, key2:%s, key3:%s .", "value1", "value2", "value3")
	logContext = logger_contetx.GetLogContext()
	funcContext = context.GetFuncContext(1)
	asserts.Equal(logContext.FuncContext.FileName, funcContext.FileName)
	asserts.Equal(logContext.FuncContext.FuncName, funcContext.FuncName)
	asserts.Equal(logContext.FuncContext.Line+2, funcContext.Line)

	//not filter
	logger.WithFilter(nil)
	logger.Infof("I am testing credit logger.Infof framework. key1:%s, key2:%s, key3:%s .", "value1", "value2", "value3")
	logContext = logger_contetx.GetLogContext()
	funcContext = context.GetFuncContext(1)
	asserts.Equal(logContext.FuncContext.FileName, funcContext.FileName)
	asserts.Equal(logContext.FuncContext.FuncName, funcContext.FuncName)
	asserts.Equal(logContext.FuncContext.Line+2, funcContext.Line)

	//filter false , monitor not nil
	mustTureFilter :=WrapLogFilterFunc(func(logContext *logger_contetx.LogContext) bool {
		return true
	})
	logger.WithFilter(mustTureFilter)
	logger.WithMonitor(func(logContext *logger_contetx.LogContext) {
		fmt.Println("Test logger.Monitor")
	})
	logger.Infof("I am testing credit logger.Infof framework. key1:%s, key2:%s, key3:%s .", "value1", "value2", "value3")
	logContext = logger_contetx.GetLogContext()
	funcContext = context.GetFuncContext(1)
	asserts.Equal(logContext.FuncContext.FileName, funcContext.FileName)
	asserts.Equal(logContext.FuncContext.FuncName, funcContext.FuncName)
	asserts.Equal(logContext.FuncContext.Line+2, funcContext.Line)
}

func TestLogger_Infow(t *testing.T) {
	asserts := require.New(t)
	logger := mockZapLogger()
	logger.Infow("I am testing credit logger.Infow framework. ", map[string]interface{}{
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
	mustFalseFilter := WrapLogFilterFunc(func(logContext *logger_contetx.LogContext) bool {
		return false
	})
	logger.WithFilter(mustFalseFilter)
	logger.Infow("I am testing credit logger.Infow framework. ", map[string]interface{}{
		"key1": "value1",
		"key2": "value2",
		"key3": "value3",
	})
	logContext = logger_contetx.GetLogContext()
	funcContext = context.GetFuncContext(1)
	asserts.Equal(logContext.FuncContext.FileName, funcContext.FileName)
	asserts.Equal(logContext.FuncContext.FuncName, funcContext.FuncName)
	asserts.Equal(logContext.FuncContext.Line+6, funcContext.Line)

	//not filter
	logger.WithFilter(nil)
	logger.Infow("I am testing credit logger.Infow framework. ", map[string]interface{}{
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
	mustTureFilter := WrapLogFilterFunc(func(logContext *logger_contetx.LogContext) bool {
		return true
	})
	logger.WithFilter(mustTureFilter)
	logger.WithMonitor(func(logContext *logger_contetx.LogContext) {
		fmt.Println("Test logger.Monitor")
	})
	logger.Infow("I am testing credit logger.Infow framework. ", map[string]interface{}{
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
