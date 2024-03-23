package config

import (
	"flag"
	"log"

	"github.com/joho/godotenv"
)

// type Config struct {
// 	ServerPort string `env:"PORT" env-default:"3000"`
// 	GRPCPort   string `env:"PORT" env-default:"8080"`
// }

// func MustLoad() *Config {
// 	configPath := flag.String("cfg", "", "path to config file")

// 	if *configPath == "" {
// 		log.Fatal("config path is nil")
// 	}

// 	var cfg Config
// 	if err := cleanenv.ReadConfig(*configPath, &cfg); err != nil {
// 		log.Fatal(err)
// 	}

// 	return &cfg
// }

// TODO: fix
func MustLoadEnv() {
	configPath := flag.String("cfg", ".env", "path to config file")

	err := godotenv.Load(*configPath)
	if err != nil {
		log.Fatal("Error load .env file")
	}
}
