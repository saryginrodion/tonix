package env_vars

import (
	"os"
	"strconv"
	"time"
	"tonix/backend/logging"
)

var logger = logging.LoggerWithOrigin("env_vars.go")

type EnvVars struct {
	JWT_SECRET                    string
	JWT_ACCESS_COOLDOWN_DURATION  time.Duration
	JWT_REFRESH_COOLDOWN_DURATION time.Duration
	REDIS_CONNECTION_URL          string
	POSTGRES_CONNECTION_URL       string
	UPLOADS_DIRECTORY             string
	UPLOADS_MAX_SIZE_MB           int
}

func loadEnvVar(key string) string {
	val, isExists := os.LookupEnv(key)

	if !isExists {
		logger.Fatalln("Environment variable is unset: ", key)
	}

	return val
}

func ParseInt(v string) int {
	res, err := strconv.Atoi(v)
	if err != nil {
		logger.Panicln("Failed to convert value ", v, " to int")
	}

	return res
}

func ParseDuration(v string) time.Duration {
	dur, err := time.ParseDuration(v)

	if err != nil {
		logger.Panicln("Failed to parse duration: ", v)
	}

	return dur
}

func LoadEnvVars() *EnvVars {
	return &EnvVars{
		POSTGRES_CONNECTION_URL:       loadEnvVar("POSTGRES_CONNECTION_URL"),
		REDIS_CONNECTION_URL:          loadEnvVar("REDIS_CONNECTION_URL"),
		JWT_SECRET:                    loadEnvVar("JWT_SECRET"),
		JWT_ACCESS_COOLDOWN_DURATION:  ParseDuration(loadEnvVar("JWT_ACCESS_COOLDOWN_DURATION")),
		JWT_REFRESH_COOLDOWN_DURATION: ParseDuration(loadEnvVar("JWT_REFRESH_COOLDOWN_DURATION")),
		UPLOADS_DIRECTORY:             loadEnvVar("UPLOADS_DIRECTORY"),
		UPLOADS_MAX_SIZE_MB:           ParseInt(loadEnvVar("UPLOADS_MAX_SIZE_MB")),
	}
}
