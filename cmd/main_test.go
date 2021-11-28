package main

import (
	"github.com/leocarmona/go-project-template/internal/test"
	"net/http"
	"testing"
)

func TestShouldRunApp(t *testing.T) {
	test.InitLogger()
	test.ComposeUp(t)
	test.StartApplication()

	_ = test.Request().
		Get("/health").
		Expect(t).
		Status(http.StatusOK).
		Type("json").
		JSON(map[string]interface{}{
			"healthy":  true,
			"read_db":  "UP",
			"write_db": "UP",
			"redis_db": "UP",
		}).Done()
}
