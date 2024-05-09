package routes

import (
	"api-backend-go/handlers"
	"api-backend-go/middleware"
	"net/http"
)


func RouteApi() http.Handler{ 

	mux := http.NewServeMux()

	mux.Handle("/", middleware.MiddlewareBacicAuth(http.HandlerFunc(handlers.UserGet)))

	return mux
}
