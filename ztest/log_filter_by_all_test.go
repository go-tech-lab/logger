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
 *  精确输出完全满足（最全面，最精确）匹配规则
 *  1. 日志上下文:all与FilterRule:all完全匹配
 *  2. rate=1(包含)
 *  filter结果：true
 */
func TestLogFilter_By_Match_All_True(t *testing.T) {
	asserts := require.New(t)
	logFilter := log_filter.NewDefaultLogFilter(&log_filter2.LogFilterRule{
		FilterRuleItems: []*log_filter2.LogFilterRuleItem{
			{
				MatchContext: &log_filter2.LogAttribute{
					Level:    "info",
					FuncName: "FuncName1",
					FileName: "FileName1",
					Line:     10,
					LogFields: map[string]interface{}{
						"tag1": "tag1",
						"tag2": "tag2",
						"tag3": "tag3",
					},
				},
				Rate: 1},
		},
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
 *  精确排除不不完全满足（最全面，最精确）匹配规则
 *  1. 日志上下文:all与FilterRule:all 不完全匹配
 *  2. rate=1（排除）
 *  filter结果：false
 */
func TestLogFilter_By_All_NotMatch_True(t *testing.T) {
	asserts := require.New(t)
	logFilter := log_filter.NewDefaultLogFilter(&log_filter2.LogFilterRule{
		FilterRuleItems: []*log_filter2.LogFilterRuleItem{
			{
				MatchContext: &log_filter2.LogAttribute{
					Level:    "info",
					FuncName: "FuncName1",
					FileName: "FileName1",
					Line:     10,
					LogFields: map[string]interface{}{
						"tag1": "tag1",
						"tag2": "tag2",
						"tag3": "tagXXXXX",
					},
				},
				Rate: 0,
			},
		},
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
			"tag3": "tagXXXX", //本行不匹配
		},
	}
	result := logFilter.Filter(logContext)
	asserts.True(result, "")
}

/**
 *  精确排除完全满足（最全面，最精确）匹配规则
 *  1. 日志上下文:all与FilterRule:all完全匹配
 *  2. rate=0（排除）
 *  filter结果：false
 */
func TestLogFilter_By_All_Match_False(t *testing.T) {
	asserts := require.New(t)
	logFilter := log_filter.NewDefaultLogFilter(&log_filter2.LogFilterRule{
		FilterRuleItems: []*log_filter2.LogFilterRuleItem{
			{
				MatchContext: &log_filter2.LogAttribute{
					Level:    "info",
					FuncName: "FuncName1",
					FileName: "FileName1",
					Line:     10,
					LogFields: map[string]interface{}{
						"tag1": "tag1",
						"tag2": "tag2",
						"tag3": "tag3",
					},
				},
				Rate: 0, //排除
			},
		},
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
