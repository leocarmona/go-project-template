package database

import (
	"database/sql"
	"fmt"
	"github.com/leocarmona/go-project-template/internal/infra/logger"
	"github.com/leocarmona/go-project-template/internal/infra/logger/attributes"
	_ "github.com/lib/pq"
	"golang.org/x/net/context"
	"sync"
	"time"
)

type (
	ConnectionStringBuilder func(cfg *SqlConfig) string

	Database struct {
		db                      *sql.DB
		config                  *SqlConfig
		connectionStringBuilder ConnectionStringBuilder
		retries                 []time.Duration
		locker                  sync.Mutex
	}
)

func NewDatabase(cfg *SqlConfig, connectionStringBuilder ConnectionStringBuilder) *Database {
	database := &Database{
		config:                  cfg,
		connectionStringBuilder: connectionStringBuilder,
		retries: []time.Duration{
			250 * time.Millisecond,
			500 * time.Millisecond,
			1000 * time.Millisecond,
			2500 * time.Millisecond,
			5000 * time.Millisecond,
		},
	}

	if !cfg.LazyConnection {
		database.initializeAndGetDB()
	}

	return database
}

func (d *Database) Connection() *sql.DB {
	return d.initializeAndGetDB()
}

func (d *Database) Close() {
	d.locker.Lock()
	defer d.locker.Unlock()

	if d.db == nil {
		return
	}

	if err := d.db.Close(); err != nil {
		logger.Error(context.Background(), fmt.Sprintf("Failed to close database [%s] with connection [%s]", d.config.Database, d.config.ConnectionName), d.configToAttribute().WithError(err))
	}

	d.db = nil
}

func (d *Database) initializeAndGetDB() *sql.DB {
	db := d.db
	if db != nil {
		return db
	}

	d.locker.Lock()
	defer d.locker.Unlock()

	// double-checked locking
	if db = d.db; db != nil {
		return db
	}

	logger.Info(context.Background(), fmt.Sprintf("Initializing database [%s] with connection [%s]", d.config.Database, d.config.ConnectionName), d.configToAttribute())
	var err error

	for retry, duration := range d.retries {
		db, err = sql.Open(d.config.Driver, d.connectionStringBuilder(d.config))

		if err != nil {
			logger.Fatal(context.Background(), fmt.Sprintf("Failed to initialize the database [%s] with connection [%s]", d.config.Database, d.config.ConnectionName), d.configToAttribute().WithError(err))
		}

		if err = d.checkConnection(db); err != nil {
			logger.Warn(context.Background(), fmt.Sprintf("Connection retry [%d]: Database [%s] with connection [%s]", retry+1, d.config.Database, d.config.ConnectionName), d.configToAttribute().WithError(err))
			time.Sleep(duration)
		}
	}

	if err := d.checkConnection(db); err != nil {
		logger.Fatal(context.Background(), fmt.Sprintf("Failed to connect to the database [%s] with connection [%s]", d.config.Database, d.config.ConnectionName), d.configToAttribute().WithError(err))
	}

	db.SetMaxIdleConns(d.config.MinConnections)
	db.SetMaxOpenConns(d.config.MaxConnections)
	db.SetConnMaxLifetime(d.config.ConnectionMaxLifetime)
	db.SetConnMaxIdleTime(d.config.ConnectionMaxIdleTime)

	d.db = db

	logger.Info(context.Background(), fmt.Sprintf("Database [%s] initialized with connection [%s]", d.config.Database, d.config.ConnectionName), d.configToAttribute())

	return db
}

func (d *Database) checkConnection(db *sql.DB) error {
	timeout, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	return db.PingContext(timeout)
}

func (d *Database) configToAttribute() attributes.Attributes {
	config := d.config
	return attributes.Attributes{
		"database.connection_name":          config.ConnectionName,
		"database.driver":                   config.Driver,
		"database.host":                     config.Host,
		"database.port":                     config.Port,
		"database.database":                 config.Database,
		"database.username":                 config.Username,
		"database.password":                 "[Masked]",
		"database.min_connections":          config.MinConnections,
		"database.max_connections":          config.MaxConnections,
		"database.connection_max_lifetime":  config.ConnectionMaxLifetime,
		"database.connection_max_idle_time": config.ConnectionMaxIdleTime,
		"database.lazy_connection":          config.LazyConnection,
	}
}
