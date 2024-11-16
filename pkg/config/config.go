package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type Configs struct {
	Listen   Listen       `mapstructure:"listen"`
	ScyllaDB ScyllaConfig `mapstructure:"scylla_db"`
}

type Listen struct {
	Host string `mapstructure:"host"`
	Port string `mapstructure:"port"`
}

type ScyllaConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	KeySpace string `mapstructure:"keyspace"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	SSL      bool   `mapstructure:"ssl"`
}

func LoadConfiguration() (*Configs, error) {

	pathConfig := "./config.yml"

	var config Configs
	viper.SetConfigFile(pathConfig)
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("Error reading config.yml file, %s", err)
		return nil, err
	}

	if err = viper.Unmarshal(&config); err != nil {
		fmt.Printf("Error unmarshalling config.yml, %s", err)
		return nil, err
	}

	return &config, nil
}
