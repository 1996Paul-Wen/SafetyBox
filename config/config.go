package config

import (
	"os"
	"time"

	"gopkg.in/yaml.v3"
)

type APPConfig struct {
	Database                 DatabaseConfig `yaml:"Database"`
	AESKeyForUserLoginPWSalt string         `yaml:"AESKeyForUserLoginPWSalt"`
	Debug                    bool           `yaml:"Debug"`
	WebSettings              WebSettings    `yaml:"WebSettings"`
}

type DatabaseConfig struct {
	DSN               string        `yaml:"DSN"`               // such as "gorm:gorm@tcp(127.0.0.1:3306)/gorm?charset=utf8&parseTime=True&loc=Local"
	DefaultStringSize uint          `yaml:"DefaultStringSize"` // string 类型字段的默认长度
	MaxOpenConns      int           `yaml:"MaxOpenConns"`
	MaxIdleConns      int           `yaml:"MaxIdleConns"`
	ConnMaxIdleTime   time.Duration `yaml:"ConnMaxIdleTime"`
}

type WebSettings struct {
	Port          int           `yaml:"Port"`
	LimitSettings LimitSettings `yaml:"LimitSettings"`
}

type LimitSettings struct {
	Limit float64 `yaml:"Limit"`
	Burst float64 `yaml:"Burst"`
}

// new returns an empty config
func new() *APPConfig {
	return &APPConfig{}
}

// Init should only be called once
func Init(appConfigFile string) {
	conf, err := os.ReadFile(appConfigFile)
	if err != nil {
		panic(err)
	}

	globalConfigContent := APPConfig{}
	err = yaml.Unmarshal(conf, &globalConfigContent)
	if err != nil {
		panic(err)
	}

	*globalConfig = globalConfigContent
}
