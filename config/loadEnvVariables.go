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
	dbURL := os.Getenv("DB_CONNECTION")
    if dbURL == "" {
        fmt.Println("DB_CONNECTION environment variable is not set.")
        return
    }

//		dbHost := os.Getenv("REMOTE_DB_URL")
	ENV = &dbURL // Assuming ENV is used to store the DB_HOST value
	fmt.Printf("This environment state \"DB_HOST=%s\" is still running...\n", *ENV)
	
}
// func LoadEnvVariables() {
//     // Assuming ENV is used to store the DB_HOST value
//     envFilePath := os.Getenv("ENV_FILE_PATH")
//     if envFilePath == "" {
//         fmt.Println("ENV_FILE_PATH environment variable not set")
//         return
//     }

//     if err := godotenv.Load(envFilePath); err != nil {
//         fmt.Println("Error loading environment file:", err)
//         return
//     }

//     dbHost := os.Getenv("REMOTE_DB_URL")
//     ENV = &dbHost
//     fmt.Printf("This environment state \"DB_HOST=%s\" is still running...\n", *ENV)
// }
