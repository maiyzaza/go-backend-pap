package repository

import (
	"PattayaAvenueProperty/db"
	models_Person "PattayaAvenueProperty/models/Person"
	models_Room "PattayaAvenueProperty/models/Room"
	"fmt"
)

type RoomRepo struct{}

func NewRoomRepo() RoomRepo {
	return RoomRepo{}
}

// Get all places, buildings, floors, rooms
func (repo RoomRepo) GetAllPlace() ([]models_Room.Place, error) {
	var model []models_Room.Place
	err := db.DB.Find(&model).Error
	if err != nil {
		fmt.Println("Error finding records:", err)
		return nil, err
	}
	return model, nil
}

func (repo RoomRepo) GetAllBuilding() ([]models_Room.Building, error) {
	var model []models_Room.Building
	err := db.DB.Find(&model).Error
	if err != nil {
		fmt.Println("Error finding records:", err)
		return nil, err
	}
	return model, nil
}

func (repo RoomRepo) GetAllFloor() ([]models_Room.Floor, error) {
	var model []models_Room.Floor
	err := db.DB.Find(&model).Error
	if err != nil {
		fmt.Println("Error finding records:", err)
		return nil, err
	}
	return model, nil
}

func (repo RoomRepo) GetAllRoom() ([]models_Room.Room, error) {
	var model []models_Room.Room
	err := db.DB.Find(&model).Error
	if err != nil {
		fmt.Println("Error finding records:", err)
		return nil, err
	}
	return model, nil
}

// Create place, building, floor, room
func (repo RoomRepo) CreatePlace(place models_Room.Place) (*models_Room.Place, error) {
	err := db.DB.Create(&place).Error
	if err != nil {
		return nil, err
	}
	return &place, nil
}

func (repo RoomRepo) CreateBuilding(building models_Room.Building) (*models_Room.Building, error) {
	err := db.DB.Create(&building).Error
	if err != nil {
		return nil, err
	}
	return &building, nil
}

func (repo RoomRepo) CreateFloor(floor models_Room.Floor) (*models_Room.Floor, error) {
	err := db.DB.Create(&floor).Error
	if err != nil {
		return nil, err
	}
	return &floor, nil
}

func (repo RoomRepo) CreateRoom(room models_Room.Room) (*models_Room.Room, error) {
	err := db.DB.Create(&room).Error
	if err != nil {
		return nil, err
	}
	return &room, nil
}

func (repo RoomRepo) ModifyRoom(room models_Room.Room) (*models_Room.Room, error) {
	err := db.DB.Save(&room).Error
	if err != nil {
		return nil, err
	}
	return &room, nil
}

func (repo RoomRepo) GetRoomByID(roomID uint) (*models_Room.Room, error) {
	var room models_Room.Room
	err := db.DB.First(&room, roomID).Error
	if err != nil {
		fmt.Println("Error finding room:", err)
		return nil, err
	}
	return &room, nil
}

func (repo RoomRepo) GetAllRoomPrices() ([]models_Room.RoomPrice, error) {
	var prices []models_Room.RoomPrice
	err := db.DB.Find(&prices).Error
	if err != nil {
		fmt.Println("Error finding records:", err)
		return nil, err
	}
	return prices, nil
}

func (repo RoomRepo) GetAllPersons() ([]models_Person.Person, error) {
	var models []models_Person.Person
	err := db.DB.Find(&models).Error
	if err != nil {
		fmt.Println("Error finding records:", err)
		return nil, err
	}
	return models, nil
}

func (repo RoomRepo) GetAllContacts() ([]models_Person.Contact, error) {
	var models []models_Person.Contact
	err := db.DB.Find(&models).Error
	if err != nil {
		fmt.Println("Error finding records:", err)
		return nil, err
	}
	return models, nil
}
