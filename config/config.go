package config

import (
	"fmt"
	"log"
	"os"
	"jekabot/models"

	"github.com/spf13/viper"
)


// var C Config

func ReadConfig() *models.Config {
	config := &models.Config{}

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config") // optionally look for config in the working directory

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println(err)
		log.Fatalln(err)
	}

	if err := viper.Unmarshal(config); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return config
}
