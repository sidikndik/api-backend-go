package routes

import (
	"api-backend-go/handlers"
	"api-backend-go/middleware"
	"net/http"
)

func RouteApi() http.Handler {

	mux := http.NewServeMux()

	mux.Handle("GET /user", middleware.MiddlewareBacicAuth(http.HandlerFunc(handlers.UserGet)))
	mux.Handle("POST /user", middleware.MiddlewareBacicAuth(http.HandlerFunc(handlers.UserPost)))
	mux.Handle("PUT /user/{id}", middleware.MiddlewareBacicAuth(http.HandlerFunc(handlers.UserPut)))
	mux.Handle("DELETE /user/{id}", middleware.MiddlewareBacicAuth(http.HandlerFunc(handlers.UserDelete)))

	return mux
}
