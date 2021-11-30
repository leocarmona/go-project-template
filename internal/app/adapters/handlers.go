package adapters

import (
	"github.com/labstack/echo/v4"
	"github.com/leocarmona/go-project-template/internal/app/adapters/handler"
	"github.com/leocarmona/go-project-template/internal/app/domain"
)

type Handlers struct {
	book   *handler.BookHandler
	health *handler.HealthHandler
}

func NewHandlers(services *domain.Services) *Handlers {
	return &Handlers{
		book:   handler.NewBookHandler(services),
		health: handler.NewHealthHandler(services),
	}
}

func (h *Handlers) Configure(server *echo.Echo) {
	h.book.Configure(server)
	h.health.Configure(server)
}
