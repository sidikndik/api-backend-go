package handlers

import (
	"api-backend-go/database"
	"api-backend-go/model"
	"log"
	"net/http"
)

func UserGet(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("welcome user get"))
}

func UserPost(w http.ResponseWriter, r *http.Request) {

	user := new(model.User)

	user.Name = "sidikndik"
	user.Age = 12
	user.Address = "semanan kalideres" 	
	user.Phone = "+6283815698856"

	db := database.DBConnect()
	result := db.Create(user)

	if nil != result.Error{
		log.Fatal(result.Error)
	}

	
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("welcome user post"))
}

func UserPut(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("welcome user update"))
}

func UserDelete(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("welcome user delete"))
}