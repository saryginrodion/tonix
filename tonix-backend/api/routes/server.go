package routes

import (
	"net/http"
	"tonix/backend/api"
	"tonix/backend/api/middleware"
	v1 "tonix/backend/api/routes/v1"
)

func HttpServer(addr string) *http.Server {
	stack := api.Stack()
	protectedStack := stack.AddUniqueHandler(middleware.AccessJWTExtractor)

	http.Handle("GET /api/v1/", stack.AddUniqueHandler(v1.GetIndex));

	// auth
	http.Handle("POST /api/v1/auth/registration", stack.AddUniqueHandler(v1.Registration));
	http.Handle("POST /api/v1/auth/login", stack.AddUniqueHandler(v1.Login));

	// profile
	http.Handle("GET /api/v1/profile/self", protectedStack.AddUniqueHandler(v1.ProfileSelf))

	s := &http.Server{
		Addr: addr,
	}

	return s
}
