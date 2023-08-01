package middleware

import (
	"net/http"
	"strings"

	"github.com/dhaaana/go-http-server/utils"
	"github.com/golang-jwt/jwt/v5"
)

const secretKey = "super-shy"

func Authenticate(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Read the "Authorization" header from the request
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			utils.JsonResponse(w, http.StatusUnauthorized, "Authorization header missing", nil)
			return
		}

		// Extract the token from the header
		tokenString := strings.Replace(authHeader, "Bearer ", "", 1)

		// Parse the token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(secretKey), nil
		})
		if err != nil || !token.Valid {
			utils.JsonResponse(w, http.StatusUnauthorized, "Invalid token", nil)
			return
		}

		// Token is valid, proceed to the next handler
		next.ServeHTTP(w, r)
	}
}
