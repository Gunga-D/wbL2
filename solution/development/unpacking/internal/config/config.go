package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	File FileConfig
}

type FileConfig struct {
	Path string `json:"path"`
	Name string `json:"name"`
}

func LoadConfig() (*viper.Viper, error) {
	v := viper.New()

	v.AddConfigPath("config")
	v.SetConfigName("config")
	v.SetConfigType("yml")
	err := v.ReadInConfig()
	if err != nil {
		return nil, err
	}
	return v, nil
}

func ParseConfig(v *viper.Viper) (*Config, error) {
	var c Config
	err := v.Unmarshal(&c)
	if err != nil {
		return nil, err
	}
	return &c, nil
}
