package app

import (
	"log"

	"github.com/joho/godotenv"
)

// LoadConfig the config from the filesystem
func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
}
