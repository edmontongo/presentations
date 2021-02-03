package main

import (
	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()
	logger.Sugar().Errorw("This is an error message", "error", "Developer ignored the error")
	logger.Debug("Unsugarred loggers",
		zap.Bool("verbose", true),
		zap.String("do I use them", "not really"),
	)
}
