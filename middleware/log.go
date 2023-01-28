package middleware

import (
	"net/http"

	"github.com/dhaaana/go-http-server/utils"
)

// when a request go through this function, this function will Log the information first
// the function on the parameter will run after that
func Logging(next http.HandlerFunc) http.HandlerFunc {
	return (http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		utils.LogInfo("Received request to", r.URL.Path)
		next.ServeHTTP(w, r)
	}))
}
