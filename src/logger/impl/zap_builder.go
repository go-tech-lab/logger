package impl

import (
	"fmt"
	"github.com/go-tech-lab/framework/logger/src/config"
	logger_contetx "github.com/go-tech-lab/framework/logger/src/log_context"
	"github.com/go-tech-lab/framework/logger/src/logger"
	trace "github.com/go-tech-lab/routinelocal/src/util"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"time"
)

var DefaultCapitalLevelEncoder = CapitalLevelEncoder

func NewZapLogger(zapConfig *config.ZapLogConfig) logger.ILogger {
	logContextLocal := logger_contetx.LogContextLocal()
	afterWriteFunc := func(p []byte, n int, err error) {
		logCtx := logContextLocal.GetLogContext()
		if logCtx != nil {
			logCtx.WriteMsgSize = n
		}
	}
	zapLogger := setupZapLogger(zapConfig, afterWriteFunc)
	loggerInstance := zapLogger.Sugar()
	logger := &zapLogImpl{
		loggerInstance:  loggerInstance,
		logContextLocal: logContextLocal,
	}
	return logger
}

func CapitalLevelEncoder(l zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(fmt.Sprintf("[%s] [%s]", trace.GetTraceId(), l.CapitalString()))
}

func setupZapLogger(config *config.ZapLogConfig, afterWrite logger.AfterWriteFunc) *zap.Logger {
	// 设置一些基本日志格式 具体含义还比较好理解，直接看zap源码也不难懂
	encoderConfig := zapcore.EncoderConfig{
		MessageKey: "msg",
		LevelKey:   "level",
		EncodeLevel: func(level zapcore.Level, encoder zapcore.PrimitiveArrayEncoder) {
			DefaultCapitalLevelEncoder(level, encoder)
		},
		TimeKey: "time",
		EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(t.Format(config.TimeFormat))
		},
		CallerKey:    "file",
		EncodeCaller: zapcore.ShortCallerEncoder,
		EncodeDuration: func(d time.Duration, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendInt64(int64(d) / 1000000)
		},
	}
	var encoder zapcore.Encoder = nil
	if config.LogFormat == "json" || config.LogFormat == "JSON" {
		encoder = zapcore.NewJSONEncoder(encoderConfig)
	} else {
		encoder = zapcore.NewConsoleEncoder(encoderConfig)
	}

	var hooks []zapcore.Core

	configHook := getHook(config, afterWrite)
	hooks = append(hooks, zapcore.NewCore(encoder, zapcore.AddSync(configHook), zapcore.DebugLevel))

	// 默认输出error级别的日志

	errorHook := getHookWithFileName(config, afterWrite, zapcore.ErrorLevel.String())
	hooks = append(hooks, zapcore.NewCore(encoder, zapcore.AddSync(errorHook), zapcore.ErrorLevel))

	core := zapcore.NewTee(hooks...)
	return zap.New(core) // 需要传入 zap.AddCaller() 才会显示打日志点的文件名和行数, 有点小坑
}
