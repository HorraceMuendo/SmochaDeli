package env

import (
	"log"

	"github.com/joho/godotenv"
)

func EnvironmentVar() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("failed to load environment variables")
	}
}
