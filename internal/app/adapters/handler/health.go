package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/leocarmona/go-project-template/internal/app/domain"
	"github.com/leocarmona/go-project-template/internal/app/interface/presenter"
	"github.com/leocarmona/go-project-template/internal/infra/logger"
	"github.com/leocarmona/go-project-template/internal/infra/logger/attributes"
	"net/http"
)

type HealthHandler struct {
	services *domain.Services
}

func NewHealthHandler(services *domain.Services) *HealthHandler {
	return &HealthHandler{
		services: services,
	}
}

func (h *HealthHandler) Configure(server *echo.Echo) {
	server.GET("/health", h.health)
}

func (h *HealthHandler) health(c echo.Context) error {
	ctx := c.Request().Context()
	read := h.services.Health.HealthReadDB(ctx)
	write := h.services.Health.HealthWriteDB(ctx)

	if read.Error != nil {
		logger.Error(ctx, "Health check error on read database", attributes.New().WithError(read.Error))
	}

	if write.Error != nil {
		logger.Error(ctx, "Health check error on write database", attributes.New().WithError(write.Error))
	}

	response := presenter.HealthResponse(read, write)
	var code int

	if response.Healthy {
		code = http.StatusOK
	} else {
		code = http.StatusServiceUnavailable
	}

	return c.JSON(code, response)
}
