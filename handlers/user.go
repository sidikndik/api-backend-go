package handlers

import (
	"api-backend-go/database"
	"api-backend-go/model"
	"encoding/json"
	"log"
	"net/http"
)

type Result struct {
	Name string
	Age  int
}

func UserGet(w http.ResponseWriter, r *http.Request) {
	db := database.DBConnect()
	var user []Result

	result := db.Raw("SELECT name, age FROM users LIMIT 1").Scan(&user)
	if result.Error != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("data not found"))
	}
	// if err := db.Find(&user).Error; err != nil {
	// }
	output, err := json.Marshal(user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("internal server error"))
	}
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(output)
}

func UserPost(w http.ResponseWriter, r *http.Request) {

	user := new(model.User)

	user.Name = "sidikndik"
	user.Age = 12
	user.Address = "semanan kalideres"
	user.Phone = "+6283815698856"

	db := database.DBConnect()
	result := db.Create(user)

	if result.RowsAffected < 0 {
		log.Fatal(result.Error)

		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("gagal menyimpan data"))
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("created"))
}

func UserPut(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("welcome user update"))
}

func UserDelete(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("welcome user delete"))
}
