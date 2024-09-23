package test

import (
	. "github.com/go-tech-lab/framework/logger/src/log_context"
	"github.com/go-tech-lab/framework/logger/src/log_context/impl"
	_ "github.com/go-tech-lab/framework/logger/src/logger/impl"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestLoggerWithModuleTag_Error(t *testing.T) {
	asserts := require.New(t)
	logger := mockZapLogger()
	//测试默认的moduleTag
	logger.Error("I am testing credit logger.Error framework. ", "key1:", "value1", "key2:", "value2", "key3:", "value3")
	logContext := GetLogContext()
	asserts.Equal(impl.DefaultModuleTag, logContext.ModuleTag)

	//测试ModuleTag = ""
	logger.WithModuleTag("").Error("I am testing credit logger.Error framework. ", "key1:", "value1", "key2:", "value2", "key3:", "value3")
	logContext = GetLogContext()
	asserts.Equal("", logContext.ModuleTag)

	//测试自定义ModuleTag
	newModuleTag := "MyModuleTag"
	logger.WithModuleTag(newModuleTag).Error("I am testing credit logger.Error framework. ", "key1:", "value1", "key2:", "value2", "key3:", "value3")
	logContext = GetLogContext()
	asserts.Equal(newModuleTag, logContext.ModuleTag)

	//测试自定义ModuleTag回归到默认值
	logger.Error("I am testing credit logger.Error framework. ", "key1:", "value1", "key2:", "value2", "key3:", "value3")
	logContext = GetLogContext()
	asserts.Equal(impl.DefaultModuleTag, logContext.ModuleTag)
}

func TestLoggerWithDefaultTag_Errorf(t *testing.T) {
	asserts := require.New(t)
	logger := mockZapLogger()
	//测试默认的moduleTag
	logger.Errorf("I am testing credit logger.Errorf framework. key1:%s, key2:%s, key3:%s .", "value1", "value2", "value3")
	logContext := GetLogContext()
	asserts.Equal(impl.DefaultModuleTag, logContext.ModuleTag)

	//测试ModuleTag = ""
	logger.WithModuleTag("").Errorf("I am testing credit logger.Errorf framework. key1:%s, key2:%s, key3:%s .", "value1", "value2", "value3")
	logContext = GetLogContext()
	asserts.Equal("", logContext.ModuleTag)

	//测试自定义ModuleTag
	newModuleTag := "MyModuleTag"
	logger.WithModuleTag(newModuleTag).Errorf("I am testing credit logger.Errorf framework. key1:%s, key2:%s, key3:%s .", "value1", "value2", "value3")
	logContext = GetLogContext()
	asserts.Equal(newModuleTag, logContext.ModuleTag)

	//测试自定义ModuleTag回归到默认值
	logger.Errorf("I am testing credit logger.Errorf framework. key1:%s, key2:%s, key3:%s .", "value1", "value2", "value3")
	logContext = GetLogContext()
	asserts.Equal(impl.DefaultModuleTag, logContext.ModuleTag)
}

func TestLoggerWithDefaultTag_Errorw(t *testing.T) {
	asserts := require.New(t)
	logger := mockZapLogger()
	//测试默认的moduleTag
	logger.Errorw("I am testing credit logger.Errorw framework. ", map[string]interface{}{
		"key1": "value1",
		"key2": "value2",
		"key3": "value3",
	})
	logContext := GetLogContext()
	asserts.Equal(impl.DefaultModuleTag, logContext.ModuleTag)

	//测试ModuleTag = ""
	logger.WithModuleTag("").Errorw("I am testing credit logger.Errorw framework. ", map[string]interface{}{
		"key1": "value1",
		"key2": "value2",
		"key3": "value3",
	})
	logContext = GetLogContext()
	asserts.Equal("", logContext.ModuleTag)

	//测试自定义ModuleTag
	newModuleTag := "MyModuleTag"
	logger.WithModuleTag(newModuleTag).Errorw("I am testing credit logger.Errorw framework. ", map[string]interface{}{
		"key1": "value1",
		"key2": "value2",
		"key3": "value3",
	})
	logContext = GetLogContext()
	asserts.Equal(newModuleTag, logContext.ModuleTag)

	//测试自定义ModuleTag回归到默认值
	logger.Errorw("I am testing credit logger.Errorw framework. ", map[string]interface{}{
		"key1": "value1",
		"key2": "value2",
		"key3": "value3",
	})
	logContext = GetLogContext()
	asserts.Equal(impl.DefaultModuleTag, logContext.ModuleTag)
}
