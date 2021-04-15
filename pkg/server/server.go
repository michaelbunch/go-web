package server

import (
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gorilla/handlers"
	"github.com/michaelbunch/go-web/pkg/app"
	log "github.com/sirupsen/logrus"
)

// StartFrontendServer starts an instance of the web server
func StartFrontendServer(config *app.Config) {
	router := loadRouter()
	srv := &http.Server{
		Handler:      handlers.LoggingHandler(os.Stdout, router),
		Addr:         ":" + strconv.Itoa(config.Server.Port),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Printf("%s is listening on port %v...", config.AppName, config.Server.Port)
	log.Fatal(srv.ListenAndServe())
}
