package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/michaelbunch/go-web/pkg/auth"
	log "github.com/sirupsen/logrus"
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

		email := "asher12000@gmail.com"
		secret := "3h9dnr2u94h7vren8902g7"

		log.Info(au)

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
