package test

import (
	. "github.com/go-tech-lab/framework/logger/src/log_context"
	"github.com/go-tech-lab/framework/logger/src/log_context/impl"
	_ "github.com/go-tech-lab/framework/logger/src/logger/impl"
	_ "github.com/go-tech-lab/framework/logger/src/logger/impl"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestLoggerWithModuleTag_Warn(t *testing.T) {
	asserts := require.New(t)
	logger := mockZapLogger()
	//测试默认的moduleTag
	logger.Warn("I am testing credit logger.Warn framework. ", "key1:", "value1", "key2:", "value2", "key3:", "value3")
	logContext := GetLogContext()
	asserts.Equal(impl.DefaultModuleTag, logContext.ModuleTag)

	//测试ModuleTag = ""
	logger.WithModuleTag("").Warn("I am testing credit logger.Warn framework. ", "key1:", "value1", "key2:", "value2", "key3:", "value3")
	logContext = GetLogContext()
	asserts.Equal("", logContext.ModuleTag)

	//测试自定义ModuleTag
	newModuleTag := "MyModuleTag"
	logger.WithModuleTag(newModuleTag).Warn("I am testing credit logger.Warn framework. ", "key1:", "value1", "key2:", "value2", "key3:", "value3")
	logContext = GetLogContext()
	asserts.Equal(newModuleTag, logContext.ModuleTag)

	//测试自定义ModuleTag回归到默认值
	logger.Warn("I am testing credit logger.Warn framework. ", "key1:", "value1", "key2:", "value2", "key3:", "value3")
	logContext = GetLogContext()
	asserts.Equal(impl.DefaultModuleTag, logContext.ModuleTag)
}

func TestLoggerWithDefaultTag_Warnf(t *testing.T) {
	asserts := require.New(t)
	logger := mockZapLogger()
	//测试默认的moduleTag
	logger.Warnf("I am testing credit logger.Warnf framework. key1:%s, key2:%s, key3:%s .", "value1", "value2", "value3")
	logContext := GetLogContext()
	asserts.Equal(impl.DefaultModuleTag, logContext.ModuleTag)

	//测试ModuleTag = ""
	logger.WithModuleTag("").Warnf("I am testing credit logger.Warnf framework. key1:%s, key2:%s, key3:%s .", "value1", "value2", "value3")
	logContext = GetLogContext()
	asserts.Equal("", logContext.ModuleTag)

	//测试自定义ModuleTag
	newModuleTag := "MyModuleTag"
	logger.WithModuleTag(newModuleTag).Warnf("I am testing credit logger.Warnf framework. key1:%s, key2:%s, key3:%s .", "value1", "value2", "value3")
	logContext = GetLogContext()
	asserts.Equal(newModuleTag, logContext.ModuleTag)

	//测试自定义ModuleTag回归到默认值
	logger.Warnf("I am testing credit logger.Warnf framework. key1:%s, key2:%s, key3:%s .", "value1", "value2", "value3")
	logContext = GetLogContext()
	asserts.Equal(impl.DefaultModuleTag, logContext.ModuleTag)
}

func TestLoggerWithDefaultTag_Warnw(t *testing.T) {
	asserts := require.New(t)
	logger := mockZapLogger()
	//测试默认的moduleTag
	logger.Warnw("I am testing credit logger.Warnw framework. ", map[string]interface{}{
		"key1": "value1",
		"key2": "value2",
		"key3": "value3",
	})
	logContext := GetLogContext()
	asserts.Equal(impl.DefaultModuleTag, logContext.ModuleTag)

	//测试ModuleTag = ""
	logger.WithModuleTag("").Warnw("I am testing credit logger.Warnw framework. ", map[string]interface{}{
		"key1": "value1",
		"key2": "value2",
		"key3": "value3",
	})
	logContext = GetLogContext()
	asserts.Equal("", logContext.ModuleTag)

	//测试自定义ModuleTag
	newModuleTag := "MyModuleTag"
	logger.WithModuleTag(newModuleTag).Warnw("I am testing credit logger.Warnw framework. ", map[string]interface{}{
		"key1": "value1",
		"key2": "value2",
		"key3": "value3",
	})
	logContext = GetLogContext()
	asserts.Equal(newModuleTag, logContext.ModuleTag)

	//测试自定义ModuleTag回归到默认值
	logger.Warnw("I am testing credit logger.Warnw framework. ", map[string]interface{}{
		"key1": "value1",
		"key2": "value2",
		"key3": "value3",
	})
	logContext = GetLogContext()
	asserts.Equal(impl.DefaultModuleTag, logContext.ModuleTag)
}
