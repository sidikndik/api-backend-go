package main

import (
	"api-backend-go/routes"
	"log"
	"net/http"
)

func main(){

	routes.RouteApi()

	log.Fatal(http.ListenAndServe(":8080", nil))

}
