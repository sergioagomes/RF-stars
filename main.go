package main

import (
	"net/http"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("not implemented yet"))
}

func SecureHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("voce pode passar"))
}

func PublicHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("voce pode ver esse carai"))
}

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("X-Api-Token")
		if len(token) == 0 {
			http.Error(w, "not authorized", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	}
}

func main() {
	http.HandleFunc("/api/auth", LoginHandler)
	http.HandleFunc("/api/public", PublicHandler)
	http.HandleFunc("/api/secure", AuthMiddleware(SecureHandler))
	http.ListenAndServe(":3000", nil)
}
