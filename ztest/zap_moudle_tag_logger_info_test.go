package test

import (
	. "github.com/go-tech-lab/framework/logger/src/log_context"
	"github.com/go-tech-lab/framework/logger/src/log_context/impl"
	_ "github.com/go-tech-lab/framework/logger/src/logger/impl"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestLoggerWithModuleTag_Info(t *testing.T) {
	asserts := require.New(t)
	logger := mockZapLogger()
	//测试默认的moduleTag
	logger.Info("I am testing credit logger.Info framework. ", "key1:", "value1", "key2:", "value2", "key3:", "value3")
	logContext := GetLogContext()
	asserts.Equal(impl.DefaultModuleTag, logContext.ModuleTag)

	//测试ModuleTag = ""
	logger.WithModuleTag("").Info("I am testing credit logger.Info framework. ", "key1:", "value1", "key2:", "value2", "key3:", "value3")
	logContext = GetLogContext()
	asserts.Equal("", logContext.ModuleTag)

	//测试自定义ModuleTag
	newModuleTag := "MyModuleTag"
	logger.WithModuleTag(newModuleTag).Info("I am testing credit logger.Info framework. ", "key1:", "value1", "key2:", "value2", "key3:", "value3")
	logContext = GetLogContext()
	asserts.Equal(newModuleTag, logContext.ModuleTag)

	//测试自定义ModuleTag回归到默认值
	logger.Info("I am testing credit logger.Info framework. ", "key1:", "value1", "key2:", "value2", "key3:", "value3")
	logContext = GetLogContext()
	asserts.Equal(impl.DefaultModuleTag, logContext.ModuleTag)
}

func TestLoggerWithDefaultTag_Infof(t *testing.T) {
	asserts := require.New(t)
	logger := mockZapLogger()
	//测试默认的moduleTag
	logger.Infof("I am testing credit logger.Infof framework. key1:%s, key2:%s, key3:%s .", "value1", "value2", "value3")
	logContext := GetLogContext()
	asserts.Equal(impl.DefaultModuleTag, logContext.ModuleTag)

	//测试ModuleTag = ""
	logger.WithModuleTag("").Infof("I am testing credit logger.Infof framework. key1:%s, key2:%s, key3:%s .", "value1", "value2", "value3")
	logContext = GetLogContext()
	asserts.Equal("", logContext.ModuleTag)

	//测试自定义ModuleTag
	newModuleTag := "MyModuleTag"
	logger.WithModuleTag(newModuleTag).Infof("I am testing credit logger.Infof framework. key1:%s, key2:%s, key3:%s .", "value1", "value2", "value3")
	logContext = GetLogContext()
	asserts.Equal(newModuleTag, logContext.ModuleTag)

	//测试自定义ModuleTag回归到默认值
	logger.Infof("I am testing credit logger.Infof framework. key1:%s, key2:%s, key3:%s .", "value1", "value2", "value3")
	logContext = GetLogContext()
	asserts.Equal(impl.DefaultModuleTag, logContext.ModuleTag)
}

func TestLoggerWithDefaultTag_Infow(t *testing.T) {
	asserts := require.New(t)
	logger := mockZapLogger()
	//测试默认的moduleTag
	logger.Infow("I am testing credit logger.Infow framework. ", map[string]interface{}{
		"key1": "value1",
		"key2": "value2",
		"key3": "value3",
	})
	logContext := GetLogContext()
	asserts.Equal(impl.DefaultModuleTag, logContext.ModuleTag)

	//测试ModuleTag = ""
	logger.WithModuleTag("").Infow("I am testing credit logger.Infow framework. ", map[string]interface{}{
		"key1": "value1",
		"key2": "value2",
		"key3": "value3",
	})
	logContext = GetLogContext()
	asserts.Equal("", logContext.ModuleTag)

	//测试自定义ModuleTag
	newModuleTag := "MyModuleTag"
	logger.WithModuleTag(newModuleTag).Infow("I am testing credit logger.Infow framework. ", map[string]interface{}{
		"key1": "value1",
		"key2": "value2",
		"key3": "value3",
	})
	logContext = GetLogContext()
	asserts.Equal(newModuleTag, logContext.ModuleTag)

	//测试自定义ModuleTag回归到默认值
	logger.Infow("I am testing credit logger.Infow framework. ", map[string]interface{}{
		"key1": "value1",
		"key2": "value2",
		"key3": "value3",
	})
	logContext = GetLogContext()
	asserts.Equal(impl.DefaultModuleTag, logContext.ModuleTag)
}
