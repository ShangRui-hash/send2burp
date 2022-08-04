package config

//CurrentConf 当前配置
var CurrentConf AppConfig

//AppConfig App配置
type AppConfig struct {
	Debug             bool   `mapstructure:"debug"`
	ConfigurationName string `mapstructure:"configuration_name"`
}
