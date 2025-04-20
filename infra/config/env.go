package config

import (
	"log/slog"

	"github.com/spf13/viper"
)

func LoadEnv() {

	viper.SetConfigName("env")
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()

	if err != nil {
		slog.Info("could not read config", "err", err.Error())
	}
}
