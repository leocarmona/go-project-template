package variables

import (
	"context"
	"fmt"
	"github.com/leocarmona/go-project-template/internal/infra/logger"
	"github.com/leocarmona/go-project-template/internal/infra/logger/attributes"
	"os"
	"strconv"
	"time"
)

type variable struct {
	key          string
	defaultValue string
}

var (
	serviceName                  = &variable{key: "SERVICE_NAME", defaultValue: "go-project-template"}
	serviceVersion               = &variable{key: "SERVICE_VERSION", defaultValue: "0.0.1"}
	environment                  = &variable{key: "ENVIRONMENT", defaultValue: "local"}
	isLambda                     = &variable{key: "LAMBDA", defaultValue: "false"}
	logLevel                     = &variable{key: "LOG_LEVEL", defaultValue: "debug"}
	serverHost                   = &variable{key: "SERVER_HOST", defaultValue: "0.0.0.0"}
	serverPort                   = &variable{key: "SERVER_PORT", defaultValue: "5000"}
	serverTimeout                = &variable{key: "SERVER_TIMEOUT", defaultValue: "30"}
	dbReadHost                   = &variable{key: "DB_READ_HOST", defaultValue: "localhost"}
	dbReadPort                   = &variable{key: "DB_READ_PORT", defaultValue: "5432"}
	dbReadName                   = &variable{key: "DB_READ_NAME", defaultValue: "go-project-template"}
	dbReadUsername               = &variable{key: "DB_READ_USERNAME", defaultValue: "postgres"}
	dbReadPassword               = &variable{key: "DB_READ_PASSWORD", defaultValue: "postgres123"}
	dbReadLazyConnection         = &variable{key: "DB_READ_LAZY_CONNECTION", defaultValue: "true"}
	dbReadMinConnections         = &variable{key: "DB_READ_MIN_CONNECTIONS", defaultValue: "2"}
	dbReadMaxConnections         = &variable{key: "DB_READ_MAX_CONNECTIONS", defaultValue: "10"}
	dbReadConnectionMaxLifeTime  = &variable{key: "DB_READ_CONNECTION_MAX_LIFE_TIME", defaultValue: "900"}
	dbReadConnectionMaxIdleTime  = &variable{key: "DB_READ_CONNECTION_MAX_IDLE_TIME", defaultValue: "60"}
	dbWriteHost                  = &variable{key: "DB_WRITE_HOST", defaultValue: "localhost"}
	dbWritePort                  = &variable{key: "DB_WRITE_PORT", defaultValue: "5432"}
	dbWriteName                  = &variable{key: "DB_WRITE_NAME", defaultValue: "go-project-template"}
	dbWriteUsername              = &variable{key: "DB_WRITE_USERNAME", defaultValue: "postgres"}
	dbWritePassword              = &variable{key: "DB_WRITE_PASSWORD", defaultValue: "postgres123"}
	dbWriteLazyConnection        = &variable{key: "DB_WRITE_LAZY_CONNECTION", defaultValue: "true"}
	dbWriteMinConnections        = &variable{key: "DB_WRITE_MIN_CONNECTIONS", defaultValue: "2"}
	dbWriteMaxConnections        = &variable{key: "DB_WRITE_MAX_CONNECTIONS", defaultValue: "10"}
	dbWriteConnectionMaxLifeTime = &variable{key: "DB_WRITE_CONNECTION_MAX_LIFE_TIME", defaultValue: "900"}
	dbWriteConnectionMaxIdleTime = &variable{key: "DB_WRITE_CONNECTION_MAX_IDLE_TIME", defaultValue: "60"}
	redisHost                    = &variable{key: "REDIS_HOST", defaultValue: "localhost"}
	redisPort                    = &variable{key: "REDIS_PORT", defaultValue: "6379"}
	redisPassword                = &variable{key: "REDIS_PASSWORD", defaultValue: ""}
	redisDB                      = &variable{key: "REDIS_DB", defaultValue: "1"}
	redisLazyConnection          = &variable{key: "REDIS_LAZY_CONNECTION", defaultValue: "true"}
)

func ServiceName() string {
	return get(serviceName)
}

func ServiceVersion() string {
	return get(serviceVersion)
}

func Environment() string {
	return get(environment)
}

func IsLambda() bool {
	return getBool(isLambda)
}

func LogLevel() string {
	return get(logLevel)
}

func ServerHost() string {
	return get(serverHost)
}

func ServerPort() int {
	return getInt(serverPort)
}

func ServerTimeout() int {
	return getInt(serverTimeout)
}

func DBReadHost() string {
	return get(dbReadHost)
}

func DBReadPort() string {
	return get(dbReadPort)
}

func DBReadName() string {
	return get(dbReadName)
}

func DBReadUsername() string {
	return get(dbReadUsername)
}

func DBReadPassword() string {
	return get(dbReadPassword)
}

func DBReadLazyConnection() bool {
	return getBool(dbReadLazyConnection)
}

func DBReadMinConnections() int {
	return getInt(dbReadMinConnections)
}

func DBReadMaxConnections() int {
	return getInt(dbReadMaxConnections)
}

func DBReadConnectionMaxLifeTime() time.Duration {
	return time.Second * time.Duration(getInt(dbReadConnectionMaxLifeTime))
}

func DBReadConnectionMaxIdleTime() time.Duration {
	return time.Second * time.Duration(getInt(dbReadConnectionMaxIdleTime))
}

func DBWriteHost() string {
	return get(dbWriteHost)
}

func DBWritePort() string {
	return get(dbWritePort)
}

func DBWriteName() string {
	return get(dbWriteName)
}

func DBWriteUsername() string {
	return get(dbWriteUsername)
}

func DBWritePassword() string {
	return get(dbWritePassword)
}

func DBWriteLazyConnection() bool {
	return getBool(dbWriteLazyConnection)
}

func DBWriteMinConnections() int {
	return getInt(dbWriteMinConnections)
}

func DBWriteMaxConnections() int {
	return getInt(dbWriteMaxConnections)
}

func DBWriteConnectionMaxLifeTime() time.Duration {
	return time.Second * time.Duration(getInt(dbWriteConnectionMaxLifeTime))
}

func DBWriteConnectionMaxIdleTime() time.Duration {
	return time.Second * time.Duration(getInt(dbWriteConnectionMaxIdleTime))
}

func RedisHost() string {
	return get(redisHost)
}

func RedisPort() int {
	return getInt(redisPort)
}

func RedisPassword() string {
	return get(redisPassword)
}

func RedisDB() int {
	return getInt(redisDB)
}

func RedisLazyConnection() bool {
	return getBool(redisLazyConnection)
}

func get(env *variable) string {
	value := os.Getenv(env.key)

	if len(value) == 0 {
		return env.defaultValue
	}

	return value
}

func getInt(env *variable) int {
	value := get(env)
	intValue, err := strconv.Atoi(value)

	if err != nil {
		logFatal(env, "int", value, err)
	}

	return intValue
}

func getBool(env *variable) bool {
	value := get(env)
	boolValue, err := strconv.ParseBool(value)

	if err != nil {
		logFatal(env, "bool", value, err)
	}

	return boolValue
}

func logFatal(env *variable, varType string, returnedValue string, err error) {
	logger.Fatal(context.Background(), fmt.Sprintf("Error while trying to convert variable [%s] with value [%s] to [%s].", env.key, returnedValue, varType), attributes.Attributes{
		"variable.key":           env.key,
		"variable.type":          varType,
		"variable.value":         returnedValue,
		"variable.default_value": env.defaultValue,
	}.WithError(err))
}
