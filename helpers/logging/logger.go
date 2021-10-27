package logging

import "go.uber.org/zap"

const (
	KeyErr = "error"
	KeyID  = "id"
)

func NewLogger() *zap.SugaredLogger {
	logger, _ := zap.NewProduction()
	return logger.Sugar()
}
