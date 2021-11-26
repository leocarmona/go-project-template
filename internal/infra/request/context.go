package request

import (
	"context"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

const (
	contextKey          ContextKey = "request-context"
	headerCid           string     = "x-cid"
	headerAuthorization string     = "authorization"
)

type (
	ContextKey string

	Context struct {
		CID           string
		Authorization string
	}
)

func BuildContext(c echo.Context) *Context {
	return buildContext(c)
}

func GetContext(c echo.Context) *Context {
	return c.Request().Context().Value(contextKey).(*Context)
}

func buildContext(c echo.Context) *Context {
	rctx := extractContext(c)
	ctx := context.WithValue(c.Request().Context(), contextKey, rctx)
	c.SetRequest(c.Request().WithContext(ctx))

	return rctx
}

func extractContext(c echo.Context) *Context {
	ctx := &Context{
		CID:           extractCid(c),
		Authorization: extractAuthorization(c),
	}

	return ctx
}

func extractCid(c echo.Context) string {
	cid := c.Request().Header.Get(headerCid)

	if len(cid) == 0 {
		cid = uuid.New().String()
		c.Request().Header.Set(headerCid, cid)
	}

	return cid
}

func extractAuthorization(c echo.Context) string {
	return c.Request().Header.Get(headerAuthorization)
}
