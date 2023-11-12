package config

import "github.com/spf13/viper"

type AppConfig struct {
	Name    string
	Version string
	Mode    string
	Port    string
}

func LoadAppConfig() AppConfig {
	return AppConfig{
		Name:    viper.GetString("APP_NAME"),
		Version: viper.GetString("APP_VERSION"),
		Mode:    viper.GetString("APP_MODE"),
		Port:    viper.GetString("APP_PORT"),
	}
}