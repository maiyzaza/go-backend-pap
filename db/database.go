package db

import (
	"PattayaAvenueProperty/config"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	config.LoadEnvVariables()
	dbREMOTE_DB_URL := os.Getenv("REMOTE_DB_URL")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	dataSourceName1 := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbUser, dbPassword, dbHost, dbPort, dbName)
	dataSourceName := dbREMOTE_DB_URL
	fmt.Println("dataSourceName: ", dataSourceName)
	fmt.Println("dataSourceName1: ", dataSourceName1)

	var err error
	DB, err = gorm.Open(mysql.Open(dataSourceName), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
}
