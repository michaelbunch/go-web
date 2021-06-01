package server

import (
	"net/http"
)

var apiIndexHandler = func(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Go-Web API is working!"))
}
