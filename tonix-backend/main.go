package main

import (
	"tonix/backend/api/routes"
	"tonix/backend/logging"
)

var logger = logging.LoggerWithOrigin("main.go")

func main() {
	s := routes.HttpServer(":8000")
	logger.Info("Starting server on :8000")
	logger.Fatal(s.ListenAndServe())
}
