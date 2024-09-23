package logger_contetx

import (
	"github.com/go-tech-lab/framework_common/src/context"
	"time"
)

//LogContext 日志上下文
type LogContext struct {
	ThroughFilterFlags *ThroughFilterFlags

	//日志模块名
	ModuleTag string

	//日志级别
	Level string

	//请求方法（RPC接口名，MQ topic名称）
	RequestMethod string

	//函数context
	FuncContext *context.FuncContext

	//日志结构化字段
	LogFields map[string]interface{}

	//过滤类型： none,flow,local
	FilterType string

	//是否被输出,true 需要输出日志
	FilterResult bool

	//输出消息字节数
	WriteMsgSize int

	//耗时
	TakenTime time.Duration
}

// ThroughFilterFlags 全链路日志过滤标签
type ThroughFilterFlags struct {
	//剪除日志标签(yes,no)
	CutLogFlag string
	//禁用服务local过滤条件(yes,no)
	DisableLocalFilterFlag string
}

const (
	//FilterTypeNone none
	FilterTypeNone = "none"
	//FilterTypeFlow 全链路过滤
	FilterTypeFlow = "flow"
	//FilterTypeLocal 本地规则过滤
	FilterTypeLocal = "local"

	ThroughCutLogFlagKey = "x-log-cut-flag"
	ThroughCutLogFlagYES = "yes"
	ThroughCutLogFlagNO  = "no"

	ThroughDisableLocalLogFilterFlagKey = "x-log-disable-local-filter-flag"
	ThroughDisableLocalLogFilterFlagYES = "yes"
	ThroughDisableLocalLogFilterFlagNO  = "no"
)
