package routes

import (
	"net/http"
	"tonix/backend/api"
	v1 "tonix/backend/api/routes/v1"
)

func HttpServer(addr string) *http.Server {
	stack := api.Stack()

	http.Handle("GET /api/v1/", stack.AddUniqueHandler(v1.GetIndex));
	http.Handle("POST /api/v1/test", stack.AddUniqueHandler(v1.PostTestMessage));

	s := &http.Server{
		Addr: addr,
	}

	return s
}
