package routes

import (
	"api-backend-go/handlers"
	"api-backend-go/middleware"
	"net/http"
)


func RouteApi() http.Handler{ 

	mux := http.NewServeMux()

	finalHandler := http.HandlerFunc(handlers.UserGet)
	mux.Handle("/", middleware.MiddlewareBacicAuth(finalHandler))

	return mux
}
