package test

import (
	. "github.com/go-tech-lab/framework/logger/src/log_context"
	"github.com/go-tech-lab/framework/logger/src/log_context/impl"
	_ "github.com/go-tech-lab/framework/logger/src/logger/impl"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestLoggerWithModuleTag_Debug(t *testing.T) {
	asserts := require.New(t)
	logger := mockZapLogger()
	//测试默认的moduleTag
	logger.Debug("I am testing credit logger.Debug framework. ", "key1:", "value1", "key2:", "value2", "key3:", "value3")
	logContext := GetLogContext()
	asserts.Equal(impl.DefaultModuleTag, logContext.ModuleTag)

	//测试ModuleTag = ""
	logger.WithModuleTag("").Debug("I am testing credit logger.Debug framework. ", "key1:", "value1", "key2:", "value2", "key3:", "value3")
	logContext = GetLogContext()
	asserts.Equal("", logContext.ModuleTag)

	//测试自定义ModuleTag
	newModuleTag := "MyModuleTag"
	logger.WithModuleTag(newModuleTag).Debug("I am testing credit logger.Debug framework. ", "key1:", "value1", "key2:", "value2", "key3:", "value3")
	logContext = GetLogContext()
	asserts.Equal(newModuleTag, logContext.ModuleTag)

	//测试自定义ModuleTag回归到默认值
	logger.Debug("I am testing credit logger.Debug framework. ", "key1:", "value1", "key2:", "value2", "key3:", "value3")
	logContext = GetLogContext()
	asserts.Equal(impl.DefaultModuleTag, logContext.ModuleTag)
}

func TestLoggerWithDefaultTag_Debugf(t *testing.T) {
	asserts := require.New(t)
	logger := mockZapLogger()
	//测试默认的moduleTag
	logger.Debugf("I am testing credit logger.Debugf framework. key1:%s, key2:%s, key3:%s .", "value1", "value2", "value3")
	logContext := GetLogContext()
	asserts.Equal(impl.DefaultModuleTag, logContext.ModuleTag)

	//测试ModuleTag = ""
	logger.WithModuleTag("").Debugf("I am testing credit logger.Debugf framework. key1:%s, key2:%s, key3:%s .", "value1", "value2", "value3")
	logContext = GetLogContext()
	asserts.Equal("", logContext.ModuleTag)

	//测试自定义ModuleTag
	newModuleTag := "MyModuleTag"
	logger.WithModuleTag(newModuleTag).Debugf("I am testing credit logger.Debugf framework. key1:%s, key2:%s, key3:%s .", "value1", "value2", "value3")
	logContext = GetLogContext()
	asserts.Equal(newModuleTag, logContext.ModuleTag)

	//测试自定义ModuleTag回归到默认值
	logger.Debugf("I am testing credit logger.Debugf framework. key1:%s, key2:%s, key3:%s .", "value1", "value2", "value3")
	logContext = GetLogContext()
	asserts.Equal(impl.DefaultModuleTag, logContext.ModuleTag)
}

func TestLoggerWithDefaultTag_Debugw(t *testing.T) {
	asserts := require.New(t)
	logger := mockZapLogger()
	//测试默认的moduleTag
	logger.Debugw("I am testing credit logger.Debugw framework. ", map[string]interface{}{
		"key1": "value1",
		"key2": "value2",
		"key3": "value3",
	})
	logContext := GetLogContext()
	asserts.Equal(impl.DefaultModuleTag, logContext.ModuleTag)

	//测试ModuleTag = ""
	logger.WithModuleTag("").Debugw("I am testing credit logger.Debugw framework. ", map[string]interface{}{
		"key1": "value1",
		"key2": "value2",
		"key3": "value3",
	})
	logContext = GetLogContext()
	asserts.Equal("", logContext.ModuleTag)

	//测试自定义ModuleTag
	newModuleTag := "MyModuleTag"
	logger.WithModuleTag(newModuleTag).Debugw("I am testing credit logger.Debugw framework. ", map[string]interface{}{
		"key1": "value1",
		"key2": "value2",
		"key3": "value3",
	})
	logContext = GetLogContext()
	asserts.Equal(newModuleTag, logContext.ModuleTag)

	//测试自定义ModuleTag回归到默认值
	logger.Debugw("I am testing credit logger.Debugw framework. ", map[string]interface{}{
		"key1": "value1",
		"key2": "value2",
		"key3": "value3",
	})
	logContext = GetLogContext()
	asserts.Equal(impl.DefaultModuleTag, logContext.ModuleTag)
}
