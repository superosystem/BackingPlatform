package config

import "github.com/spf13/viper"

type PayGateConfig struct {
	PGClientKey string
	PGServerKey string
}

func LoadPayGateConfig() PayGateConfig {
	return PayGateConfig{
		PGClientKey: viper.GetString("PG_CLIENT_KEY"),
		PGServerKey: viper.GetString("PG_SERVER_KEY"),
	}
}