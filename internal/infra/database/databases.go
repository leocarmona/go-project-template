package database

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/leocarmona/go-project-template/internal/infra/variables"
)

type Databases struct {
	Read  *Database
	Write *Database
	Redis *Redis
}

func NewDatabases() *Databases {
	return &Databases{
		Read:  newReadDatabase(),
		Write: newWriteDatabase(),
		Redis: newRedisDatabase(),
	}
}

func (d *Databases) Close() {
	d.Read.Close()
	d.Write.Close()
	d.Redis.Close()
}

func newReadDatabase() *Database {
	return NewPostgres(&SqlConfig{
		ConnectionName:        variables.ServiceName() + "-read",
		Host:                  variables.DBReadHost(),
		Port:                  variables.DBReadPort(),
		Database:              variables.DBReadName(),
		Username:              variables.DBReadUsername(),
		Password:              variables.DBReadPassword(),
		MinConnections:        variables.DBReadMinConnections(),
		MaxConnections:        variables.DBReadMaxConnections(),
		ConnectionMaxLifetime: variables.DBReadConnectionMaxLifeTime(),
		ConnectionMaxIdleTime: variables.DBReadConnectionMaxIdleTime(),
		LazyConnection:        variables.IsLambda(),
	})
}

func newWriteDatabase() *Database {
	return NewPostgres(&SqlConfig{
		ConnectionName:        variables.ServiceName() + "-write",
		Host:                  variables.DBWriteHost(),
		Port:                  variables.DBWritePort(),
		Database:              variables.DBWriteName(),
		Username:              variables.DBWriteUsername(),
		Password:              variables.DBWritePassword(),
		MinConnections:        variables.DBWriteMinConnections(),
		MaxConnections:        variables.DBWriteMaxConnections(),
		ConnectionMaxLifetime: variables.DBWriteConnectionMaxLifeTime(),
		ConnectionMaxIdleTime: variables.DBWriteConnectionMaxIdleTime(),
		LazyConnection:        variables.IsLambda(),
	})
}

func newRedisDatabase() *Redis {
	return NewRedis(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", variables.RedisHost(), variables.RedisPort()),
		Password: variables.RedisPassword(),
		DB:       variables.RedisDB(),
	}, variables.IsLambda())
}
