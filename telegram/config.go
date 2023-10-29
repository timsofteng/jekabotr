package main

import (
	"os"

	"github.com/spf13/viper"
)

// var C Config

func ReadConfig() (*Config, error) {
	config := &MultiEnvConfig{}

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".") // optionally look for config in the working directory

	err := viper.ReadInConfig()

	err = viper.Unmarshal(config)

	env := os.Getenv("ENV")

	if env == "prod" {
		return &config.Prod, err
	}

	return &config.Dev, err
}
