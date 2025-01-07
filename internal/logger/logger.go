package logger

import (
	"go.uber.org/zap"
)

var logger *zap.Logger
var sugar *zap.SugaredLogger

func Init() {
	logger, _ = zap.NewProduction()
	defer logger.Sync()
	sugar = logger.Sugar()
}

func Sync() {
	logger.Sync()
}

func Error() {

}
