package controllers

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/michaelbunch/go-web/pkg/auth"
)

// AuthUser is the web request for authentication
type AuthUser struct {
	Email  string `json:"email"`
	Secret string `json:"secret"`
}

// LoginHandler accepts a username and password then returns a JWT
var LoginHandler = http.HandlerFunc(
	func(w http.ResponseWriter, r *http.Request) {
		au := new(AuthUser)
		d := json.NewDecoder(r.Body)
		d.Decode(&au)

		email := os.Getenv("API_USER_EMAIL")
		secret := os.Getenv("API_USER_SECRET")

		if au.Email == email && au.Secret == secret {
			token, err := auth.JwtGenerateToken()
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}

			SendJSONResponse(w, token)
		} else {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
		}
	},
)
