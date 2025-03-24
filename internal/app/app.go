package app

import (
	"go/rest/internal/app/delivery"
	"go/rest/internal/app/routing"
	"go/rest/internal/repository"
	"go/rest/internal/usecase"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap/zapcore"
)

type ILogger interface {
	Debug(msg string, fields ...zapcore.Field)
	Error(msg string, fields ...zapcore.Field)
	Fatal(msg string, fields ...zapcore.Field)
	Info(msg string, fields ...zapcore.Field)
}

type App struct {
	engine *gin.Engine
	logger ILogger
}

func New(logger ILogger) *App {

	engine := gin.Default()

	handler := delivery.New(usecase.New(repository.New()))
	routing.APIrout(engine, handler)

	return &App{
		engine: engine,
		logger: logger,
	}

}
