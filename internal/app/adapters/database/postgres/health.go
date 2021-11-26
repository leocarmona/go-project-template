package postgres

import (
	"github.com/leocarmona/go-project-template/internal/app/domain/health"
	"github.com/leocarmona/go-project-template/internal/infra/database"
	"golang.org/x/net/context"
)

type HealthRepository struct {
	dbs *database.Databases
}

func NewHealthRepository(dbs *database.Databases) *HealthRepository {
	return &HealthRepository{
		dbs: dbs,
	}
}

func (r *HealthRepository) HealthReadDB(ctx context.Context) *health.Health {
	return r.checkConnection(ctx, r.dbs.Read)
}

func (r *HealthRepository) HealthWriteDB(ctx context.Context) *health.Health {
	return r.checkConnection(ctx, r.dbs.Write)
}

func (r *HealthRepository) checkConnection(ctx context.Context, db *database.Database) *health.Health {
	row := db.Connection().QueryRowContext(ctx, healthCheckSql)

	if row.Err() != nil {
		return health.New(false, row.Err())
	}

	var version string
	err := row.Scan(&version)

	if err != nil {
		return health.New(false, row.Err())
	}

	return health.New(true, nil)
}

const (
	healthCheckSql = "SELECT version()"
)
