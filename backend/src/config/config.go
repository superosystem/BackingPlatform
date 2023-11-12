package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	App     AppConfig
	DB      DBConfig
	PayGate PayGateConfig
}

func NewConfig() *Config {
	viper.SetConfigFile(`.env`)
	viper.AutomaticEnv()
	viper.ReadInConfig()

	if viper.GetBool(`DEBUG`) {
		log.Panicln("Service RUN on DEBUG mode")
	}

	return &Config{
		App:     LoadAppConfig(),
		DB:      LoadDBConfig(),
		PayGate: LoadPayGateConfig(),
	}
}