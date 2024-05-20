package middleware

import (
	"log"
	"net/http"
)

var (
	USERNAME = "sidikndik"
	PASSWORD = "12345"
)

func MiddlewareBacicAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		log.Print("Executing MiddlewareBacicAuth")
		username, password, ok := r.BasicAuth()
		if !ok {
			w.WriteHeader(401)
			w.Write([]byte(`something went wrong`))
			return
		}

		isValid := (username == USERNAME) && (password == PASSWORD)
		if !isValid {
			w.WriteHeader(401)
			w.Write([]byte(`wrong username/password`))
			return
		}

		next.ServeHTTP(w, r)
	})
}
