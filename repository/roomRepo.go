package repository

import (
	"PattayaAvenueProperty/db"
	models_Room "PattayaAvenueProperty/models/Room"
	"fmt"
)

type RoomRepo struct{}

func NewRoomRepo() RoomRepo {
	return RoomRepo{}
}

func (repo RoomRepo) GetAllPlace() ([]models_Room.Place, error) {
	var model []models_Room.Place
	err := db.DB.Preload("Buildings.Floors").Find(&model).Error
	if err != nil {
		fmt.Println("Error finding records:", err)
		return nil, err
	}
	return model, nil
}

func (repo RoomRepo) CreatePlace(place models_Room.Place) (*models_Room.Place, error) {
	err := db.DB.Create(&place).Error
	if err != nil {
		return nil, err
	}
	return &place, nil
}
