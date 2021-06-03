package main

import (
	"os"
	"strconv"

	"github.com/michaelbunch/go-web/pkg/app"
	"github.com/michaelbunch/go-web/pkg/server"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "web",
		Usage: "Run the web server",
		Action: func(c *cli.Context) error {
			run()
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func run() {
	app.LoadConfig()

	port, _ := strconv.Atoi(os.Getenv("PORT"))
	server.StartFrontendServer(server.LoadRouter(), port, os.Getenv("APP_NAME"))
}
