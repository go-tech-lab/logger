package impl

import (
	"github.com/go-tech-lab/framework/logger/src/level"
	. "github.com/go-tech-lab/framework/logger/src/log_context"
	"github.com/go-tech-lab/routinelocal/src"
)

func init() {
	routineFactory := routinelocal.RefRoutineLocalFactory()
	logContextLocal := &LogContextLocalImpl{
		defaultModuleTag:        DefaultModuleTag,
		enableLogLevel:          DefaultEnableLogLevel,
		requestMethodLocal:      routineFactory.NewRoutineLocal("requestMethodLocal"),
		moduleTagLocal:          routineFactory.NewRoutineLocal("moduleTagLocal"),
		logContextLocal:         routineFactory.NewRoutineLocal("logContextLocal"),
		logWrapDepthLocal:       routineFactory.NewRoutineLocal("logWrapDepthLocal"),
		throughFilterFlagsLocal: routineFactory.NewRoutineLocal("throughFilterFlagsLocal"),
	}
	InjectLogContextLocal(logContextLocal)
}

var DefaultSkipDepth = 3
var DefaultModuleTag = "Service"
var DefaultEnableLogLevel = level.InfoLevel

type LogContextLocalImpl struct {

	//默认ModuleTag
	defaultModuleTag string
	// 允许打印的日志等级
	enableLogLevel level.LogLevel
	//RoutineLocal requestMethodLocal
	requestMethodLocal routinelocal.IRoutineLocal
	//RoutineLocal moduleTag
	moduleTagLocal routinelocal.IRoutineLocal
	//RoutineLocal logContex
	logContextLocal routinelocal.IRoutineLocal
	//RoutineLocal logWrapDepth
	logWrapDepthLocal routinelocal.IRoutineLocal
	//RoutineLocal  throughFilterFlagsLocal
	throughFilterFlagsLocal routinelocal.IRoutineLocal
}

func (impl *LogContextLocalImpl) GetRequestMethod() string {
	ret := impl.requestMethodLocal.Value()
	if ret == nil {
		return ""
	}
	return impl.requestMethodLocal.Value().(string)
}

func (impl *LogContextLocalImpl) SkipDepth() int {
	ret := impl.logWrapDepthLocal.Value()
	if ret == nil {
		impl.logWrapDepthLocal.Put(DefaultSkipDepth)
		return DefaultSkipDepth
	}
	return impl.logWrapDepthLocal.Value().(int)
}

func (impl *LogContextLocalImpl) AddSkipDepth(addSkipDepth int) {
	logWrapDepth := impl.SkipDepth()
	logWrapDepth += addSkipDepth
	impl.logWrapDepthLocal.Put(logWrapDepth)
}

func (impl *LogContextLocalImpl) GetLogContext() *LogContext {
	ret := impl.logContextLocal.Value()
	if ret == nil {
		return nil
	}
	var logContext *LogContext = nil
	logContext = ret.(*LogContext)
	return logContext
}

func (impl *LogContextLocalImpl) PutLogContext(logContext *LogContext) {
	impl.logContextLocal.Put(logContext)
}

func (impl *LogContextLocalImpl) RequestMethod() string {
	ret := impl.requestMethodLocal.Value()
	if ret == nil {
		return ""
	}
	return impl.requestMethodLocal.Value().(string)
}

func (impl *LogContextLocalImpl) PutRequestMethod(requestMethod string) {
	impl.requestMethodLocal.Put(requestMethod)
}

func (impl *LogContextLocalImpl) ThroughFilterFlags() *ThroughFilterFlags {
	ret := impl.throughFilterFlagsLocal.Value()
	if ret == nil {
		return nil
	}
	return impl.throughFilterFlagsLocal.Value().(*ThroughFilterFlags)
}

func (impl *LogContextLocalImpl) PutThroughFilterFlags(flags *ThroughFilterFlags) {
	impl.throughFilterFlagsLocal.Put(flags)
}

func (impl *LogContextLocalImpl) ModuleTag() string {
	ret := impl.moduleTagLocal.Value()
	if ret == nil {
		impl.moduleTagLocal.Put(impl.defaultModuleTag)
		return impl.defaultModuleTag
	}
	return impl.moduleTagLocal.Value().(string)
}

func (impl *LogContextLocalImpl) WithModuleTag(moduleTag string) {
	impl.moduleTagLocal.Put(moduleTag)
}

func (impl *LogContextLocalImpl) ResetModuleTag() {
	impl.moduleTagLocal.Put(impl.defaultModuleTag)
}

func (impl *LogContextLocalImpl) SetDefaultModuleTag(moduleTag string) {
	impl.defaultModuleTag = moduleTag
}

func (impl *LogContextLocalImpl) Clear() {
	impl.logContextLocal.Remove()
	impl.logWrapDepthLocal.Remove()
	impl.moduleTagLocal.Remove()
}

func (impl *LogContextLocalImpl) SetEnableLogLevel(enableLogLevel string) {
	if enableLogLevel == "" {
		return
	}
	logLevel, err := level.Unmarshal(enableLogLevel)
	if err != nil {
		return
	}
	impl.enableLogLevel = logLevel
}

func (impl *LogContextLocalImpl) EnableLogLevel() level.LogLevel {
	return impl.enableLogLevel
}
