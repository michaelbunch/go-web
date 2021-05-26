package server

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/michaelbunch/go-web/pkg/app"
	"github.com/michaelbunch/go-web/pkg/auth"
)

func loadRouter() *mux.Router {
	r := mux.NewRouter()
	r.StrictSlash(true)

	// Frontend handler
	r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("./frontend/dist/"))))

	// API handlers
	r.HandleFunc("/api/auth", auth.AuthHandler)).Methods(http.MethodPost)
	r.HandleFunc("/api", jwtTokenCheck(apiIndexHandler)).Methods(http.MethodGet)

	return r
}
