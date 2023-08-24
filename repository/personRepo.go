package repository

import (
	// "PattayaAvenueProperty/models"

	"PattayaAvenueProperty/db"
	models_Person "PattayaAvenueProperty/models/Person"
	"fmt"

	"gorm.io/gorm"
	// "github.com/google/uuid"
	// "go.mongodb.org/mongo-driver/bson"
)

type PersonRepo struct{}

func NewPersonRepo() PersonRepo {
	return PersonRepo{}
}

func ActiveOnlyPerson(db *gorm.DB) *gorm.DB {
	return db.Where("is_active = ?", 1)
}

func (repo PersonRepo) FindAll() ([]models_Person.Person, error) {
	var model []models_Person.Person
	err := ActiveOnlyPerson(db.DB).Find(&models_Person.Person{}).Preload("BankAccounts").Find(&model).Error
	if err != nil {
		fmt.Println("Error finding records:", err)
		return nil, err
	}
	return model, nil
}

func (repo PersonRepo) FindPersonById(personID uint) (*models_Person.Person, error) {
	var model models_Person.Person
	err := ActiveOnlyPerson(db.DB).Find(&models_Person.Person{}).Preload("BankAccounts").First(&model).Error
	if err != nil {
		fmt.Println("Error finding records:", err)
		return nil, err
	}
	return &model, nil
}
