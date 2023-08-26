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
	err := ActiveOnlyPerson(db.DB).Where("id = ?", personID).Find(&models_Person.Person{}).Preload("BankAccounts").First(&model).Error
	if err != nil {
		fmt.Println("Error finding records:", err)
		return nil, err
	}
	return &model, nil
}

func (repo PersonRepo) GetAllContact() ([]models_Person.Contact, error) {
	var model []models_Person.Contact
	err := ActiveOnlyPerson(db.DB).Find(&models_Person.Contact{}).Find(&model).Error
	if err != nil {
		fmt.Println("Error finding records:", err)
		return nil, err
	}
	return model, nil
}

// create person and edit person
func (repo PersonRepo) CreatePerson(person *models_Person.Person) (*models_Person.Person, error) {
	err := db.DB.Create(&person).Error
	if err != nil {
		fmt.Println("Error creating records:", err)
		return nil, err
	}
	return person, nil
}

func (repo PersonRepo) UpdatePerson(person *models_Person.Person) (*models_Person.Person, error) {
	err := db.DB.Save(&person).Error
	if err != nil {
		fmt.Println("Error updating records:", err)
		return nil, err
	}
	return person, nil
}

// create contact and edit contact and delete contact
func (repo PersonRepo) CreateContact(contact *models_Person.Contact) (*models_Person.Contact, error) {
	err := db.DB.Create(&contact).Error
	if err != nil {
		fmt.Println("Error creating records:", err)
		return nil, err
	}
	return contact, nil
}

func (repo PersonRepo) UpdateContact(contact *models_Person.Contact) (*models_Person.Contact, error) {
	err := db.DB.Save(&contact).Error
	if err != nil {
		fmt.Println("Error updating records:", err)
		return nil, err
	}
	return contact, nil
}

func (repo PersonRepo) DeleteContact(contact *models_Person.Contact) (*models_Person.Contact, error) {
	err := db.DB.Model(&models_Person.Contact{}).Where("id = ?", contact.ID).Update("is_active", 0).Error
	if err != nil {
		fmt.Println("Error deleting records:", err)
		return nil, err
	}
	return contact, nil
}

// create bank account and edit bank account
func (repo PersonRepo) CreateBankAccount(bankAccount *models_Person.BankAccount) (*models_Person.BankAccount, error) {
	err := db.DB.Create(&bankAccount).Error
	if err != nil {
		fmt.Println("Error creating records:", err)
		return nil, err
	}
	return bankAccount, nil
}

func (repo PersonRepo) UpdateBankAccount(bankAccount *models_Person.BankAccount) (*models_Person.BankAccount, error) {
	err := db.DB.Save(&bankAccount).Error
	if err != nil {
		fmt.Println("Error updating records:", err)
		return nil, err
	}
	return bankAccount, nil
}
