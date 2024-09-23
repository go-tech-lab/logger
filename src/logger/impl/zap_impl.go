package impl

import (
	"fmt"
	"github.com/go-tech-lab/framework/logger/src/level"
	logger_contetx "github.com/go-tech-lab/framework/logger/src/log_context"
	"github.com/go-tech-lab/framework/logger/src/log_filter"
	"github.com/go-tech-lab/framework/logger/src/logger"
	"github.com/go-tech-lab/framework_common/src/context"
	"github.com/go-tech-lab/framework_common/src/util"
	"go.uber.org/zap"
	"time"
)

type zapLogImpl struct {
	loggerInstance      *zap.SugaredLogger
	logFilter           log_filter.ILogFilter
	logFilterDecideFunc log_filter.LogFilterDecider
	logWrapper        logger.ILogWrapper
	logContextLocal     logger_contetx.ILogContextLocal
}

func (impl *zapLogImpl) Debugf(format string, params ...interface{}) {
	if !impl.filterByLogLevel(level.DebugLevel) {
		return
	}

	startTime := time.Now()
	if impl.logWrapper != nil {
		format, params = impl.logWrapper.LogfWrapper(format, params...)
	}

	logContext := impl.initLogContext("debug")
	defer func() {
		impl.takenTime(logContext, startTime)
		impl.clearLocal()
	}()
	if impl.decideLogFilter() && impl.logFilter != nil && !impl.logFilter.Filter(logContext) {
		logContext.FilterResult = false
		return
	}
	logPosition := impl.formatLogPosition(logContext)
	newFormat := logPosition + format
	impl.loggerInstance.Debugf(newFormat, params...)
	logContext.FilterResult = true
}

func (impl *zapLogImpl) Infof(format string, params ...interface{}) {
	if !impl.filterByLogLevel(level.InfoLevel) {
		return
	}

	startTime := time.Now()
	if impl.logWrapper != nil {
		format, params = impl.logWrapper.LogfWrapper(format, params...)
	}

	logContext := impl.initLogContext("info")
	defer func() {
		impl.takenTime(logContext, startTime)
		impl.clearLocal()
	}()
	if impl.decideLogFilter() && impl.logFilter != nil && !impl.logFilter.Filter(logContext) {
		logContext.FilterResult = false
		return
	}
	logPosition := impl.formatLogPosition(logContext)
	newFormat := logPosition + format
	impl.loggerInstance.Infof(newFormat, params...)
	logContext.FilterResult = true
}

func (impl *zapLogImpl) Warnf(format string, params ...interface{}) {
	if !impl.filterByLogLevel(level.WarnLevel) {
		return
	}

	startTime := time.Now()
	if impl.logWrapper != nil {
		format, params = impl.logWrapper.LogfWrapper(format, params...)
	}

	logContext := impl.initLogContext("warn")
	defer func() {
		impl.takenTime(logContext, startTime)
		impl.clearLocal()
	}()
	if impl.decideLogFilter() && impl.logFilter != nil && !impl.logFilter.Filter(logContext) {
		logContext.FilterResult = false
		return
	}
	logPosition := impl.formatLogPosition(logContext)
	newFormat := logPosition + format
	impl.loggerInstance.Warnf(newFormat, params...)
	logContext.FilterResult = true
}

func (impl *zapLogImpl) Errorf(format string, params ...interface{}) {
	if !impl.filterByLogLevel(level.ErrorLevel) {
		return
	}

	startTime := time.Now()
	if impl.logWrapper != nil {
		format, params = impl.logWrapper.LogfWrapper(format, params...)
	}

	logContext := impl.initLogContext("error")
	defer func() {
		impl.takenTime(logContext, startTime)
		impl.clearLocal()
	}()
	if impl.decideLogFilter() && impl.logFilter != nil && !impl.logFilter.Filter(logContext) {
		logContext.FilterResult = false
		return
	}
	logPosition := impl.formatLogPosition(logContext)
	newFormat := logPosition + format
	impl.loggerInstance.Errorf(newFormat, params...)
	logContext.FilterResult = true
}

func (impl *zapLogImpl) Debug(v ...interface{}) {
	if !impl.filterByLogLevel(level.DebugLevel) {
		return
	}

	startTime := time.Now()
	if impl.logWrapper != nil {
		v = impl.logWrapper.LogWrapper(v...)
	}

	logContext := impl.initLogContext("debug")
	defer func() {
		impl.takenTime(logContext, startTime)
		impl.clearLocal()
	}()
	if impl.decideLogFilter() && impl.logFilter != nil && !impl.logFilter.Filter(logContext) {
		logContext.FilterResult = false
		return
	}
	logPosition := impl.formatLogPosition(logContext)
	newV := append([]interface{}{logPosition}, v...)
	impl.loggerInstance.Debug(newLogMessage(newV).String())
	logContext.FilterResult = true
}

func (impl *zapLogImpl) Info(v ...interface{}) {
	if !impl.filterByLogLevel(level.InfoLevel) {
		return
	}

	startTime := time.Now()
	if impl.logWrapper != nil {
		v = impl.logWrapper.LogWrapper(v...)
	}

	logContext := impl.initLogContext("info")
	defer func() {
		impl.takenTime(logContext, startTime)
		impl.clearLocal()
	}()
	if impl.decideLogFilter() && impl.logFilter != nil && !impl.logFilter.Filter(logContext) {
		logContext.FilterResult = false
		return
	}
	logPosition := impl.formatLogPosition(logContext)
	newV := append([]interface{}{logPosition}, v...)
	impl.loggerInstance.Info(newV...)
	logContext.FilterResult = true
}

func (impl *zapLogImpl) Warn(v ...interface{}) {
	if !impl.filterByLogLevel(level.WarnLevel) {
		return
	}

	startTime := time.Now()
	if impl.logWrapper != nil {
		v = impl.logWrapper.LogWrapper(v...)
	}

	logContext := impl.initLogContext("warn")
	defer func() {
		impl.takenTime(logContext, startTime)
		impl.clearLocal()
	}()
	if impl.decideLogFilter() && impl.logFilter != nil && !impl.logFilter.Filter(logContext) {
		logContext.FilterResult = false
		return
	}
	logPosition := impl.formatLogPosition(logContext)
	newV := append([]interface{}{logPosition}, v...)
	impl.loggerInstance.Warn(newLogMessage(newV).String())
	logContext.FilterResult = true
}

func (impl *zapLogImpl) Error(v ...interface{}) {
	if !impl.filterByLogLevel(level.ErrorLevel) {
		return
	}

	startTime := time.Now()
	if impl.logWrapper != nil {
		v = impl.logWrapper.LogWrapper(v...)
	}

	logContext := impl.initLogContext("error")
	defer func() {
		impl.takenTime(logContext, startTime)
		impl.clearLocal()
	}()
	if impl.decideLogFilter() && impl.logFilter != nil && !impl.logFilter.Filter(logContext) {
		logContext.FilterResult = false
		return
	}
	logPosition := impl.formatLogPosition(logContext)
	newV := append([]interface{}{logPosition}, v...)
	impl.loggerInstance.Error(newLogMessage(newV).String())
	logContext.FilterResult = true
}

func (impl *zapLogImpl) Debugw(msg string, logFields logger.LogFields) {
	if !impl.filterByLogLevel(level.DebugLevel) {
		return
	}

	startTime := time.Now()
	if impl.logWrapper != nil {
		msg, logFields = impl.logWrapper.LogwWrapper(msg, logFields)
	}

	logContext := impl.initLogContext("debug")
	logContext.LogFields = logFields
	defer func() {
		impl.takenTime(logContext, startTime)
		impl.clearLocal()
	}()
	if impl.decideLogFilter() && impl.logFilter != nil && !impl.logFilter.Filter(logContext) {
		logContext.FilterResult = false
		return
	}
	logPosition := impl.formatLogPosition(logContext)
	newV := append([]interface{}{logPosition}, msg)
	newMsg := newLogMessage(newV).String()
	flatParams := util.FlatMap2Array(logFields)
	impl.loggerInstance.Debugw(newMsg, flatParams...)
	logContext.FilterResult = true
}

func (impl *zapLogImpl) Infow(msg string, logFields logger.LogFields) {
	if !impl.filterByLogLevel(level.InfoLevel) {
		return
	}

	startTime := time.Now()
	if impl.logWrapper != nil {
		msg, logFields = impl.logWrapper.LogwWrapper(msg, logFields)
	}

	logContext := impl.initLogContext("info")
	logContext.LogFields = logFields
	defer func() {
		impl.takenTime(logContext, startTime)
		impl.clearLocal()
	}()
	if impl.decideLogFilter() && impl.logFilter != nil && !impl.logFilter.Filter(logContext) {
		logContext.FilterResult = false
		return
	}
	logPosition := impl.formatLogPosition(logContext)
	newV := append([]interface{}{logPosition}, msg)
	newMsg := newLogMessage(newV).String()
	flatParams := util.FlatMap2Array(logFields)
	impl.loggerInstance.Infow(newMsg, flatParams...)
	logContext.FilterResult = true
}

func (impl *zapLogImpl) Warnw(msg string, logFields logger.LogFields) {
	if !impl.filterByLogLevel(level.WarnLevel) {
		return
	}

	startTime := time.Now()
	if impl.logWrapper != nil {
		msg, logFields = impl.logWrapper.LogwWrapper(msg, logFields)
	}

	logContext := impl.initLogContext("warn")
	logContext.LogFields = logFields
	defer func() {
		impl.takenTime(logContext, startTime)
		impl.clearLocal()
	}()
	if impl.decideLogFilter() && impl.logFilter != nil && !impl.logFilter.Filter(logContext) {
		logContext.FilterResult = false
		return
	}
	logPosition := impl.formatLogPosition(logContext)
	newV := append([]interface{}{logPosition}, msg)
	newMsg := newLogMessage(newV).String()
	flatParams := util.FlatMap2Array(logFields)
	impl.loggerInstance.Warnw(newMsg, flatParams...)
	logContext.FilterResult = true
}

func (impl *zapLogImpl) Errorw(msg string, logFields logger.LogFields) {
	if !impl.filterByLogLevel(level.ErrorLevel) {
		return
	}

	startTime := time.Now()
	if impl.logWrapper != nil {
		msg, logFields = impl.logWrapper.LogwWrapper(msg, logFields)
	}

	logContext := impl.initLogContext("error")
	logContext.LogFields = logFields
	defer func() {
		impl.takenTime(logContext, startTime)
		impl.clearLocal()
	}()
	if impl.decideLogFilter() && impl.logFilter != nil && !impl.logFilter.Filter(logContext) {
		logContext.FilterResult = false
		return
	}
	logPosition := impl.formatLogPosition(logContext)
	newV := append([]interface{}{logPosition}, msg)
	newMsg := newLogMessage(newV).String()
	flatParams := util.FlatMap2Array(logFields)
	impl.loggerInstance.Errorw(newMsg, flatParams...)
	logContext.FilterResult = true
}

func (impl *zapLogImpl) Filter() log_filter.ILogFilter {
	return impl.logFilter
}

func (impl *zapLogImpl) WithFilter(filter log_filter.ILogFilter) logger.ILogger {
	impl.logFilter = filter
	return impl
}

func (impl *zapLogImpl) WithModuleTag(moduleTag string) logger.ILogger {
	impl.logContextLocal.WithModuleTag(moduleTag)
	return impl
}

func (impl *zapLogImpl) SetDefaultModuleTag(defaultModuleTag string) logger.ILogger {
	impl.logContextLocal.SetDefaultModuleTag(defaultModuleTag)
	return impl
}

func (impl *zapLogImpl) WithMonitor(monitorFunc func(logContext *logger_contetx.LogContext)) logger.ILogger {
	return impl
}

func (impl *zapLogImpl) LogWrapper() logger.ILogWrapper {
	return impl.logWrapper
}

func (impl *zapLogImpl) WithLogWrapper(logDecorator logger.ILogWrapper) logger.ILogger {
	impl.logWrapper = logDecorator
	return impl
}

func (impl *zapLogImpl) LogFilterDecider() log_filter.LogFilterDecider {
	return impl.logFilterDecideFunc
}

func (impl *zapLogImpl) WithLogFilterDecider(logFilterDecider log_filter.LogFilterDecider) logger.ILogger {
	impl.logFilterDecideFunc = logFilterDecider
	return impl
}

func (impl *zapLogImpl) LogContextLocal() logger_contetx.ILogContextLocal {
	return impl.logContextLocal
}

func (impl *zapLogImpl) RawLogger() interface{} {
	return impl.loggerInstance
}

func (impl *zapLogImpl) initLogContext(level string, params ...map[string]interface{}) *logger_contetx.LogContext {
	skipDepth := impl.logContextLocal.SkipDepth()
	requestMethod := impl.logContextLocal.GetRequestMethod()
	funcContext := context.GetFuncContext(skipDepth)
	var logTags map[string]interface{} = nil
	if len(params) > 0 {
		logTags = params[0]
	}
	logContext := &logger_contetx.LogContext{
		ThroughFilterFlags: impl.logContextLocal.ThroughFilterFlags(),
		RequestMethod:      requestMethod,
		ModuleTag:          impl.logContextLocal.ModuleTag(),
		Level:              level,
		FuncContext: &context.FuncContext{
			FileName: funcContext.FileName,
			FuncName: funcContext.FuncName,
			Line:     funcContext.Line,
		},
		LogFields:    logTags,
		FilterResult: true,
		TakenTime:    0,
	}
	impl.logContextLocal.PutLogContext(logContext)
	return logContext
}

func (impl *zapLogImpl) clearLocal() {
	//impl.logContextLocal.Clear()
	impl.logContextLocal.ResetModuleTag()
}

func (impl *zapLogImpl) takenTime(logContext *logger_contetx.LogContext, startTime time.Time) {
	if logContext.FilterResult {
		endTime := time.Now()
		takenTime := endTime.Sub(startTime)
		logContext.TakenTime = takenTime
	}
}

func (impl *zapLogImpl) formatLogPosition(logContext *logger_contetx.LogContext) string {
	funcContext := logContext.FuncContext
	moduleTag := impl.logContextLocal.ModuleTag()
	logPosition := ""
	if len(moduleTag) > 0 {
		logPosition = fmt.Sprintf("[%s] [%s/%d: %s] ", moduleTag, funcContext.FileName, funcContext.Line, funcContext.FuncName)
	} else {
		logPosition = fmt.Sprintf("[%s/%d: %s] ", funcContext.FileName, funcContext.Line, funcContext.FuncName)
	}
	return logPosition
}

//日志过滤决策：针对一些特殊场景强制忽略日志过滤器使用
func (impl *zapLogImpl) decideLogFilter() bool {
	//有日志过滤决策函数，则让决策函数决策是否走日志过滤器
	if impl.logFilterDecideFunc != nil {
		return impl.logFilterDecideFunc()
	}
	//没有日志决策函数，默认是需要走日志过滤器的
	return true
}

func (impl *zapLogImpl) SetEnableLogLevel(enableLogLevel string) logger.ILogger {
	impl.logContextLocal.SetEnableLogLevel(enableLogLevel)
	return impl
}

//打印比local中存储的级别更高的日志
func (impl *zapLogImpl) filterByLogLevel(logLevel level.LogLevel) bool {
	enableLogLevel := impl.logContextLocal.EnableLogLevel()
	return enableLogLevel.Enabled(logLevel)
}
