package main

import (
	"github.com/leocarmona/go-project-template/internal/app"
	"github.com/leocarmona/go-project-template/internal/infra/logger"
	"github.com/leocarmona/go-project-template/internal/infra/variables"
)

func main() {
	logger.Init(&logger.Option{
		ServiceName:    variables.ServiceName(),
		ServiceVersion: variables.ServiceVersion(),
		Environment:    variables.Environment(),
		LogLevel:       variables.LogLevel(),
	})

	defer logger.Sync()

	application := app.Instance()
	application.Start(false)
}
