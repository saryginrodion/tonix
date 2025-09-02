package context

import (
	"tonix/backend/env_vars"

	"github.com/jmoiron/sqlx"
)

type SharedState struct {
	Environment env_vars.EnvVars
	DB          *sqlx.DB
}
