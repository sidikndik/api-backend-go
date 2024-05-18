package handlers

import (
	"api-backend-go/database"
	"api-backend-go/model"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type Result struct {
	Name string
	Age  int
}

// get all user
func UserGet(w http.ResponseWriter, r *http.Request) {
	db := database.DBConnect()
	var user []Result

	result := db.Raw("SELECT name, age FROM users").Scan(&user)
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

// store data user
func UserPost(w http.ResponseWriter, r *http.Request) {

	payload, _ := io.ReadAll(r.Body)

	data := []byte(payload)
	user := new(model.User)
	err := json.Unmarshal(data, &user)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("request gagal"))
	}

	db := database.DBConnect()
	result := db.Create(user)

	if result.RowsAffected < 0 {
		log.Fatal(result.Error)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("request gagal"))
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("created"))
}

// update dat user
func UserPut(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	payload, _ := io.ReadAll(r.Body)
	data := []byte(payload)

	var user model.User

	db := database.DBConnect()
	db.First(&user, id)

	err := json.Unmarshal(data, &user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("request feild"))
	}

	db.Save(&user)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("success updated"))
}

// delete data user
func UserDelete(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	db := database.DBConnect()

	db.Delete(&model.User{}, id)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("deleted"))
}
