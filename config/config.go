package config

import (
	"github.com/spf13/viper"
	"go.uber.org/fx"
)

type Config struct {
	Port      string       `mapstructure:"port"`
	RunPodApi RunPodConfig `mapstructure:"runpod-api"`
}

type RunPodConfig struct {
	BaseUrl string `mapstructure:"baseurl"`
	ApiKey  string `mapstructure:"apikey"`
}

var cfg = &Config{}
var CfgPath = "config/config.yml"
var Module = fx.Provide(initCfg)

func initCfg() *Config {
	// 使用viper读取配置
	viper.SetConfigFile(CfgPath)
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	err := viper.Unmarshal(cfg)
	if err != nil {
		panic(err)
	}
	return cfg
}

func GetConfig() *Config {
	return cfg
}
