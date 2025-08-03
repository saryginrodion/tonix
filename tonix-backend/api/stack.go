package api

import (
	"tonix/backend/api/context"
	"tonix/backend/api/middleware"
	"tonix/backend/env_vars"

	"github.com/radyshenkya/stackable"
	stackableMiddleware "github.com/radyshenkya/stackable/middleware"
)

var stack *stackable.Stackable[context.SharedState, context.LocalState] = nil

func buildNewStack() stackable.Stackable[context.SharedState, context.LocalState] {
	envVars := env_vars.LoadEnvVars()

	shared := &context.SharedState{
		Environment: *envVars,
	}

	newStack := stackable.Stackable[context.SharedState, context.LocalState]{
		Shared: shared,
	}

	requestIdMW := &stackableMiddleware.RequestIdMiddleware[context.SharedState, context.LocalState]{}
	corsMW := &stackableMiddleware.CORSMiddleware[context.SharedState, context.LocalState]{
		AllowedOrigins:   []string{"http://localhost", "http://127.0.0.1", "http://0.0.0.0"},
		AllowedMethods:   []string{"*"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	}

	newStack.AddHandler(requestIdMW)
	newStack.AddHandler(corsMW)
	newStack.AddHandler(middleware.LoggingMiddleware)

	return newStack
}

/*
Stack() is returning default stackable.Stackable for routes. Working as singleton. This function set ups Shared state + every middleware used by every route
*/
func Stack() *stackable.Stackable[context.SharedState, context.LocalState] {
	if stack == nil {
		newStack := buildNewStack()
		stack = &newStack
	}

	return stack
}
