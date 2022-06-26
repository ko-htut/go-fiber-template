package log

import (
	. "github.com/hotrungnhan/go-fiber-template/pkg/configs"
	"go.uber.org/zap"
)

var log *zap.SugaredLogger

func SetupLogger() {
	var logger *zap.Logger
	if Get().STAGE == DEVELOPMENT {
		logger, _ = zap.NewDevelopment()
	} else {
		logger, _ = zap.NewProduction()
	}
	defer logger.Sync() // flushes buffer, if any
	log = logger.Sugar()
}

func Get() *zap.SugaredLogger {
	return log
}
