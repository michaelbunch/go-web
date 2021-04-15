package main

import (
	"os"

	"github.com/michaelbunch/go-web/pkg/app"
	"github.com/michaelbunch/go-web/pkg/server"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "web",
		Usage: "Run the web server",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "config",
				Aliases: []string{"c"},
				Value:   "config.yaml",
				Usage:   "Load configuration from `FILE`",
			},
		},
		Action: func(c *cli.Context) error {
			run(c.String("config"))
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func run(configPath string) {
	c, err := app.LoadConfig(configPath)
	if err != nil {
		log.Fatalln(err)
	}

	server.StartFrontendServer(c)
}
