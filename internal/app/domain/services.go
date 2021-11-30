package domain

import (
	dbAdapters "github.com/leocarmona/go-project-template/internal/app/adapters/database"
	"github.com/leocarmona/go-project-template/internal/app/adapters/database/postgres"
	"github.com/leocarmona/go-project-template/internal/app/domain/book"
	"github.com/leocarmona/go-project-template/internal/app/domain/health"
	"github.com/leocarmona/go-project-template/internal/infra/database"
)

type Services struct {
	Book   *book.Service
	Health *health.Service
}

func NewServices(dbs *database.Databases) *Services {
	return &Services{
		Book:   book.NewService(postgres.NewBookRepository(dbs)),
		Health: health.NewService(dbAdapters.NewHealthRepository(dbs)),
	}
}
