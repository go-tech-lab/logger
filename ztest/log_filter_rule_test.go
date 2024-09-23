package test

import (
	"encoding/json"
	log_filter2 "github.com/go-tech-lab/framework/logger/src/log_filter"
	_ "github.com/go-tech-lab/framework/logger/src/logger/impl"
	"github.com/go-tech-lab/framework_common/src/util"
	"testing"
)

func TestLogFilterRuleMarshal(t *testing.T) {
	logger := mockZapLogger()
	ruleItem := &log_filter2.LogFilterRuleItem{
		MatchContext: &log_filter2.LogAttribute{
			ModuleTag:     "MYSQL",
			RequestMethod: "PayManager.QueryUserInstallmentInfo",
			Level:         "INFO",
			FuncName:      "Funcx",
			FileName:      "Funcx",
			Line:          0,
			LogFields:     nil,
		},
		Rate: 0.50,
	}
	logFilterRule := log_filter2.LogFilterRule{FilterRuleItems: []*log_filter2.LogFilterRuleItem{ruleItem, ruleItem, ruleItem}}
	jsonStr, err := json.Marshal(logFilterRule)
	util.CheckErr(err)
	logger.Info(string(jsonStr))
}

func TestLogFilterRuleUnmarshal(t *testing.T) {
	logFilterRule := &log_filter2.LogFilterRuleItem{}
	bytes := []byte("{\"MatchContext\":{\"Level\":\"INFO\",\"FuncContext\":{\"FuncName\":\"Funcx\",\"FileName\":\"Funcx\",\"Line\":0},\"RequestContext\":{\"Method\":\"\",\"ClientIp\":\"\",\"FromServer\":\"\",\"TraceId\":\"\",\"User\":\"\"},\"LogFields\":null},\"Rate\":0.5}")
	err := json.Unmarshal(bytes, logFilterRule)
	util.CheckErr(err)
}
