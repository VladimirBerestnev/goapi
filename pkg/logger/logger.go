package logger

import (
	"sync"

	"go.uber.org/zap"
)

var (
	once   sync.Once
	logger *zap.Logger
)

func GetLogger() *zap.Logger {
	once.Do(func() {
		var err error
		logger, err = zap.NewDevelopment()
		if err != nil {
			panic(err)
		}
	})
	return logger
}
