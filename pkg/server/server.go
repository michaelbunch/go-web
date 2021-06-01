package server

import (
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

// StartFrontendServer starts an instance of the web server
func StartFrontendServer(router *mux.Router, port int, appName string) {
	srv := &http.Server{
		Handler:      handlers.LoggingHandler(os.Stdout, router),
		Addr:         ":" + strconv.Itoa(port),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Printf("%s is listening on port %v...", appName, port)
	log.Fatal(srv.ListenAndServe())
}
