package main

import (
	"github.com/leocarmona/go-project-template/internal/app"
	"github.com/leocarmona/go-project-template/internal/infra/logger"
)

func main() {
	logger.Init()
	defer func() {
		_ = logger.Sync()
	}()

	application := app.Instance()
	application.Start()
}
