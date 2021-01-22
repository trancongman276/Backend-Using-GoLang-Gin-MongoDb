package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func EnvVar(key string, defaulVal string) string {
	godotenv.Load(".env")
	value := os.Getenv(key)
	if len(value) == 0 {
		return defaulVal
	}
	log.Fatal("Loaded")
	return value
}
