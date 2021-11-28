package app

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/facebookgo/grace/gracehttp"
	"github.com/labstack/echo/v4"
	"github.com/leocarmona/go-project-template/internal/app/adapters"
	"github.com/leocarmona/go-project-template/internal/app/domain"
	"github.com/leocarmona/go-project-template/internal/infra/aws"
	"github.com/leocarmona/go-project-template/internal/infra/database"
	"github.com/leocarmona/go-project-template/internal/infra/logger"
	"github.com/leocarmona/go-project-template/internal/infra/logger/attributes"
	"github.com/leocarmona/go-project-template/internal/infra/server"
	"github.com/leocarmona/go-project-template/internal/infra/variables"
	"sync"
	"time"
)

type App struct {
	running bool
	locker  sync.Mutex

	server    *echo.Echo
	handlers  *adapters.Handlers
	services  *domain.Services
	databases *database.Databases
}

var app = new(App)

func Instance() *App {
	return app
}

func (app *App) Start(async bool) {
	app.locker.Lock()

	if app.running {
		app.locker.Unlock()
		return
	}

	start := time.Now()
	logger.Info(context.Background(), fmt.Sprintf("Starting application %s:%s", variables.ServiceName(), variables.ServiceVersion()), nil)

	app.build()

	if async {
		go app.startServer(start)
	} else {
		app.startServer(start)
	}
}

func (app *App) Stop() {
	app.locker.Lock()

	if !app.running {
		app.locker.Unlock()
		return
	}

	defer app.setRunning(false)
	defer app.locker.Unlock()

	logger.Warn(context.Background(), "Stopping application", nil)

	if err := app.server.Close(); err != nil {
		logger.Error(context.Background(), "Error while trying to close echo server", attributes.New().WithError(err))
	}

	app.databases.Close()
	app.dispose()

	logger.Warn(context.Background(), "Application stopped", nil)
}

func (app *App) IsRunning() bool {
	return app.running
}

func (app *App) startServer(start time.Time) {
	defer app.setRunning(false)
	go func() {
		app.printElapsed(start)
		app.locker.Unlock()
	}()

	if variables.IsLambda() {
		lambdaAdapter := &aws.LambdaAdapter{Echo: app.server}
		lambda.Start(lambdaAdapter.Handler)
		logger.Warn(context.Background(), "Application stopped [Lambda]", nil)
	} else {
		err := gracehttp.Serve(app.server.Server)
		logger.Warn(context.Background(), "Application stopped gracefully", attributes.New().WithError(err))
	}
}

func (app *App) build() {
	app.databases = database.NewDatabases()
	app.services = domain.NewServices(app.databases)
	app.handlers = adapters.NewHandlers(app.services)
	app.server = server.New()
	app.handlers.Configure(app.server)
}

func (app *App) dispose() {
	app.server = nil
	app.handlers = nil
	app.services = nil
	app.databases = nil
}

func (app *App) printElapsed(start time.Time) {
	elapsed := time.Since(start)
	logger.Info(context.Background(), fmt.Sprintf("Application %s:%s started in %v", variables.ServiceName(), variables.ServiceVersion(), elapsed.String()), nil)
}

func (app *App) setRunning(run bool) {
	app.running = run
}
