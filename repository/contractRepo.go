package repository

import (
	"PattayaAvenueProperty/db"
	models_Contract "PattayaAvenueProperty/models/Contract"

	"gorm.io/gorm"
)

type ContractRepo struct{}

func NewContractRepo() ContractRepo {
	return ContractRepo{}
}

func ActiveOnlyContract(db *gorm.DB) *gorm.DB {
	return db.Where("is_active = ?", 1)
}

func (repo ContractRepo) GetAllRoomContract() ([]models_Contract.RoomContract, error) {
	var model []models_Contract.RoomContract
	err := ActiveOnlyContract(db.DB).Find(&model).Error
	if err != nil {
		return nil, err
	}
	return model, nil
}

// func (repo ContractRepo) GetAllSecondaryContract() ([]models_Contract.SecondaryContract, error) {
// 	var model []models_Contract.SecondaryContract
// 	err := ActiveOnlyContract(db.DB).Find(&model).Error
// 	if err != nil {
// 		return nil, err
// 	}
// 	return model, nil
// }

func (repo ContractRepo) GetAllPersonContract() ([]models_Contract.PersonContract, error) {
	var model []models_Contract.PersonContract
	err := ActiveOnlyContract(db.DB).Find(&model).Error
	if err != nil {
		return nil, err
	}
	return model, nil
}

func (repo ContractRepo) GetRoomContractByRoomID(id uint) ([]models_Contract.RoomContract, error) {
	var model []models_Contract.RoomContract
	err := db.DB.Where("room_id = ?", id).Find(&model).Error
	if err != nil {
		return nil, err
	}
	return model, nil
}

// func (repo ContractRepo) GetSecondaryContractByRoomContractID(id uint) (models_Contract.SecondaryContract, error) {
// 	var model models_Contract.SecondaryContract
// 	err := ActiveOnlyContract(db.DB).Where("room_contract_id = ?", id).First(&model).Error
// 	if err != nil {
// 		return models_Contract.SecondaryContract{}, err
// 	}
// 	return model, nil
// }

func (repo ContractRepo) GetPersonContractByPersonContractID(id uint) ([]models_Contract.PersonContract, error) {
	var model []models_Contract.PersonContract
	err := db.DB.Where("room_contract_id = ?", id).Find(&model).Error
	if err != nil {
		return nil, err
	}
	return model, nil
}

func (repo ContractRepo) GetRoomContractByID(id uint) (*models_Contract.RoomContract, error) {
	var model models_Contract.RoomContract
	err := db.DB.Where("id = ?", id).First(&model).Error
	if err != nil {
		return nil, err
	}
	return &model, nil
}
