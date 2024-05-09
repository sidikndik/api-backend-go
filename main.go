package main

import (
	"api-backend-go/routes"
	"fmt"
	"log"
	"net/http"
)

func main(){

	mux := routes.RouteApi()
	
	fmt.Println("Listening port 8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
