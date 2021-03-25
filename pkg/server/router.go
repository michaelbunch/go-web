package server

import (
	"net/http"

	"github.com/gorilla/mux"
)

func loadRouter() *mux.Router {
	r := mux.NewRouter()
	r.StrictSlash(true)

	r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("./frontend/dist/"))))

	return r
}
