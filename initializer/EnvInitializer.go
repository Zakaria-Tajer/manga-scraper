package initializer

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnvironment() {
	err := godotenv.Load("/home/cosmic/client-consumer-api/.env")

	if err != nil {

		log.Fatal("Error loading environment")

	}
}
