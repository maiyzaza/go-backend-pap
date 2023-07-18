package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var ENV *string

func LoadEnvVariables() {
	if err := godotenv.Load(".env"); err != nil {
		fmt.Println("No environment file found")
	}

	dbHost := os.Getenv("DB_HOST")
	ENV = &dbHost // Assuming ENV is used to store the DB_HOST value
	fmt.Printf("This environment state \"DB_HOST=%s\" is still running...\n", *ENV)
}
