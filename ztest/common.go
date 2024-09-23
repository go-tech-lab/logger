package test

import (
	"encoding/json"
	"github.com/go-tech-lab/framework/logger/src/config"
	"github.com/go-tech-lab/framework/logger/src/log_filter"
	"github.com/go-tech-lab/framework/logger/src/log_filter/impl"
	"github.com/go-tech-lab/framework/logger/src/logger"
	impl2 "github.com/go-tech-lab/framework/logger/src/logger/impl"
	"github.com/go-tech-lab/framework_common/src/util"
)

func mockZapLogger() logger.ILogger {
	zapConfig := &config.ZapLogConfig{
		OutputDir:      "/Users/changlinguo/shopee/gopath/src/git.garena.com/shopee/loan-service/credit_backend/credit_framework/framework_logger/log",
		FileBase:       "all.log",
		TimeFormat:     "2006-01-02 15:04:05.000000",
		MaxAgeHour:     72,
		RotationHour:   1,
		FileTailFormat: ".%Y-%m-%d-%H",
		EnableDebug:    true,
	}
	zapLogger := impl2.NewZapLogger(zapConfig)

	logFilterRuleStr := `{"FilterRuleItems":[
							{
								"MatchContext":{
									"Level":"info",
									"FuncContext":{
										"FuncName":"Funcx",
										"FileName":"Funcx",
										"Line":0
									},
									"RequestContext":{
										"Method":"",
										"ClientIp":"",
										"FromServer":"",
										"TraceId":"",
										"User":""
									},
									"LogFields":null
								},
								"Rate":0.5
							}
						]
					}`
	logFilterRule := &log_filter.LogFilterRule{}
	err := json.Unmarshal([]byte(logFilterRuleStr), logFilterRule)
	util.CheckErr(err)
	logFilter := impl.NewDefaultLogFilter(logFilterRule)
	zapLogger.WithFilter(logFilter)
	zapLogger.WithLogWrapper(impl2.NewTroughLogTagsWrapper(func() logger.LogFields {
		return map[string]interface{}{
			"logDecorator_key1": "logDecorator_value1",
			"logDecorator_key2": "logDecorator_value2",
			"logDecorator_key3": "logDecorator_value3",
		}
	}))
	return zapLogger
}
