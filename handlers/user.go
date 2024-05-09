package handlers

import "net/http"

func UserGet(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome user get"))
}