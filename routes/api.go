package routes

import (
	"api-backend-go/handlers"
	"net/http"
)


func RouteApi(){

	http.NewServeMux()
	http.HandleFunc("/",handlers.UserGet)
}
