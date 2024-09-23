package impl

import (
	logger_contetx "github.com/go-tech-lab/framework/logger/src/log_context"
	"github.com/go-tech-lab/framework/logger/src/log_filter"
	"github.com/go-tech-lab/framework_common/src/context"
	"time"
)

//NewDefaultLogFilter 创建系统默认的日志过滤器
func NewDefaultLogFilter(logFilterRule *log_filter.LogFilterRule) log_filter.ILogFilter {
	return &logFilterImpl{
		logFilterRule: logFilterRule,
	}
}

//日志过滤器接口默认实现
type logFilterImpl struct {
	logFilterRule *log_filter.LogFilterRule
}

func (impl *logFilterImpl) Filter(logContext *logger_contetx.LogContext) bool {
	logContext.FilterType = logger_contetx.FilterTypeNone
	//全链路过滤只是针对info和debug
	throughFilterFlags := logContext.ThroughFilterFlags
	//全链路过滤只是针对info和debug
	if (logContext.Level == "info" || logContext.Level == "debug") && throughFilterFlags != nil {
		//全链路过滤决策
		disableLocalFilterFlag := throughFilterFlags.DisableLocalFilterFlag
		//禁用服务local过滤，有全链路过滤条件决策
		if disableLocalFilterFlag == "yes" {
			if throughFilterFlags.CutLogFlag == "yes" {
				logContext.FilterType = logger_contetx.FilterTypeFlow
				return false
			}
			if throughFilterFlags.CutLogFlag == "no" {
				return true
			}
		}
	}
	//没有过滤规则则会返回true,默认打印日志
	logFilterRule := impl.logFilterRule
	if logFilterRule != nil && len(logFilterRule.FilterRuleItems) > 0 {
		for _, rule := range logFilterRule.FilterRuleItems {
			//规则没有匹配上则忽略
			if !impl.match(logContext, rule) {
				continue
			}
			//0-不输出
			if rule.Rate <= 0 {
				logContext.FilterType = logger_contetx.FilterTypeLocal
				return false
			}
			//1-输出
			if rule.Rate >= 1 {
				return true
			}
			//概率输出
			randomRate := float64(time.Now().Nanosecond()%100) / 100.0
			if randomRate <= rule.Rate {
				return true
			}
			return false
		}
	}
	//没有服务local过滤规则，则使用全链路过滤规则
	if (logContext.Level == "info" || logContext.Level == "debug") && throughFilterFlags != nil {
		if throughFilterFlags.CutLogFlag == "yes" {
			logContext.FilterType = logger_contetx.FilterTypeFlow
			return false
		}
		if throughFilterFlags.CutLogFlag == "no" {
			return true
		}
	}
	//默认返回true
	return true
}

/***
 * 日志过滤器匹配
 * logContext 日志打印上下文
 * logFilterRule 日志过滤规则
 * 是否打印
 */
func (impl *logFilterImpl) match(logContext *logger_contetx.LogContext, logFilterRule *log_filter.LogFilterRuleItem) bool {
	ruleContext := logFilterRule.MatchContext
	//nil规则
	if ruleContext == nil {
		return false
	}
	if ruleContext.IsEmpty() {
		return false
	}

	if len(ruleContext.RequestMethod) > 0 && logContext.RequestMethod != ruleContext.RequestMethod {
		return false
	}

	if len(ruleContext.ModuleTag) > 0 && logContext.ModuleTag != ruleContext.ModuleTag {
		return false
	}

	if len(ruleContext.Level) > 0 && logContext.Level != ruleContext.Level {
		return false
	}
	ruleFuncContext := &context.FuncContext{
		FileName: ruleContext.FileName,
		FuncName: ruleContext.FuncName,
		Line:     ruleContext.Line,
	}
	if !impl.matchFuncContext(logContext.FuncContext, ruleFuncContext) {
		return false
	}
	if !impl.matchLogFields(logContext.LogFields, ruleContext.LogFields) {
		return false
	}
	return true
}

func (impl *logFilterImpl) matchFuncContext(src *context.FuncContext, rule *context.FuncContext) bool {
	if rule == nil {
		return true
	}
	if len(rule.FileName) > 0 && rule.FileName != src.FileName {
		return false
	}
	if len(rule.FuncName) > 0 && rule.FuncName != src.FuncName {
		return false
	}
	if rule.Line > 0 && rule.Line != src.Line {
		return false
	}
	return true
}

func (impl *logFilterImpl) matchLogFields(src map[string]interface{}, rule map[string]interface{}) bool {
	for key, value := range rule {
		if len(key) > 0 && src[key] != value {
			return false
		}
	}
	return true
}
