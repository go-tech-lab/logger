package config

//ZapLogConfig zap日志配置项
type ZapLogConfig struct {
	//日志问价输出目录
	OutputDir string `ini:"output_dir"`
	//日志中的时间格式
	TimeFormat string `ini:"time_format"`
	//日志文件最大保存时间-小时
	MaxAgeHour int `ini:"max_age_hour"`
	//日志文件滚动周期-小时
	RotationHour int `ini:"rotation_hour"`
	//日志文件名（前缀）
	FileBase string `ini:"file_base"`
	//日志文件名后缀格式
	FileTailFormat string `ini:"file_tail_format"`
	//日志输出格式(console,json) : 默认是console格式，
	LogFormat string `ini:"log_format"`
	//是否开启日志监控:默认不开启，开启对监控资源消耗较大，一般不建议开启
	EnableMonitor bool `ini:"enable_monitor"`

	EnableDebug bool `ini:"enable_debug"`

	//对接日志平台
	UseLogPlatform bool `ini:"use_log_platform"`
	//写操作 sync or async
	WriteMode string `ini:"write_mode"`
	//写操作 sync or async
	TimeRollingPattern string `ini:"time_rolling_pattern"`

	EnableLogLevel string `ini:"enable_log_level"`
}
