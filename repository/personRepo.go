package repository

import (
	// "PattayaAvenueProperty/models"

	"PattayaAvenueProperty/db"
	models_Person "PattayaAvenueProperty/models/Person"
	"fmt"
	// "github.com/google/uuid"
	// "go.mongodb.org/mongo-driver/bson"
)

type PersonRepo struct{}

func NewPersonRepo() PersonRepo {
	return PersonRepo{}
}

func (repo PersonRepo) FindAll() ([]models_Person.Person, error) {
	var model []models_Person.Person
	err := db.DB.Model(&models_Person.Person{}).Preload("BankAccounts").Find(&model).Error
	if err != nil {
		fmt.Println("Error finding records:", err)
		return nil, err
	}
	// var result []models_Person.Person
	// db.DB.Raw("SELECT * FROM persons INNER JOIN bank_accounts on persons.id = bank_accounts.person_id ;").Scan(&result)
	return model, nil
	// ON persons.id=bank_accounts.person_id
}

func (repo PersonRepo) FindById(id uint) (*models_Person.Person, error) {

	var model models_Person.Person
	err := db.DB.Preload("bank_account").Where(models_Person.Person{ID: id}).First(&model).Error

	if err != nil {
		return nil, err
	}
	return &model, nil
}
