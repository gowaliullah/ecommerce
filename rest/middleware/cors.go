package middleware

import "net/http"

func Cors(next *http.ServeMux) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Habib") // preflight request wiht OPTIONS
		w.Header().Set("Content-Type", "application/json")

		next.ServeHTTP(w, r)
	})
}
