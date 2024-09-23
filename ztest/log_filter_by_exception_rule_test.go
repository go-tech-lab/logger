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
func TestLogFilter_Nil_Rule(t *testing.T) {
	asserts := require.New(t)

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

	logFilter := log_filter.NewDefaultLogFilter(&log_filter2.LogFilterRule{
		FilterRuleItems: []*log_filter2.LogFilterRuleItem{
			{}}})
	result := logFilter.Filter(logContext)
	asserts.True(result, "")

	logFilter = log_filter.NewDefaultLogFilter(nil)
	result = logFilter.Filter(logContext)
	asserts.True(result, "")
}
