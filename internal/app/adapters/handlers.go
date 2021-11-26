package adapters

import (
	"github.com/labstack/echo/v4"
	"github.com/leocarmona/go-project-template/internal/app/adapters/handler"
	"github.com/leocarmona/go-project-template/internal/app/domain"
)

type Handlers struct {
	health *handler.HealthHandler
}

func NewHandlers(services *domain.Services) *Handlers {
	return &Handlers{
		health: handler.NewHealthHandler(services),
	}
}

func (h *Handlers) Configure(server *echo.Echo) {
	h.health.Configure(server)
}
