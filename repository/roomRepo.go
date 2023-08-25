package repository

import (
	"PattayaAvenueProperty/db"
	models_Document "PattayaAvenueProperty/models/Document"
	models_Person "PattayaAvenueProperty/models/Person"
	models_Room "PattayaAvenueProperty/models/Room"
	"fmt"

	"gorm.io/gorm"
)

type RoomRepo struct{}

func NewRoomRepo() RoomRepo {
	return RoomRepo{}
}

func ActiveOnlyRoom(db *gorm.DB) *gorm.DB {
	return db.Where("is_active = ?", 1)
}

// Get all places, buildings, floors, rooms
func (repo RoomRepo) GetAllPlace() ([]models_Room.Place, error) {
	var model []models_Room.Place
	err := ActiveOnlyRoom(db.DB).Find(&model).Error
	if err != nil {
		fmt.Println("Error finding active records:", err)
		return nil, err
	}
	return model, nil
}

func (repo RoomRepo) GetAllBuilding() ([]models_Room.Building, error) {
	var model []models_Room.Building
	err := ActiveOnlyRoom(db.DB).Find(&model).Error
	if err != nil {
		fmt.Println("Error finding records:", err)
		return nil, err
	}
	return model, nil
}

func (repo RoomRepo) GetAllFloor() ([]models_Room.Floor, error) {
	var model []models_Room.Floor
	err := ActiveOnlyRoom(db.DB).Find(&model).Error
	if err != nil {
		fmt.Println("Error finding records:", err)
		return nil, err
	}
	return model, nil
}

func (repo RoomRepo) GetAllRoom() ([]models_Room.Room, error) {
	var model []models_Room.Room
	err := ActiveOnlyRoom(db.DB).Find(&model).Error
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
	var models models_Room.Room
	err := ActiveOnlyRoom(db.DB).First(&models, roomID).Error
	if err != nil {
		fmt.Println("Error finding room:", err)
		return nil, err
	}
	return &models, nil
}

func (repo RoomRepo) GetAllRoomPrices() ([]models_Room.RoomPrice, error) {
	var models []models_Room.RoomPrice
	err := ActiveOnlyRoom(db.DB).Find(&models).Error
	if err != nil {
		fmt.Println("Error finding records:", err)
		return nil, err
	}
	return models, nil
}

func (repo RoomRepo) GetAllRoomPictures() ([]models_Room.RoomPicture, error) {
	var models []models_Room.RoomPicture
	err := ActiveOnlyRoom(db.DB).Find(&models).Error
	if err != nil {
		fmt.Println("Error finding records:", err)
		return nil, err
	}
	return models, nil
}

func (repo RoomRepo) GetAllPersons() ([]models_Person.Person, error) {
	var models []models_Person.Person
	err := ActiveOnlyRoom(db.DB).Find(&models).Error
	if err != nil {
		fmt.Println("Error finding records:", err)
		return nil, err
	}
	return models, nil
}

func (repo RoomRepo) GetAllContacts() ([]models_Person.Contact, error) {
	var models []models_Person.Contact
	err := ActiveOnlyRoom(db.DB).Find(&models).Error
	if err != nil {
		fmt.Println("Error finding records:", err)
		return nil, err
	}
	return models, nil
}

func (repo RoomRepo) GetAllRoomDocuments() ([]models_Document.RoomDocument, error) {
	var models []models_Document.RoomDocument
	err := ActiveOnlyRoom(db.DB).Find(&models).Error
	if err != nil {
		fmt.Println("Error finding records:", err)
		return nil, err
	}
	return models, nil
}

func (repo RoomRepo) GetRoomDocumentByRoomID(roomID uint) ([]models_Document.RoomDocument, error) {
	var models []models_Document.RoomDocument
	err := ActiveOnlyRoom(db.DB).Where("room_id = ?", roomID).Find(&models).Error
	if err != nil {
		fmt.Println("Error finding room document:", err)
		return nil, err
	}
	return models, nil
}

func (repo RoomRepo) GetRoomPriceByRoomID(roomID uint) ([]models_Room.RoomPrice, error) {
	var models []models_Room.RoomPrice
	err := ActiveOnlyRoom(db.DB).Where("room_id = ?", roomID).Find(&models).Error
	if err != nil {
		fmt.Println("Error finding room price:", err)
		return nil, err
	}
	return models, nil
}

func (repo RoomRepo) GetRoomPictureByRoomID(roomID uint) ([]models_Room.RoomPicture, error) {
	var models []models_Room.RoomPicture
	err := ActiveOnlyRoom(db.DB).Where("room_id = ?", roomID).Find(&models).Error
	if err != nil {
		fmt.Println("Error finding room picture:", err)
		return nil, err
	}
	return models, nil
}

func (repo RoomRepo) EditPlace(place models_Room.Place) (*models_Room.Place, error) {
	err := db.DB.Save(&place).Error
	if err != nil {
		return nil, err
	}
	return &place, nil
}

func (repo RoomRepo) GetPlaceByID(placeID uint) (*models_Room.Place, error) {
	var model models_Room.Place
	err := ActiveOnlyRoom(db.DB).First(&model, placeID).Error
	if err != nil {
		fmt.Println("Error finding place:", err)
		return nil, err
	}
	return &model, nil
}

// edit building and get building by id
func (repo RoomRepo) EditBuilding(building models_Room.Building) (*models_Room.Building, error) {
	err := db.DB.Save(&building).Error
	if err != nil {
		return nil, err
	}
	return &building, nil
}

func (repo RoomRepo) GetBuildingByID(buildingID uint) (*models_Room.Building, error) {
	var model models_Room.Building
	err := ActiveOnlyRoom(db.DB).First(&model, buildingID).Error
	if err != nil {
		fmt.Println("Error finding building:", err)
		return nil, err
	}
	return &model, nil
}

// create RoomPrice and delete RoomPrice by change is_active to 0
func (repo RoomRepo) CreateRoomPrice(roomPrice models_Room.RoomPrice) (*models_Room.RoomPrice, error) {
	err := db.DB.Create(&roomPrice).Error
	if err != nil {
		return nil, err
	}
	return &roomPrice, nil
}

func (repo RoomRepo) DeleteRoomPrice(roomPriceID uint) error {
	err := db.DB.Model(&models_Room.RoomPrice{}).Where("id = ?", roomPriceID).Update("is_active", 0).Error
	if err != nil {
		return err
	}
	return nil
}

// create RoomPicture and delete RoomPicture by change is_active to 0
func (repo RoomRepo) CreateRoomPicture(roomPicture models_Room.RoomPicture) (*models_Room.RoomPicture, error) {
	err := db.DB.Create(&roomPicture).Error
	if err != nil {
		return nil, err
	}
	return &roomPicture, nil
}

func (repo RoomRepo) DeleteRoomPicture(roomPictureID uint) error {
	err := db.DB.Model(&models_Room.RoomPicture{}).Where("id = ?", roomPictureID).Update("is_active", 0).Error
	if err != nil {
		return err
	}
	return nil
}

// create RoomDocument and delete RoomDocument by change is_active to 0
func (repo RoomRepo) CreateRoomDocument(roomDocumentPicture models_Document.RoomDocument) (*models_Document.RoomDocument, error) {
	err := db.DB.Create(&roomDocumentPicture).Error
	if err != nil {
		return nil, err
	}
	return &roomDocumentPicture, nil
}

func (repo RoomRepo) DeleteRoomDocument(roomDocumentPictureID uint) error {
	err := db.DB.Model(&models_Document.RoomDocument{}).Where("id = ?", roomDocumentPictureID).Update("is_active", 0).Error
	if err != nil {
		return err
	}
	return nil
}
