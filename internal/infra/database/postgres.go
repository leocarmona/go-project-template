package database

import (
	"fmt"
)

const (
	postgresDriver           = "postgres"
	postgresConnectionString = "postgres://%s:%s@%s:%s/%s?sslmode=disable&fallback_application_name=%s"
)

func NewPostgres(cfg *SqlConfig) *Database {
	if len(cfg.Driver) == 0 {
		cfg.Driver = postgresDriver
	}

	return NewDatabase(cfg, postgresConnectionStringBuilder)
}

func postgresConnectionStringBuilder(cfg *SqlConfig) string {
	return fmt.Sprintf(postgresConnectionString, cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Database, cfg.ConnectionName)
}
