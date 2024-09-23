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

func TestLogFilter_By_LogTag_True(t *testing.T) {
	asserts := require.New(t)
	logFilter := log_filter.NewDefaultLogFilter(&log_filter2.LogFilterRule{
		FilterRuleItems: []*log_filter2.LogFilterRuleItem{
			{
				MatchContext: &log_filter2.LogAttribute{
					LogFields: map[string]interface{}{
						"tag1": "tag1",
						"tag2": "tag2",
						"tag3": "tag3",
					},
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

func TestLogFilter_By_LogTag_False(t *testing.T) {
	asserts := require.New(t)
	logFilter := log_filter.NewDefaultLogFilter(&log_filter2.LogFilterRule{
		FilterRuleItems: []*log_filter2.LogFilterRuleItem{
			{
				MatchContext: &log_filter2.LogAttribute{
					LogFields: map[string]interface{}{
						"tag1": "tag1",
						"tag2": "tag2",
						"tag3": "tag3",
					},
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

func TestLogFilter_By_LogTag_NotMatch(t *testing.T) {
	asserts := require.New(t)
	logFilter := log_filter.NewDefaultLogFilter(&log_filter2.LogFilterRule{
		FilterRuleItems: []*log_filter2.LogFilterRuleItem{
			{
				MatchContext: &log_filter2.LogAttribute{
					LogFields: map[string]interface{}{
						"tag1": "tag1",
						"tag2": "tag2",
						"tag3": "tag3",
					},
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
			"tag3": "tagXXXXX", //本行不匹配
		},
	}
	result := logFilter.Filter(logContext)
	asserts.True(result, "")
}
