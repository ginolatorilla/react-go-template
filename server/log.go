package main

import "go.uber.org/zap"

func logger() *zap.SugaredLogger {
	return zap.S()
}
