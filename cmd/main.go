package main

import (
	"go/rest/internal/app"
	"go/rest/pkg/logger"
)

func main() {

	appLogger := logger.Initialize()

	app.New(appLogger)

}
