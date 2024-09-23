package boot

import (
	"github.com/go-tech-lab/framework/logger/src/config"
	"github.com/go-tech-lab/framework/logger/src/level"
	"github.com/go-tech-lab/framework/logger/src/logger"
	"github.com/go-tech-lab/framework/logger/src/logger/impl"
)

//系统默认的logger
var defaultLogger logger.ILogger

//CreateLoggerFromFile 通过配置文件创建新的日志组件
//所谓的组件实际上就是对外暴露接口（当然要有实例）
func CreateLoggerFromFile(fileConfigPath string) logger.ILogger {
	zapLogConfig := loadZapLoggerConfig(fileConfigPath)
	newLogger := impl.NewZapLogger(zapLogConfig)
	newLogger.SetDefaultModuleTag("Service")
	newLogger.SetEnableLogLevel(getEnableLogLevel(zapLogConfig))
	//加载日志过滤器规则
	configFilter := LoadLogFilter(fileConfigPath)
	newLogger.WithFilter(configFilter)
	//注册日志过滤规则刷新任务
	registerFilterRefreshTask(newLogger, fileConfigPath)
	//防止无流量服务基础日志文件被归档
	registerPreventLogBaseFileArchiveTask(newLogger)
	if defaultLogger == nil {
		defaultLogger = newLogger
	}
	return newLogger
}

//DefaultLogger 获取默认的logger
func DefaultLogger() logger.ILogger {
	return defaultLogger
}

func ChangeDefaultLogger(logger logger.ILogger) {
	defaultLogger = logger
}

func getEnableLogLevel(config *config.ZapLogConfig) string {
	if config.EnableLogLevel != "" {
		return config.EnableLogLevel
	}
	if config.EnableDebug {
		return level.DebugLevel.String()
	}
	return level.InfoLevel.String()
}
