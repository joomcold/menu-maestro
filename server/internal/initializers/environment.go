package initializers

import (
	"github.com/joho/godotenv"
)

func Environment() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
}
