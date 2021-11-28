package test

import (
	"context"
	"fmt"
	"github.com/leocarmona/go-project-template/internal/app"
	"github.com/leocarmona/go-project-template/internal/infra/logger"
	"github.com/leocarmona/go-project-template/internal/infra/variables"
	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/baloo.v3"
	"os"
	"os/exec"
	"testing"
	"time"
)

func InitLogger() {
	logger.Init(&logger.Option{
		ServiceName:    variables.ServiceName(),
		ServiceVersion: variables.ServiceVersion(),
		Environment:    variables.Environment(),
		LogLevel:       variables.LogLevel(),
	})
}

func ComposeUp(t *testing.T) {
	Shell(t, "cd ../ && make compose-up")
	time.Sleep(1 * time.Second)
}

func ComposeDown(t *testing.T) {
	Shell(t, "cd ../ && make compose-down")
	time.Sleep(1 * time.Second)
}

func StartApplication() {
	_ = os.Setenv("DB_READ_PORT", "6432")
	_ = os.Setenv("DB_WRITE_PORT", "6432")

	if app.Instance().IsRunning() {
		return
	}

	app.Instance().Start(true)
	time.Sleep(10 * time.Millisecond)
}

func Request() *baloo.Client {
	return baloo.New("http://localhost:5000")
}

func Shell(t *testing.T, command string) {
	logger.Info(context.Background(), fmt.Sprintf("Executing command [%s]", command), nil)
	out, err := exec.Command("/bin/sh", "-c", command).Output()

	if err != nil {
		if ee, ok := err.(*exec.ExitError); ok {
			assert.FailNow(t, string(ee.Stderr))
		}

		assert.FailNow(t, err.Error())
	}

	logger.Info(context.Background(), fmt.Sprintf("Response [%s]", out), nil)
}
