package models

import (
	models "PattayaAvenueProperty/models/Person"
	"time"
)

type Room struct {
	ID                 uint       `json:"ID" gorm:"primaryKey"`
	OwnerID            *uint      `json:"OwnerID" gorm:"owner_id"`
	FloorID            uint       `json:"FloorID" gorm:"floor_id"`
	RoomName           *string    `json:"RoomName" gorm:"room_name"`
	RoomNumber         string     `json:"RoomNumber" gorm:"room_number"`
	RoomAddress        string     `json:"RoomAddress" gorm:"room_address"`
	ElectricNumber     *string    `json:"ElectricNumber" gorm:"electric_number"`
	ElectricUserNumber *string    `json:"ElectricUserNumber" gorm:"electric_user_number"`
	AmountOfBedRoom    *int32     `json:"AmountOfBedRoom" gorm:"amount_of_bed_room"`
	AmountOfToiletRoom *int32     `json:"AmountOfToiletRoom" gorm:"amount_of_toilet_room"`
	AmountOfLivingRoom *int32     `json:"AmountOfLivingRoom" gorm:"amount_of_living_room"`
	SizeSQM            float32    `json:"SizeSQM" gorm:"size_sqm"`
	TypeOfView         string     `json:"TypeOfView" gorm:"type_of_view"`
	Remark             *string    `json:"Remark" gorm:"remark"`
	StatusOfRoom       string     `json:"StatusOfRoom" gorm:"status_of_room"`
	IsActive           bool       `json:"IsActive" gorm:"is_active"`
	CreatedAt          time.Time  `json:"CreatedAt" gorm:"created_at"`
	UpdatedAt          *time.Time `json:"UpdatedAt" gorm:"updated_at"`

	Person models.Person `json:"Person" gorm:"foreignKey:OwnerID"`
	Floor  Floor         `json:"Floor" gorm:"foreignKey:FloorID"`
}
