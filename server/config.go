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
	viper.BindEnv("enable_cors")
	viper.SetDefault("enable_cors", false)

	cfg := serverConfig{
		listenAddress: viper.GetString("listen_address"),
		enableCORS:    viper.GetBool("enable_cors"),
	}

	logger().Debugf("Config loaded from environment: %#v", cfg)
	return cfg
}
