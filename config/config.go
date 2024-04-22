package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	PORT        string
	ADDRESS     string
	DB_HOST     string
	DB_USER     string
	DB_PASSWORD string
	DB_NAME     string
	DB_PORT     string
}

var C *Config

func InitConfig() {
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	if err := viper.Unmarshal(&C); err != nil {
		panic(fmt.Errorf("unable to decode into struct, %w", err))
	}
}
