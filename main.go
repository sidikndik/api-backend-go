package main

import (
	config "api-backend-go/configs"
	"api-backend-go/routes"
	"fmt"
	"log"
	"net/http"
	"os"
)

func init() {
	config.Environment()
}

func main() {

	mux := routes.RouteApi()

	fmt.Println("Listening port 8080")
	fmt.Println(os.Getenv("DATABASE_NAME"))

	log.Fatal(http.ListenAndServe(":8080", mux))

}
