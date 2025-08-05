package context

import (
	"database/sql"
	"tonix/backend/env_vars"
)


type SharedState struct {
	Environment env_vars.EnvVars
	DB *sql.DB
}
