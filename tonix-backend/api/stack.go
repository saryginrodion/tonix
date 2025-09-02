package api

import (
	"tonix/backend/api/context"
	"tonix/backend/api/middleware"
	"tonix/backend/database"
	"tonix/backend/env_vars"
	"tonix/backend/logging"

	"github.com/saryginrodion/stackable"
	stackableMiddleware "github.com/saryginrodion/stackable/middleware"
	"github.com/sirupsen/logrus"
)

var log = logging.LoggerWithOrigin("stack.go")

var stack *stackable.Stackable[context.SharedState, context.LocalState] = nil

func buildNewStack() stackable.Stackable[context.SharedState, context.LocalState] {
	envVars := env_vars.LoadEnvVars()

	dbConnection, err := database.Connect(envVars.POSTGRES_CONNECTION_URL)
	if err != nil {
		log.Fatal("Failed to connect to db: ", err.Error())
	}

	shared := &context.SharedState{
		Environment: *envVars,
		DB:          dbConnection,
	}

	newStack := stackable.NewStackable[context.SharedState, context.LocalState](
		shared,
	)

	newStack.SetLogLevel(logrus.DebugLevel)

	requestIdMW := &stackableMiddleware.RequestIdMiddleware[context.SharedState, context.LocalState]{}
	corsMW := &stackableMiddleware.CORSMiddleware[context.SharedState, context.LocalState]{
		AllowedOrigins:   []string{"http://localhost", "http://127.0.0.1", "http://0.0.0.0"},
		AllowedMethods:   []string{"*"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	}

	newStack.AddHandler(middleware.LoggingMiddleware)
	newStack.AddHandler(middleware.ErrorsHandlerMiddleware)
	newStack.AddHandler(requestIdMW)
	newStack.AddHandler(corsMW)

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
