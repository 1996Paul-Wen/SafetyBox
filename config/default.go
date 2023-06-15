package config

var globalConfig *APPConfig = new()

func GlobalConfig() *APPConfig {
	return globalConfig
}
