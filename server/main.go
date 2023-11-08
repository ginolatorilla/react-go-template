package main

import "go.uber.org/zap"

var version string
var app string

func main() {
	logger, _ := zap.NewDevelopment()
	zap.ReplaceGlobals(logger)
	defer logger.Sync()

	server := NewServer(ReadConfigFromEnv())
	server.Serve()
}
