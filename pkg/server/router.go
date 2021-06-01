package server

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/michaelbunch/go-web/pkg/app"
	"github.com/michaelbunch/go-web/pkg/controllers"
)

// LoadRouter prepares and returns the application router
func LoadRouter(config *app.Config) *mux.Router {
	r := mux.NewRouter()
	r.StrictSlash(true)

	// API handler
	r.HandleFunc("/api", jwtTokenCheck(apiIndexHandler)).Methods(http.MethodGet)
	r.HandleFunc("/api/login", controllers.LoginHandler).Methods(http.MethodPost)

	// Frontend handler
	r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("./frontend/dist/"))))

	return r
}
