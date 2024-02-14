package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

const (
	DebugMode   = "debug"
	TestMode    = "test"
	ReleaseMode = "release"

	ClientTypeSuper = "SUPERADMIN"
)

type Config struct {
	Environment string

	ServerHost string
	HTTPPort   string

	SecretKey string

	PostgresHost          string
	PostgresUser          string
	PostgresDatabase      string
	PostgresPassword      string
	PostgresPort          int
	PostgresMaxConnection int32

	DefaultOffset int
	DefaultLimit  int
}

func Load() Config {

	if err := godotenv.Load(".env"); err != nil {
		fmt.Println("No .env file found")
	}

	cfg := Config{}

	cfg.DefaultOffset = 0
	cfg.DefaultLimit = 10

	cfg.Environment = cast.ToString(getOrReturnDefaultValue("ENVIRONMENT", DebugMode))

	cfg.ServerHost = cast.ToString(getOrReturnDefaultValue("SERVER_HOST", "https://lms-vuny.onrender.com"))
	cfg.HTTPPort = cast.ToString(getOrReturnDefaultValue("HTTP_PORT", ":8081"))

	// cfg.ServerHost = cast.ToString(getOrReturnDefaultValue("SERVER_HOST", "localhost"))
	// cfg.HTTPPort = cast.ToString(getOrReturnDefaultValue("HTTP_PORT", ":8081"))

	cfg.SecretKey = cast.ToString(getOrReturnDefaultValue("SECRET_KEY", "=huiowp34"))

	cfg.PostgresHost = cast.ToString(getOrReturnDefaultValue("POSTGRES_HOST", "dpg-cn279egl5elc73ea1tmg-a"))
	cfg.PostgresUser = cast.ToString(getOrReturnDefaultValue("POSTGRES_USER", "lms_38y0_user"))
	cfg.PostgresDatabase = cast.ToString(getOrReturnDefaultValue("POSTGRES_DATABASE", "lms_38y0"))
	cfg.PostgresPassword = cast.ToString(getOrReturnDefaultValue("POSTGRES_PASSWORD", "WJxKaw4BGrzNxg1zfk4bRyAbYcEAunm1"))
	cfg.PostgresPort = cast.ToInt(getOrReturnDefaultValue("POSTGRES_PORT", 5432))

	// cfg.PostgresHost = cast.ToString(getOrReturnDefaultValue("POSTGRES_HOST", "localhost"))
	// cfg.PostgresUser = cast.ToString(getOrReturnDefaultValue("POSTGRES_USER", "postgres"))
	// cfg.PostgresDatabase = cast.ToString(getOrReturnDefaultValue("POSTGRES_DATABASE", "generic"))
	// cfg.PostgresPassword = cast.ToString(getOrReturnDefaultValue("POSTGRES_PASSWORD", "1234"))
	// cfg.PostgresPort = cast.ToInt(getOrReturnDefaultValue("POSTGRES_PORT", 5432))

	cfg.PostgresMaxConnection = cast.ToInt32(getOrReturnDefaultValue("POSTGRES_PORT", 30))

	cfg.DefaultOffset = cast.ToInt(getOrReturnDefaultValue("OFFSET", 0))
	cfg.DefaultLimit = cast.ToInt(getOrReturnDefaultValue("LIMIT", 10))

	return cfg
}

func getOrReturnDefaultValue(key string, defaultValue interface{}) interface{} {
	val, exists := os.LookupEnv(key)

	if exists {
		return val
	}

	return defaultValue
}
