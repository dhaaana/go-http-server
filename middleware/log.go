package middleware

import (
	"net/http"

	"github.com/dhaaana/go-http-server/utils"
)

func Logging(next http.HandlerFunc) http.HandlerFunc {
	return (http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		utils.LogInfo("Received", r.Method, "request to", r.URL.Path)
		next.ServeHTTP(w, r)
	}))
}
