package config

import (
	"errors"
	"github.com/spf13/viper"
)

type Config struct {
	App   AppConfig
	MySQL MySQLConfig
}

type AppConfig struct {
	Port    string
	Mode    string
	Version string
}
type MySQLConfig struct {
	Port     string
	Host     string
	User     string
	Password string
	DBName   string
}

func LoadConfig(filename string) (*Config, error) {

	v := viper.New()

	v.SetConfigFile(filename)
	v.AutomaticEnv()

	if err := v.ReadInConfig(); err != nil {
		// Kiểm tra có phải là lỗi không tìm thấy file config hay không

		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return nil, errors.New("config file not found")
		}
		return nil, err
	}

	var c Config
	if err := v.Unmarshal(&c); err != nil {
		return nil, err
	}
	return &c, nil
}
