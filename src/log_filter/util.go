package log_filter

import . "github.com/go-tech-lab/framework/logger/src/log_context"

//WrapLogFilterFunc 将logFilterFunc包装成ILogFilter，方便普通函数到ILogFilter的转换
func WrapLogFilterFunc(logFilterFunc LogFilterFunc) ILogFilter {
	return &logFilterWrapper{
		logFilterFunc: logFilterFunc,
	}
}

type LogFilterFunc func(logContext *LogContext) bool

type logFilterWrapper struct {
	logFilterFunc LogFilterFunc
}

func (wrapper *logFilterWrapper) Filter(logContext *LogContext) bool {
	return wrapper.logFilterFunc(logContext)
}
