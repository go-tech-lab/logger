package impl

import (
	"fmt"
	"github.com/arthurkiller/rollingwriter"
	"github.com/go-tech-lab/framework/logger/src/config"
	"github.com/go-tech-lab/framework/logger/src/logger"
	"github.com/go-tech-lab/framework_common/src/util"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"io"
	"path/filepath"
	"time"
)

type logMessage struct {
	params []interface{}
}

func (message *logMessage) String() string {
	return fmt.Sprint(message.params...)
}

func newLogMessage(params []interface{}) fmt.Stringer {
	message := new(logMessage)
	message.params = params
	return message
}

func getLogOutFilePath(dir string, name string) string {
	if len(name) <= 0 {
		return ""
	}
	return filepath.Join(dir, name)
}

func getHookWithFileName(logConfig *config.ZapLogConfig, afterWrite logger.AfterWriteFunc, fileName string) io.Writer {
	newLogConfig := &config.ZapLogConfig{
		OutputDir:          logConfig.OutputDir,
		TimeFormat:         logConfig.TimeFormat,
		MaxAgeHour:         logConfig.MaxAgeHour,
		RotationHour:       logConfig.RotationHour,
		FileBase:           fileName,
		FileTailFormat:     logConfig.FileTailFormat,
		UseLogPlatform:     logConfig.UseLogPlatform,
		WriteMode:          logConfig.WriteMode,
		TimeRollingPattern: logConfig.TimeRollingPattern,
	}
	return getHook(newLogConfig, afterWrite)
}

func getHook(config *config.ZapLogConfig, afterFunc logger.AfterWriteFunc) io.Writer {
	var writer io.Writer = nil
	if config == nil || len(config.FileBase) <= 0 {
		writer = &stdoutHook{}
	} else {
		if config.UseLogPlatform {
			writer = getLogPlatformWriter(config)
		} else {
			writer = getWriter(config)
		}
	}
	return &writerWrapper{
		writer:    writer,
		afterFunc: afterFunc,
	}
}

func getWriter(config *config.ZapLogConfig) io.Writer {
	absOutputDir, err := filepath.Abs(config.OutputDir)
	util.CheckErr(err)
	logFileBase := getLogOutFilePath(absOutputDir, config.FileBase)
	softLinkFilePath := filepath.Join(absOutputDir, config.FileBase)
	// 生成rotatelogs的Logger 实际生成的文件名 XX.log.YYmmddHH
	// 保存MaxAge(小时)内的日志，每RotationTime(整点)分割一次日志
	writer, err := rotatelogs.New(
		logFileBase+config.FileTailFormat, // 没有使用go风格反人类的format格式
		rotatelogs.WithLinkName(softLinkFilePath),
		rotatelogs.WithMaxAge(time.Hour*time.Duration(config.MaxAgeHour)),
		rotatelogs.WithRotationTime(time.Hour*time.Duration(config.RotationHour)),
	)
	util.CheckErr(err)

	return writer
}

func getLogPlatformWriter(config *config.ZapLogConfig) io.Writer {
	absOutputDir, err := filepath.Abs(config.OutputDir)
	util.CheckErr(err)
	writeConfig := rollingwriter.Config{
		LogPath:       absOutputDir,          //日志路径
		TimeTagFormat: config.FileTailFormat, //时间格式串
		FileName:      config.FileBase,       //日志文件名
		MaxRemain:     config.MaxAgeHour,     //配置日志最大存留数
		// 目前有2中滚动策略: 按照时间滚动按照大小滚动
		// - 时间滚动: 配置策略如同 crontable, 例如,每天0:0切分, 则配置 0 0 0 * * *
		// - 大小滚动: 配置单个日志文件(未压缩)的滚动大小门限, 入1G, 500M
		RollingPolicy:      rollingwriter.TimeRolling, //配置滚动策略 norolling timerolling volumerolling
		RollingTimePattern: config.TimeRollingPattern, //配置时间滚动策略
		//RollingVolumeSize:  "2k",                      //配置截断文件下限大小

		// writer 支持4种不同的 mode:
		// 1. none 2. lock
		// 3. async 4. buffer
		// - 无保护的 writer: 不提供并发安全保障
		// - lock 保护的 writer: 提供由 mutex 保护的并发安全保障
		// - 异步 writer: 异步 write, 并发安全. 异步开启后忽略 Lock 选项
		WriterMode: config.WriteMode,
		// BufferWriterThershould in B
		//BufferWriterThershould: 8 * 1024 * 1024,
		// Compress will compress log file with gzip
		Compress: false,
	}

	// 创建一个 writer
	writer, err := rollingwriter.NewWriterFromConfig(&writeConfig)
	util.CheckErr(err)

	return writer
}

type stdoutHook struct {
}

func (hook *stdoutHook) Write(p []byte) (n int, err error) {
	fmt.Println(string(p))
	return len(p), nil
}

type writerWrapper struct {
	writer    io.Writer
	afterFunc logger.AfterWriteFunc
}

func (wrapper *writerWrapper) Write(p []byte) (n int, err error) {
	n, err = wrapper.writer.Write(p)
	if wrapper.afterFunc != nil {
		wrapper.afterFunc(p, n, err)
	}
	return
}
