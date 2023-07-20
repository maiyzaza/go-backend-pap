package main

import (
	"PattayaAvenueProperty/config"
	models_Contract "PattayaAvenueProperty/models/Contract"
	models_Document "PattayaAvenueProperty/models/Document"
	models_Person "PattayaAvenueProperty/models/Person"
	models_Room "PattayaAvenueProperty/models/Room"
	models_Transaction "PattayaAvenueProperty/models/TransactionModel"
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	config.LoadEnvVariables()
	dbREMOTE_DB_URL := os.Getenv("REMOTE_DB_URL")
	dataSourceName := dbREMOTE_DB_URL
	fmt.Println("dataSourceName: ", dataSourceName)

	// var err error
	// DB, err = gorm.Open(mysql.Open(dataSourceName), &gorm.Config{})
	// if err != nil {
	// 	log.Fatal(err)
	// }

	db, err := gorm.Open(mysql.Open(dataSourceName), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// Create the new "PattayaAvenueProperty" database
	// err = db.Exec("CREATE DATABASE IF NOT EXISTS PattayaAvenueProperty").Error
	// if err != nil {
	// 	panic(err)
	// }

	// Perform migrations on the "PattayaAvenueProperty" database
	db.AutoMigrate(
		&models_Person.Person{},
		&models_Room.Place{},
		&models_Room.Building{},
		&models_Room.Floor{},
		&models_Room.Room{},
		&models_Contract.RoomContract{},
		&models_Contract.SecondaryContract{},
		&models_Contract.PersonContract{},
		&models_Document.Document{},
		&models_Document.RoomContractDocument{},
		&models_Person.BankAccount{},
		&models_Person.Contact{},
		&models_Person.Employee{},
		&models_Person.Role{},
		&models_Room.Amenity{},
		&models_Room.RoomAmenity{},
		&models_Room.RoomPicture{},
		&models_Room.RoomPrice{},
		&models_Transaction.InternalTransaction{},
		&models_Transaction.InternalTransactionDocument{},
		&models_Transaction.Transaction{},
	)
}
