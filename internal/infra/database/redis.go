package database

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/leocarmona/go-project-template/internal/infra/logger"
	"github.com/leocarmona/go-project-template/internal/infra/logger/attributes"
	"sync"
	"time"
)

type Redis struct {
	rdb    *redis.Client
	opt    *redis.Options
	locker sync.Mutex
}

func NewRedis(opt *redis.Options, lazyConnection bool) *Redis {
	rdb := &Redis{
		opt: opt,
	}

	if !lazyConnection {
		_ = rdb.initializeAndGetRedis()
	}

	return rdb
}

func (r *Redis) Connection() *redis.Client {
	return r.initializeAndGetRedis()
}

func (r *Redis) Close() {
	r.locker.Lock()
	defer r.locker.Unlock()

	if r.rdb == nil {
		return
	}

	if err := r.rdb.Close(); err != nil {
		logger.Error(context.Background(), "Failed to close Redis", r.configToAttribute().WithError(err))
	}

	r.rdb = nil
}

func (r *Redis) initializeAndGetRedis() *redis.Client {
	db := r.rdb
	if db != nil {
		return db
	}

	r.locker.Lock()
	defer r.locker.Unlock()

	// double-checked locking
	if db := r.rdb; db != nil {
		return db
	}

	logger.Info(context.Background(), "Initializing Redis", r.configToAttribute())

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	rdb := redis.NewClient(r.rdb.Options())

	if err := rdb.Ping(ctx).Err(); err != nil {
		logger.Error(context.Background(), "Failed to ping Redis database", r.configToAttribute().WithError(err))
	} else {
		logger.Info(context.Background(), "Redis initialized", r.configToAttribute())
	}

	r.rdb = rdb
	return rdb
}

func (r *Redis) configToAttribute() attributes.Attributes {
	config := r.opt
	return attributes.Attributes{
		"redis.address":  config.Addr,
		"redis.db":       config.DB,
		"redis.password": "[Masked]",
	}
}
