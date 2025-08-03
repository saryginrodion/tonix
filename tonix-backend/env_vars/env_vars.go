package env_vars

import (
	"os"
	"tonix/backend/logging"
)

var logger = logging.Logger("env_vars.go")

type EnvVars struct {
	POSTGRES_CONNECTION_URL string
}

func loadEnvVar(key string) string {
	val, isExists := os.LookupEnv(key);

	if !isExists {
		logger.Fatalf("Environment variable is unset")
	}

	return val
}

func LoadEnvVars() *EnvVars {
	return &EnvVars{
		POSTGRES_CONNECTION_URL: loadEnvVar("POSTGRES_CONNECTION_URL"),
	}
}
