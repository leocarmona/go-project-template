package database

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/leocarmona/go-project-template/internal/infra/variables"
	"sync"
)

type Databases struct {
	Read  *Database
	Write *Database
	Redis *Redis
}

func NewDatabases() *Databases {
	dbs := &Databases{}
	var waitGroup sync.WaitGroup
	defer waitGroup.Wait()

	dbs.buildReadDatabase(&waitGroup)
	dbs.buildWriteDatabase(&waitGroup)
	dbs.buildRedisDatabase(&waitGroup)

	return dbs
}

func (d *Databases) Close() {
	d.Read.Close()
	d.Write.Close()
	d.Redis.Close()
}

func (d *Databases) buildReadDatabase(waitGroup *sync.WaitGroup) {
	lazyConnection := variables.DBReadLazyConnection()
	cfg := &SqlConfig{
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
		LazyConnection:        lazyConnection,
	}

	if lazyConnection {
		d.Read = NewPostgres(cfg)
	} else {
		waitGroup.Add(1)
		go func() {
			defer waitGroup.Done()
			d.Read = NewPostgres(cfg)
		}()
	}
}

func (d *Databases) buildWriteDatabase(waitGroup *sync.WaitGroup) {
	lazyConnection := variables.DBWriteLazyConnection()
	cfg := &SqlConfig{
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
		LazyConnection:        lazyConnection,
	}

	if lazyConnection {
		d.Write = NewPostgres(cfg)
	} else {
		waitGroup.Add(1)
		go func() {
			defer waitGroup.Done()
			d.Write = NewPostgres(cfg)
		}()
	}
}

func (d *Databases) buildRedisDatabase(waitGroup *sync.WaitGroup) {
	lazyConnection := variables.RedisLazyConnection()
	opt := &redis.Options{
		Addr:     fmt.Sprintf("%s:%d", variables.RedisHost(), variables.RedisPort()),
		Password: variables.RedisPassword(),
		DB:       variables.RedisDB(),
	}

	if lazyConnection {
		d.Redis = NewRedis(opt, lazyConnection)
	} else {
		waitGroup.Add(1)
		go func() {
			defer waitGroup.Done()
			d.Redis = NewRedis(opt, lazyConnection)
		}()
	}
}
