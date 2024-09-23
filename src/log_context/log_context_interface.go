package logger_contetx

import "github.com/go-tech-lab/framework/logger/src/level"

type ILogContextLocal interface {

	//GetLogContext 获取当前日志上下文
	GetLogContext() *LogContext

	//PutLogContext 设置当前日志上下文
	PutLogContext(logContext *LogContext)

	//GetRequestMethod get RequestMethod
	GetRequestMethod() string

	//PutRequestMethod 设置RequestMethod
	PutRequestMethod(requestMethod string)

	//ThroughFilterFlags get ThroughFilterFlags
	ThroughFilterFlags() *ThroughFilterFlags

	//PutThroughFilterFlags 设置ThroughFilterFlags
	PutThroughFilterFlags(flags *ThroughFilterFlags)

	//SkipDepth log caller 函数栈深度
	SkipDepth() int

	//AddSkipDepth 调整log caller 函数栈深度
	AddSkipDepth(addSkipDepth int)

	//ModuleTag 获取RoutineLocal日志模块tag
	ModuleTag() string

	//WithModuleTag 修改RoutineLocal日志模块tag
	WithModuleTag(moduleTag string)

	//ResetModuleTag 重制日志模块tag为默认值DefaultModuleTag
	ResetModuleTag()

	//SetDefaultModuleTag 设置默认你日志模块tag
	SetDefaultModuleTag(moduleTag string)

	//SetEnableLogLevel 设置允许打印的日志等级
	SetEnableLogLevel(logLevel string)

	//EnableLogLevel 获取日志打印级别
	EnableLogLevel() level.LogLevel

	//Clear 清理RoutineLocal，防止内存泄漏
	Clear()
}

var logContextLocal ILogContextLocal

func LogContextLocal() ILogContextLocal {
	return logContextLocal
}

func InjectLogContextLocal(_logContextLocal ILogContextLocal) {
	logContextLocal = _logContextLocal
}

func GetLogContext() *LogContext {
	return LogContextLocal().GetLogContext()
}
