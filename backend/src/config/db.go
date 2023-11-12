package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type DBConfig struct {
	DSN string
}

func LoadDBConfig() DBConfig {
	return DBConfig{
		DSN: fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
			viper.GetString("DB_USER"),
			viper.GetString("DB_PASSWORD"),
			viper.GetString("DB_HOST"),
			viper.GetString("DB_PORT"),
			viper.GetString("DB_NAME"),
		),
	}
}