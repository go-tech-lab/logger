package log_filter

//LogFilterRule 过滤规则
type LogFilterRule struct {
	//规则条目数组
	FilterRuleItems []*LogFilterRuleItem
}

type LogFilterRuleItem struct {
	//上下文匹配
	MatchContext *LogAttribute
	//概率0-1
	Rate float64
}

type ThroughFilterRule struct {
	RequestMethod      string
	Through            bool
	DisableLocalFilter bool
	Exceptive          map[string]string
}

//LogAttribute 日志过滤基本属性
type LogAttribute struct {
	//日志模块名
	ModuleTag string

	//日志级别
	Level string

	//请求方法（接口名）
	RequestMethod string

	//函数名
	FuncName string

	//文件名
	FileName string

	//代码行
	Line int

	//日志结构化字段
	LogFields map[string]interface{}
}

func (logMatchAttribute *LogAttribute) IsEmpty() bool {

	if len(logMatchAttribute.RequestMethod) > 0 {
		return false
	}
	if len(logMatchAttribute.ModuleTag) > 0 {
		return false
	}
	if len(logMatchAttribute.Level) > 0 {
		return false
	}
	if len(logMatchAttribute.LogFields) > 0 {
		return false
	}
	if len(logMatchAttribute.FileName) > 0 {
		return false
	}
	if len(logMatchAttribute.FuncName) > 0 {
		return false
	}
	if logMatchAttribute.Line > 0 {
		return false
	}
	return true
}
