package logger

type LogFieldsGetter func() LogFields

//ILogWrapper 日志内容包装器
//通常用于从应用之外修改，调整，丰富日志内容，比如全链路日志携带特定标签
type ILogWrapper interface {

	//LogfWrapper 应用于***f版本的输出日志函数
	LogfWrapper(format string, params ...interface{}) (string, []interface{})

	//LogWrapper 应用于普通版本输出日志函数
	LogWrapper(v ...interface{}) []interface{}

	//LogwWrapper 应用于***w版本的输出日志函数
	LogwWrapper(msg string, logFiled LogFields) (string, LogFields)
}
