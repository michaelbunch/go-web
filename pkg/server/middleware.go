package server

import (
	"net/http"
	"strings"

	services "github.com/michaelbunch/go-web/pkg/auth"
)

func jwtTokenCheck(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			auth := r.Header.Get("authorization")
			if auth == "" {
				http.Error(w, "Missing authorization token.", http.StatusBadRequest)
				return
			}

			parsed := strings.Split(auth, " ")
			if len(parsed) != 2 {
				http.Error(w, "Authorization header is malformed.", http.StatusBadRequest)
				return
			}

			tokenStatus := services.JwtValidateToken(parsed[1])
			if !tokenStatus {
				http.Error(w, "Invalid token provided.", http.StatusUnauthorized)
				return
			}

			next.ServeHTTP(w, r)
		},
	)
}
