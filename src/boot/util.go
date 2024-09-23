package boot

import (
	"encoding/json"
	"fmt"
	"github.com/go-tech-lab/framework/logger/src/config"
	"github.com/go-tech-lab/framework/logger/src/log_filter"
	"github.com/go-tech-lab/framework/logger/src/log_filter/impl"
	"github.com/go-tech-lab/framework/logger/src/logger"
	"github.com/go-tech-lab/framework_common/src/util"
	"github.com/robfig/cron"
	"gopkg.in/ini.v1"
)

//从日志配置文件中加载zap配置
func loadZapLoggerConfig(configFilePath string) *config.ZapLogConfig {
	iniFile, err := ini.Load(configFilePath) //加载配置文件
	util.CheckErr(err)
	var zapConfig config.ZapLogConfig
	err = iniFile.Section("zap_logger_config").MapTo(&zapConfig)
	util.CheckErr(err)
	return &zapConfig
}


//LoadLogFilter 从配置文件加载过滤器
func LoadLogFilter(fileConfigPath string) log_filter.ILogFilter {
	logFilterRule := loadLogFilterRuleConfig(fileConfigPath)
	logFilter := impl.NewDefaultLogFilter(logFilterRule)
	return logFilter
}

//从日志配置文件中加载日志过滤器
func loadLogFilterRuleConfig(configFilePath string) *log_filter.LogFilterRule {
	iniFile, err := ini.Load(configFilePath) //加载配置文件
	util.CheckErr(err)
	filterRulesStr := iniFile.Section("logger_filter_rule").Key("filter_rule_items").String()
	if len(filterRulesStr) > 0 {
		fmt.Printf("loadLogFilterRuleConfig: %s\n", filterRulesStr)
		logFilterRule := &log_filter.LogFilterRule{}
		err = json.Unmarshal([]byte(filterRulesStr), logFilterRule)
		util.CheckErr(err)
		return logFilterRule
	}
	return nil
}

//加载日志过滤规则刷新周期
func loadLogFilterRefreshCron(configFilePath string) string {
	iniFile, err := ini.Load(configFilePath) //加载配置文件
	util.CheckErr(err)
	filterRefreshCron := iniFile.Section("logger_filter_rule").Key("filter_refresh_cron").String()
	fmt.Printf("loadLogFilterRefreshCron: %s\n", filterRefreshCron)
	return filterRefreshCron
}

//给logger注册日志过滤刷新task
func registerFilterRefreshTask(logger logger.ILogger, fileConfigPath string) {
	logFilterRefreshCron := loadLogFilterRefreshCron(fileConfigPath)
	if len(logFilterRefreshCron) > 0 {
		cronMaster := cron.New()
		cronMaster.AddFunc(logFilterRefreshCron, func() {
			defer func() {
				if ret := recover(); ret != nil {
					err, ok := ret.(error)
					if ok {
						logger.Errorf("refresh log filter rule error: %s", err.Error())
					} else {
						logger.Errorf("refresh log filter rule error: %v", ret)
					}
				} else {
					logger.Infof("refresh log filter rule success")
				}
			}()
			refreshLogFilter := LoadLogFilter(fileConfigPath)
			logger.WithFilter(refreshLogFilter)
		})
		cronMaster.Start()
	}
}


//定时输出日志，防止无流量服务基础日志文件被归档，输出日志异常
func registerPreventLogBaseFileArchiveTask(logger logger.ILogger) {
	logPreventBaseFileArchiveCron := "0 */20 * * * *"
	logger.Info("output timing log")
	cronMaster := cron.New()
	cronMaster.AddFunc(logPreventBaseFileArchiveCron, func() {
		logger.Info("output timing log")
	})
	cronMaster.Start()

}
