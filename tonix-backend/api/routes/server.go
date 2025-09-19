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

	http.Handle("GET /api/v1/", stack.AddUniqueHandler(v1.GetIndex))

	// auth
	http.Handle("POST /api/v1/auth/registration", stack.AddUniqueHandler(v1.Registration))
	http.Handle("POST /api/v1/auth/login", stack.AddUniqueHandler(v1.Login))
	http.Handle("POST /api/v1/auth/logout", protectedStack.AddUniqueHandler(v1.Logout))
	http.Handle("POST /api/v1/auth/refresh", stack.AddUniqueHandler(middleware.RefreshJWTExtractor).AddUniqueHandler(v1.Refresh))

	// profile
	http.Handle("GET /api/v1/profile/self", protectedStack.AddUniqueHandler(v1.ProfileSelf))
	http.Handle("GET /api/v1/profile/{id}", stack.AddUniqueHandler(v1.Profile))

	// file
	http.Handle("GET /api/v1/file/{id}", stack.AddUniqueHandler(v1.ReadFile))
	http.Handle("POST /api/v1/file", protectedStack.AddUniqueHandler(v1.UploadFile))

	// tags
	http.Handle("GET /api/v1/tag", stack.AddUniqueHandler(v1.SearchTags))

	s := &http.Server{
		Addr: addr,
	}

	return s
}
