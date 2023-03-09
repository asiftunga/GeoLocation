package config

import (
	env "github.com/joho/godotenv"
	"log"
)

const (
	ENV_MODE string = "my"
)

func init() {
	errLoadEnv := env.Load("internal/configs/" + ENV_MODE + ".env")
	if errLoadEnv != nil {
		log.Fatal("Error loading .env file")
	}
}
