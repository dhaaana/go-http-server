package middleware

import "net/http"

func Cors(next http.HandlerFunc) http.HandlerFunc {
	return (http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

		next.ServeHTTP(w, r)
	}))
}
