package domain

import (
	"github.com/leocarmona/go-project-template/internal/app/adapters/database/postgres"
	"github.com/leocarmona/go-project-template/internal/app/domain/health"
	database2 "github.com/leocarmona/go-project-template/internal/infra/database"
)

type Services struct {
	Health *health.Service
}

func NewServices(dbs *database2.Databases) *Services {
	return &Services{
		Health: health.NewService(postgres.NewHealthRepository(dbs)),
	}
}
