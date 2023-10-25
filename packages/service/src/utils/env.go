package utils

import (
	"fmt"

	"github.com/joho/godotenv"
)

func LoadEnvs() {
	err := godotenv.Load(".env.local")
	if err != nil {
		fmt.Println("Error loading .env.local file")
	}
}
