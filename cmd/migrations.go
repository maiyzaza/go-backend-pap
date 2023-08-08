package main

import (
	models_Contract "PattayaAvenueProperty/models/Contract"
	models_Document "PattayaAvenueProperty/models/Document"
	models_Person "PattayaAvenueProperty/models/Person"
	models_Room "PattayaAvenueProperty/models/Room"
	models_Transaction "PattayaAvenueProperty/models/TransactionModel"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// config.LoadEnvVariables()

	dataSourceName := "doadmin:AVNS_nnRbPmxqVlOi3mYo5LT@tcp(db-mysql-sgp1-88785-do-user-14398363-0.b.db.ondigitalocean.com:25060)/?charset=utf8mb4&parseTime=True&loc=Local"
	// dataSourceName := "root:my-secret-pw@tcp(localhost:1433)/?charset=utf8mb4&parseTime=True&loc=Local"
	// dbREMOTE_DB_URL := os.Getenv("REMOTE_DB_URL")
	// dataSourceName := dbREMOTE_DB_URL
	fmt.Println("dataSourceName: ", dataSourceName)

	db, err := gorm.Open(mysql.Open(dataSourceName), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// Create the new "PattayaAvenueProperty" database
	err = db.Exec("CREATE DATABASE IF NOT EXISTS PattayaAvenueProperty_Staging").Error
	if err != nil {
		panic(err)
	}

	// dataSourceName1 := "root:my-secret-pw@tcp(localhost:1433)/PattayaAvenueProperty?charset=utf8mb4&parseTime=True&loc=Local"
	dataSourceName1 := "doadmin:AVNS_nnRbPmxqVlOi3mYo5LT@tcp(db-mysql-sgp1-88785-do-user-14398363-0.b.db.ondigitalocean.com:25060)/PattayaAvenueProperty_Staging?charset=utf8mb4&parseTime=True&loc=Local"

	db1, err := gorm.Open(mysql.Open(dataSourceName1), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// Perform migrations on the "PattayaAvenueProperty" database
	db1.AutoMigrate(
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
