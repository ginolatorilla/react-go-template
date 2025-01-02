package cmd

import (
	"fmt"
	"os"
	"strings"

	u "github.com/ginolatorilla/go-template/pkg/utils"
	"github.com/ginolatorilla/react-go-template/server"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func Execute() {
	cmd := newCommand(server.AppName)
	u.Check(cmd.Execute())
}

func newCommand(appName string) *cobra.Command {
	var configFile string
	var verbosity int

	cobra.OnInitialize(
		func() { setUpLogger(verbosity) },
		func() { configure(configFile, appName) },
	)

	cmd := &cobra.Command{
		Use:   appName,
		Short: "Runs the application web server",
		Run: func(cmd *cobra.Command, args []string) {
			opts := server.Options{
				ListenAddress: viper.GetString("listen-address"),
				EnableCORS:    viper.GetBool("enable-cors"),
			}
			srv := server.NewServer(opts)
			u.Check(srv.Serve())
		},
	}

	cmd.Flags().StringVar(
		&configFile,
		"config",
		"",
		fmt.Sprintf("Read configuration from this file (default is $HOME/.%s.yaml)", appName),
	)
	cmd.Flags().CountVarP(
		&verbosity,
		"verbose",
		"v",
		"Verbosity level. Use -v for verbose, -vv for more verbose, etc.",
	)
	cmd.Flags().String(
		"listen-address",
		"127.0.0.1:8080",
		"Listen on this address",
	)
	viper.BindPFlag("listen-address", cmd.Flags().Lookup("listen-address"))
	cmd.Flags().Bool(
		"enable-cors",
		false,
		"Allow CORS requests",
	)
	viper.BindPFlag("enable-cors", cmd.Flags().Lookup("enable-cors"))

	return cmd
}

// setUpLogger sets up the logger based on the verbosity level.
//
// This function mimics the default logging level of Python's logger (starts at WARNING).
func setUpLogger(verbosity int) {
	lvl := zap.InfoLevel
	trace := false

	switch verbosity {
	case 0:
		lvl = zap.InfoLevel
		trace = false
	case 1:
		lvl = zap.DebugLevel
		trace = false
	default:
		lvl = zap.DebugLevel
		trace = true
	}

	config := zap.NewDevelopmentConfig()
	config.Level = zap.NewAtomicLevelAt(lvl)
	config.DisableStacktrace = !trace
	zap.ReplaceGlobals(zap.Must(config.Build()))
}

// configure reads application options from a file and environment variables.
func configure(configFile, appName string) {
	if configFile != "" {
		viper.SetConfigFile(configFile)
		return
	}

	zap.S().Debug("No config file specified, searching for default config file")
	home := u.Must(os.UserHomeDir())
	viper.AddConfigPath(home)
	viper.SetConfigType("yaml")
	viper.SetConfigName(fmt.Sprintf(".%s", appName))

	viper.SetEnvPrefix(appName)
	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		zap.S().Info("Using config file:", viper.ConfigFileUsed())
	}
}
