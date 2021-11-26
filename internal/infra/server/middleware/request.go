package middleware

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/leocarmona/go-project-template/internal/infra/logger"
	"github.com/leocarmona/go-project-template/internal/infra/logger/attributes"
	"github.com/leocarmona/go-project-template/internal/infra/request"
	"net/http"
	"strings"
	"time"
)

func ConfigRequest() echo.MiddlewareFunc {
	return newRequest
}

func newRequest(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		start := time.Now()

		_ = request.BuildContext(c)

		response := next(c)
		elapsed := time.Since(start)

		logPostRequest(c.Request(), elapsed)

		return response
	}
}

func logPostRequest(req *http.Request, duration time.Duration) {
	details := generateDefaultAttributes(req)
	details["request.duration"] = duration

	logger.Info(req.Context(), fmt.Sprintf("Request handled [%s %s] [%s]", details["request.method"], details["request.uri"], duration.String()), details)
}

func generateDefaultAttributes(req *http.Request) attributes.Attributes {
	details := attributes.New()
	method, path, params, uri := extractUri(req)

	details["request.method"] = method
	details["request.path"] = path
	details["request.params"] = params
	details["request.uri"] = uri

	return details
}

func extractUri(r *http.Request) (method string, path string, params string, uri string) {
	method = r.Method
	uri = r.RequestURI
	parts := strings.Split(uri, "?")

	if partsLen := len(parts); partsLen > 0 {
		path = parts[0]

		for i := 1; i < partsLen; i++ {
			params += parts[i] + "?"
		}

		if paramsLen := len(params); paramsLen > 0 {
			params = params[:paramsLen-1]
		}
	}

	return
}
