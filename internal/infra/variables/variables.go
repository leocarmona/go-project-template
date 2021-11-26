package variables

import (
	"os"
	"strconv"
	"time"
)

type variable struct {
	key          string
	defaultValue string
}

var (
	appName                 = &variable{key: "APP_NAME", defaultValue: "go-project-template"}
	appVersion              = &variable{key: "APP_VERSION", defaultValue: "0.0.1"}
	env                     = &variable{key: "ENV", defaultValue: "local"}
	isLambda                = &variable{key: "LAMBDA", defaultValue: "false"}
	logLevel                = &variable{key: "LOG_LEVEL", defaultValue: "debug"}
	serverHost              = &variable{key: "SERVER_HOST", defaultValue: "0.0.0.0"}
	serverPort              = &variable{key: "SERVER_PORT", defaultValue: "5000"}
	serverTimeout           = &variable{key: "SERVER_TIMEOUT", defaultValue: "30"}
	dbHost                  = &variable{key: "DB_HOST", defaultValue: "localhost"}
	dbPort                  = &variable{key: "DB_PORT", defaultValue: "5432"}
	dbName                  = &variable{key: "DB_NAME", defaultValue: "go-project-template"}
	dbUsername              = &variable{key: "DB_USERNAME", defaultValue: "postgres"}
	dbPassword              = &variable{key: "DB_PASSWORD", defaultValue: "postgres123"}
	dbMinConnections        = &variable{key: "DB_MIN_CONNECTIONS", defaultValue: "1"}
	dbMaxConnections        = &variable{key: "DB_MAX_CONNECTIONS", defaultValue: "10"}
	dbConnectionMaxLifeTime = &variable{key: "DB_CONNECTION_MAX_LIFE_TIME", defaultValue: "0"}
	dbConnectionMaxIdleTime = &variable{key: "DB_CONNECTION_MAX_IDLE_TIME", defaultValue: "15"}
	redisEndpoint           = &variable{key: "REDIS_ENDPOINT", defaultValue: "localhost:6379"}
	redisPassword           = &variable{key: "REDIS_PASSWORD", defaultValue: ""}
	redisDB                 = &variable{key: "REDIS_DB", defaultValue: "1"}
)

func AppName() string {
	return get(appName)
}

func AppVersion() string {
	return get(appVersion)
}

func Env() string {
	return get(env)
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

func DBHost() string {
	return get(dbHost)
}

func DBPort() string {
	return get(dbPort)
}

func DBName() string {
	return get(dbName)
}

func DBUsername() string {
	return get(dbUsername)
}

func DBPassword() string {
	return get(dbPassword)
}

func DBMinConnections() int {
	return getInt(dbMinConnections)
}

func DBMaxConnections() int {
	return getInt(dbMaxConnections)
}

func DBConnectionMaxLifeTime() time.Duration {
	return time.Second * time.Duration(getInt(dbConnectionMaxLifeTime))
}

func DBConnectionMaxIdleTime() time.Duration {
	return time.Second * time.Duration(getInt(dbConnectionMaxIdleTime))
}

func RedisEndpoint() string {
	return get(redisEndpoint)
}

func RedisPassword() string {
	return get(redisPassword)
}

func RedisDB() int {
	return getInt(redisDB)
}

func get(env *variable) string {
	value := os.Getenv(env.key)

	if value == "" {
		return env.defaultValue
	}

	return value
}

func getInt(env *variable) int {
	value := get(env)
	intValue, err := strconv.Atoi(value)

	if err != nil {
		//logger.Warn(nil, fmt.Sprintf("Error while trying to convert %s with value %s to int. Using default value %s", env.key, value, env.defaultValue), logger.Details{
		//	"environment_variable": dbMinConnections.key,
		//	"error_value":          value,
		//	"default_value":        env.defaultValue,
		//	"error":                err,
		//})

		intValue, _ = strconv.Atoi(env.defaultValue)
	}

	return intValue
}

func getBool(env *variable) bool {
	value := get(env)
	boolValue, err := strconv.ParseBool(value)

	if err != nil {
		//logger.Warn(nil, fmt.Sprintf("Error while trying to convert %s with value %s to bool. Using default value %s", env.key, value, env.defaultValue), logger.Details{
		//	"environment_variable": dbMinConnections.key,
		//	"error_value":          value,
		//	"default_value":        env.defaultValue,
		//	"error":                err,
		//})

		boolValue, _ = strconv.ParseBool(env.defaultValue)
	}

	return boolValue
}
