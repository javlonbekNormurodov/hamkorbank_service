package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

const (
	// DebugMode indicates service mode is debug.
	Debug = "debug"
	// TestMode indicates service mode is test.
	Test = "test"
	// ReleaseMode indicates service mode is release.
	Release = "release"
)

type Config struct {
	ServiceName string
	Environment string // debug, test, release
	Version     string

	ServiceHost string
	HTTPPort    string
	HTTPScheme  string

	DefaultOffset string
	DefaultLimit  string

	PostgresHost     string
	PostgresPort     int
	PostgresUser     string
	PostgresPassword string
	PostgresDatabase string

	PostgresMaxConnections int32

	PosServiceHost string
	PosGRPCPort    string
}

// Load ...
func Load() Config {
	if err := godotenv.Load("/app/.env"); err != nil {
		fmt.Println("No .env file found")
	}

	config := Config{}

	config.ServiceName = cast.ToString(getOrReturnDefaultValue("SERVICE_NAME", "epa_go_api_gateway"))
	config.Environment = cast.ToString(getOrReturnDefaultValue("ENVIRONMENT", Debug))
	config.Version = cast.ToString(getOrReturnDefaultValue("VERSION", "1.0"))

	config.ServiceHost = cast.ToString(getOrReturnDefaultValue("SERVICE_HOST", "localhost"))
	config.HTTPPort = cast.ToString(getOrReturnDefaultValue("HTTP_PORT", ":8000"))
	config.HTTPScheme = cast.ToString(getOrReturnDefaultValue("HTTP_SCHEME", "http"))

	config.PostgresHost = cast.ToString(getOrReturnDefaultValue("POSTGRES_HOST", "localhost"))
	config.PostgresPort = cast.ToInt(getOrReturnDefaultValue("POSTGRES_PORT", 5432))
	config.PostgresUser = cast.ToString(getOrReturnDefaultValue("POSTGRES_USER", "postgres"))
	config.PostgresPassword = cast.ToString(getOrReturnDefaultValue("POSTGRES_PASSWORD", "postgres"))
	config.PostgresDatabase = cast.ToString(getOrReturnDefaultValue("POSTGRES_DATABASE", "postgres"))

	config.PostgresMaxConnections = cast.ToInt32(getOrReturnDefaultValue("POSTGRES_MAX_CONNECTIONS", 30))

	config.DefaultOffset = cast.ToString(getOrReturnDefaultValue("DEFAULT_OFFSET", "0"))
	config.DefaultLimit = cast.ToString(getOrReturnDefaultValue("DEFAULT_LIMIT", "10000000"))

	config.PosServiceHost = cast.ToString(getOrReturnDefaultValue("POS_SERVICE_HOST", "localhost"))
	config.PosGRPCPort = cast.ToString(getOrReturnDefaultValue("POS_SERVICE_GRPC_PORT", ":8000"))

	return config
}

func getOrReturnDefaultValue(key string, defaultValue interface{}) interface{} {
	val, exists := os.LookupEnv(key)

	if exists {
		return val
	}

	return defaultValue
}
