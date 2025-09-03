package context

import (
	"tonix/backend/env_vars"

	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
)

type SharedState struct {
	Environment  env_vars.EnvVars
	DB           *sqlx.DB
	RedisClient  *redis.Client
}
