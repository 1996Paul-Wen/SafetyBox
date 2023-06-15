package config

import (
	"os"
	"time"

	"gopkg.in/yaml.v3"
)

type APPConfig struct {
	Database                    DatabaseConfig `yaml:"Database"`
	KeyToEncryptUserLoginPWSalt string         `yaml:"KeyToEncryptUserLoginPWSalt"`
	Debug                       bool           `yaml:"Debug"`
}

type DatabaseConfig struct {
	DSN               string        `yaml:"DSN"`               // such as "gorm:gorm@tcp(127.0.0.1:3306)/gorm?charset=utf8&parseTime=True&loc=Local"
	DefaultStringSize uint          `yaml:"DefaultStringSize"` // string 类型字段的默认长度
	MaxOpenConns      int           `yaml:"MaxOpenConns"`
	MaxIdleConns      int           `yaml:"MaxIdleConns"`
	ConnMaxIdleTime   time.Duration `yaml:"ConnMaxIdleTime"`
}

// new returns an empty config
func new() *APPConfig {
	return &APPConfig{}
}

// Init should only be called once for a single APPConfig instance
func (c *APPConfig) Init(appConfigFile string) {
	conf, err := os.ReadFile(appConfigFile)
	if err != nil {
		panic(err)
	}

	globalConfigContent := APPConfig{}
	err = yaml.Unmarshal(conf, &globalConfigContent)
	if err != nil {
		panic(err)
	}

	*c = globalConfigContent
}
