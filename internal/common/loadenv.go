package common

import (
	"github.com/joho/godotenv"
	"log"
)

func LoadEnv() error {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}
	return err
}
