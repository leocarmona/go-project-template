package database

import (
	"github.com/leocarmona/go-project-template/internal/infra/variables"
)

type Databases struct {
	Read  *Database
	Write *Database
}

func NewDatabases() *Databases {
	read := NewPostgres(&SqlConfig{
		ConnectionName:        variables.AppName() + "-read",
		Host:                  variables.DBHost(),
		Port:                  variables.DBPort(),
		Database:              variables.DBName(),
		Username:              variables.DBUsername(),
		Password:              variables.DBPassword(),
		MinConnections:        variables.DBMinConnections(),
		MaxConnections:        variables.DBMaxConnections(),
		ConnectionMaxLifetime: variables.DBConnectionMaxIdleTime(),
		ConnectionMaxIdleTime: variables.DBConnectionMaxIdleTime(),
		LazyConnection:        variables.IsLambda(),
	})

	write := NewPostgres(&SqlConfig{
		ConnectionName:        variables.AppName() + "-write",
		Host:                  variables.DBHost(),
		Port:                  variables.DBPort(),
		Database:              variables.DBName(),
		Username:              variables.DBUsername(),
		Password:              variables.DBPassword(),
		MinConnections:        variables.DBMinConnections(),
		MaxConnections:        variables.DBMaxConnections(),
		ConnectionMaxLifetime: variables.DBConnectionMaxIdleTime(),
		ConnectionMaxIdleTime: variables.DBConnectionMaxIdleTime(),
		LazyConnection:        variables.IsLambda(),
	})

	return &Databases{
		Read:  read,
		Write: write,
	}
}

func (d *Databases) Close() {
	d.Read.Close()
	d.Write.Close()
}
