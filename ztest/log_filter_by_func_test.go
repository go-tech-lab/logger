package test

import (
	logger_contetx "github.com/go-tech-lab/framework/logger/src/log_context"
	log_filter2 "github.com/go-tech-lab/framework/logger/src/log_filter"
	log_filter "github.com/go-tech-lab/framework/logger/src/log_filter/impl"
	_ "github.com/go-tech-lab/framework/logger/src/logger/impl"
	"github.com/go-tech-lab/framework_common/src/context"
	"github.com/stretchr/testify/require"
	"testing"
)

/**
 *  精确输出某个文件某行日志
 *  1. 日志上下文与FilterRule FuncContext 匹配
 *  2. rate=1
 *  filter结果：true
 */
func TestLogFilter_By_FuncContext_True(t *testing.T) {
	asserts := require.New(t)
	logFilter := log_filter.NewDefaultLogFilter(&log_filter2.LogFilterRule{
		FilterRuleItems: []*log_filter2.LogFilterRuleItem{
			{
				MatchContext: &log_filter2.LogAttribute{
					FuncName: "FuncName1",
					FileName: "FileName1",
					Line:     10,
				},

				Rate: 1,
			}},
	})
	logContext := &logger_contetx.LogContext{
		Level: "info",
		FuncContext: &context.FuncContext{
			FuncName: "FuncName1",
			FileName: "FileName1",
			Line:     10,
		},
		LogFields: map[string]interface{}{
			"tag1": "tag1",
			"tag2": "tag2",
			"tag3": "tag3",
		},
	}
	result := logFilter.Filter(logContext)
	asserts.True(result, "")
}

/**
 *  精确排除某个文件某行日志
 *  1. 日志上下文与FilterRule FuncContext 匹配
 *  2. rate=1
 *  filter结果：true
 */
func TestLogFilter_By_FuncContext_False(t *testing.T) {
	asserts := require.New(t)
	logFilter := log_filter.NewDefaultLogFilter(&log_filter2.LogFilterRule{
		FilterRuleItems: []*log_filter2.LogFilterRuleItem{
			{
				MatchContext: &log_filter2.LogAttribute{
					FuncName: "FuncName1",
					FileName: "FileName1",
					Line:     10,
				},
				Rate: 0,
			}},
	})
	logContext := &logger_contetx.LogContext{
		Level: "info",
		FuncContext: &context.FuncContext{
			FuncName: "FuncName1",
			FileName: "FileName1",
			Line:     10,
		},
		LogFields: map[string]interface{}{
			"tag1": "tag1",
			"tag2": "tag2",
			"tag3": "tag3",
		},
	}
	result := logFilter.Filter(logContext)
	asserts.False(result, "")
}

/**
 *  精确排除某个文件某行日志
 *  1. 日志上下文与FilterRule FuncContext 不匹配
 *  2. rate=1
 *  filter结果：true
 */
func TestLogFilter_By_FuncContext_NotMatch_True1(t *testing.T) {
	asserts := require.New(t)
	logFilter := log_filter.NewDefaultLogFilter(&log_filter2.LogFilterRule{
		FilterRuleItems: []*log_filter2.LogFilterRuleItem{
			{
				MatchContext: &log_filter2.LogAttribute{
					FuncName: "FuncName1",
					FileName: "FileName1",
					Line:     10,
				},
				Rate: 0,
			}},
	})
	logContext := &logger_contetx.LogContext{
		Level: "info",
		FuncContext: &context.FuncContext{
			FuncName: "FuncName2", //本行不匹配
			FileName: "FileName1",
			Line:     10,
		},
		LogFields: map[string]interface{}{
			"tag1": "tag1",
			"tag2": "tag2",
			"tag3": "tag3",
		},
	}
	result := logFilter.Filter(logContext)
	asserts.True(result, "")
}

/**
 *  精确排除某个文件某行日志
 *  1. 日志上下文与FilterRule FuncContext 不匹配
 *  2. rate=1
 *  filter结果：true
 */
func TestLogFilter_By_FuncContext_NotMatch_True2(t *testing.T) {
	asserts := require.New(t)
	logFilter := log_filter.NewDefaultLogFilter(&log_filter2.LogFilterRule{
		FilterRuleItems: []*log_filter2.LogFilterRuleItem{
			{
				MatchContext: &log_filter2.LogAttribute{
					FuncName: "FuncName1",
					FileName: "FileName1",
					Line:     10,
				},
				Rate: 0,
			}},
	})
	logContext := &logger_contetx.LogContext{
		Level: "info",
		FuncContext: &context.FuncContext{
			FuncName: "FuncName1",
			FileName: "FileName2", //本行不匹配
			Line:     10,
		},
		LogFields: map[string]interface{}{
			"tag1": "tag1",
			"tag2": "tag2",
			"tag3": "tag3",
		},
	}
	result := logFilter.Filter(logContext)
	asserts.True(result, "")
}

/**
 *  精确排除某个文件某行日志
 *  1. 日志上下文与FilterRule FuncContext 不匹配
 *  2. rate=1
 *  filter结果：true
 */
func TestLogFilter_By_FuncContext_NotMatch_True3(t *testing.T) {
	asserts := require.New(t)
	logFilter := log_filter.NewDefaultLogFilter(&log_filter2.LogFilterRule{
		FilterRuleItems: []*log_filter2.LogFilterRuleItem{
			{
				MatchContext: &log_filter2.LogAttribute{
					FuncName: "FuncName1",
					FileName: "FileName1",
					Line:     10,
				},
				Rate: 0,
			}},
	})
	logContext := &logger_contetx.LogContext{
		Level: "info",
		FuncContext: &context.FuncContext{
			FuncName: "FuncName1",
			FileName: "FileName2", //本行不匹配
			Line:     100,
		},
		LogFields: map[string]interface{}{
			"tag1": "tag1",
			"tag2": "tag2",
			"tag3": "tag3",
		},
	}
	result := logFilter.Filter(logContext)
	asserts.True(result, "")
}
