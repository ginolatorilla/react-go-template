package main

import (
	"strings"

	"github.com/spf13/viper"
)

func ReadConfigFromEnv() serverConfig {
	viper.SetEnvPrefix(app)
	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
	viper.BindEnv("listen_address")
	viper.SetDefault("listen_address", "localhost:8080")

	cfg := serverConfig{
		listenAddress: viper.GetString("listen_address"),
	}

	return cfg
}
