package config

import (
	"flag"
	"log"

	"github.com/joho/godotenv"
)

func MustLoadEnv() {
	configPath := flag.String("с", ".env", "path to config file")

	err := godotenv.Load(*configPath)
	if err != nil {
		log.Fatal("LoadEnv", "Error load .env file")
	}
}
