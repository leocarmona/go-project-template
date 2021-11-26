package database

import "time"

type SqlConfig struct {
	ConnectionName        string
	Driver                string
	Host                  string
	Port                  string
	Database              string
	Username              string
	Password              string
	MinConnections        int
	MaxConnections        int
	ConnectionMaxLifetime time.Duration // 0, connections are reused forever
	ConnectionMaxIdleTime time.Duration
	LazyConnection        bool
}
