package server

import (
	"net/http"

	"github.com/gorilla/mux"
)

func loadRouter() *mux.Router {
	r := mux.NewRouter()
	r.StrictSlash(true)

	// Frontend handler
	r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("./frontend/dist/"))))

	// API handlers
	r.HandleFunc("/api", jwtTokenCheck(apiIndexHandler)).Methods(http.MethodGet)

	return r
}
