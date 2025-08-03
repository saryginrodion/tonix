package context

import (
	"tonix/backend/env_vars"
)


type SharedState struct {
	Environment env_vars.EnvVars
}
