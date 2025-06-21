package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Println(".env not found, using system ENV")
	}
}

func GetEnv(key string) string{
	val := os.Getenv(key)
	if val == "" {
		log.Fatalf("%s not set in environment", key)
	}
	return val
}
