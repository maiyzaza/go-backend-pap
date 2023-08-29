package repository

import (
	"PattayaAvenueProperty/db"
	models_Document "PattayaAvenueProperty/models/Document"
	models_Transaction "PattayaAvenueProperty/models/Transaction"

	"gorm.io/gorm"
)

type TransactionRepo struct{}

func NewTransactionRepo() TransactionRepo {
	return TransactionRepo{}
}

func ActiveOnlyTransaction(db *gorm.DB) *gorm.DB {
	return db.Where("is_active = ?", 1)
}

// Get all transactions and internal transactions
func (repo TransactionRepo) GetAllTransaction() ([]models_Transaction.Transaction, error) {
	var model []models_Transaction.Transaction
	err := ActiveOnlyTransaction(db.DB).Find(&model).Error
	if err != nil {
		return nil, err
	}
	return model, nil
}

// get transaction by id
func (repo TransactionRepo) GetTransactionByID(id uint) (*models_Transaction.Transaction, error) {
	var model models_Transaction.Transaction
	err := db.DB.Where("id = ?", id).First(&model).Error
	if err != nil {
		return nil, err
	}
	return &model, nil
}

// get documents
func (repo TransactionRepo) GetAllDocument() ([]models_Document.Document, error) {
	var model []models_Document.Document
	err := ActiveOnlyTransaction(db.DB).Find(&model).Error
	if err != nil {
		return nil, err
	}
	return model, nil
}

// create transaction
func (repo TransactionRepo) CreateTransaction(model *models_Transaction.Transaction) (*models_Transaction.Transaction, error) {
	err := db.DB.Create(&model).Error
	if err != nil {
		return nil, err
	}
	return model, nil
}

// create document
func (repo TransactionRepo) CreateDocument(model *models_Document.Document) (*models_Document.Document, error) {
	err := db.DB.Create(&model).Error
	if err != nil {
		return nil, err
	}
	return model, nil
}

// delete transaction by setting is_active to false
func (repo TransactionRepo) DeleteTransaction(id uint) error {
	err := db.DB.Model(&models_Transaction.Transaction{}).Where("id = ?", id).Update("is_active", false).Error
	if err != nil {
		return err
	}
	return nil
}

// get all trasanction that is deleted
func (repo TransactionRepo) GetAllDeletedTransaction() ([]models_Transaction.Transaction, error) {
	var model []models_Transaction.Transaction
	err := db.DB.Where("is_active = ?", 0).Find(&model).Error
	if err != nil {
		return nil, err
	}
	return model, nil
}
