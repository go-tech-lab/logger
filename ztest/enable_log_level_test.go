package test

import (
	"github.com/go-tech-lab/framework/logger/src/config"
	"github.com/go-tech-lab/framework/logger/src/level"
	"github.com/go-tech-lab/framework/logger/src/logger"
	impl2 "github.com/go-tech-lab/framework/logger/src/logger/impl"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestEnableLogLevel(t *testing.T) {
	zapConfig := &config.ZapLogConfig{
		OutputDir:      "/Users/xiaoxinchen/gopath/framework_logger/log",
		FileBase:       "all.log",
		TimeFormat:     "2006-01-02 15:04:05.000000",
		MaxAgeHour:     72,
		RotationHour:   1,
		FileTailFormat: ".%Y-%m-%d-%H",
		EnableDebug:    true,
		EnableLogLevel: "info",
	}
	zapLog := impl2.NewZapLogger(zapConfig)

	logOutput(t, zapLog, "debug", level.DebugLevel)
	logOutput(t, zapLog, "info", level.InfoLevel)
	logOutput(t, zapLog, "warn", level.WarnLevel)
	logOutput(t, zapLog, "error", level.ErrorLevel)

}

func logOutput(t *testing.T, logger logger.ILogger, logLevelConfig string, logLevel level.LogLevel) {
	asserts := require.New(t)

	logger.SetEnableLogLevel(logLevelConfig)
	asserts.Equal(logger.LogContextLocal().EnableLogLevel().String(), logLevelConfig)

	// Debug
	logger.Debug("test Debug log under enableLogLevel config is " + logLevel.String())
	asserts.Equal(logger.LogContextLocal().EnableLogLevel(), logLevel)

	logger.Debugf("test Debugf log under enableLogLevel config is %s,key:%s", logLevel.String(), "value1")
	asserts.Equal(logger.LogContextLocal().EnableLogLevel(), logLevel)

	logger.Debugw("test Debugw log under enableLogLevel config is ", map[string]interface{}{
		"logLevel": logLevel.String(),
		"key1":     "value1",
		"key2":     "value2",
		"key3":     "value3",
	})
	asserts.Equal(logger.LogContextLocal().EnableLogLevel(), logLevel)

	// Info
	logger.Info("test Info log under enableLogLevel config is " + logLevel.String())
	asserts.Equal(logger.LogContextLocal().EnableLogLevel(), logLevel)

	logger.Infof("test Infof log under enableLogLevel config is %s,key:%s", logLevel.String(), "value1")
	asserts.Equal(logger.LogContextLocal().EnableLogLevel(), logLevel)

	logger.Infow("test Infow log under enableLogLevel config is ", map[string]interface{}{
		"logLevel": logLevel.String(),
		"key1":     "value1",
		"key2":     "value2",
		"key3":     "value3",
	})
	asserts.Equal(logger.LogContextLocal().EnableLogLevel(), logLevel)

	// Warn
	logger.Warn("test Warn log under enableLogLevel config is " + logLevel.String())
	asserts.Equal(logger.LogContextLocal().EnableLogLevel(), logLevel)

	logger.Warnf("test Warnf log under enableLogLevel config is %s,key:%s", logLevel.String(), "value1")
	asserts.Equal(logger.LogContextLocal().EnableLogLevel(), logLevel)

	logger.Warnw("test Warnw log under enableLogLevel config is ", map[string]interface{}{
		"logLevel": logLevel.String(),
		"key1":     "value1",
		"key2":     "value2",
		"key3":     "value3",
	})
	asserts.Equal(logger.LogContextLocal().EnableLogLevel(), logLevel)

	// Error
	logger.Error("test Error log under enableLogLevel config is " + logLevel.String())
	asserts.Equal(logger.LogContextLocal().EnableLogLevel(), logLevel)

	logger.Errorf("test Errorf log under enableLogLevel config is %s,key:%s", logLevel.String(), "value1")
	asserts.Equal(logger.LogContextLocal().EnableLogLevel(), logLevel)

	logger.Errorw("test Errorw log under enableLogLevel config is ", map[string]interface{}{
		"logLevel": logLevel.String(),
		"key1":     "value1",
		"key2":     "value2",
		"key3":     "value3",
	})
	asserts.Equal(logger.LogContextLocal().EnableLogLevel(), logLevel)
}
