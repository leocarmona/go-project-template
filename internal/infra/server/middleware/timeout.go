package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/leocarmona/go-project-template/internal/infra/logger"
	"github.com/leocarmona/go-project-template/internal/infra/logger/attributes"
	"github.com/leocarmona/go-project-template/internal/infra/variables"
	"time"
)

var (
	DefaultTimeoutConfig = middleware.TimeoutConfig{
		Skipper:      middleware.DefaultSkipper,
		ErrorMessage: "{\"error\":\"Request Timeout\"}",
		OnTimeoutRouteErrorHandler: func(err error, c echo.Context) {
			logger.Warn(c.Request().Context(), "Request Timeout", attributes.Attributes{
				"uri": c.Request().RequestURI,
			}.WithError(err))
		},
		Timeout: time.Second * time.Duration(variables.ServerTimeout()),
	}
)

// ConfigTimeout middleware adds a `timeout`
func ConfigTimeout() echo.MiddlewareFunc {
	return middleware.TimeoutWithConfig(DefaultTimeoutConfig)
}
