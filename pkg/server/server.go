package server

import (
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gorilla/handlers"
	log "github.com/sirupsen/logrus"
)

var (
	port = 3000
)

// StartFrontendServer starts an instance of the web server
func StartFrontendServer() {
	router := loadRouter()
	srv := &http.Server{
		Handler:      handlers.LoggingHandler(os.Stdout, router),
		Addr:         ":" + strconv.Itoa(port),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Printf("Webapp is listening on port %v...", port)
	log.Fatal(srv.ListenAndServe())
}
