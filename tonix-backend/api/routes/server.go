package routes

import (
	"net/http"
	"tonix/backend/api"
	v1 "tonix/backend/api/routes/v1"
)

func HttpServer(addr string) *http.Server {
	stack := api.Stack()

	http.Handle("GET /api/v1/", stack.AddUniqueHandler(v1.GetIndex));
	http.Handle("POST /api/v1/auth/registration", stack.AddUniqueHandler(v1.Registration));
	http.Handle("POST /api/v1/auth/login", stack.AddUniqueHandler(v1.Login));

	s := &http.Server{
		Addr: addr,
	}

	return s
}
