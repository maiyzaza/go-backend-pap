package models

import (
	"time"

	models_Person "PattayaAvenueProperty/models/Person"
)

type Room struct {
	ID                 uint       `gorm:"primaryKey"`
	OwnerID            *uint      `gorm:"owner_id"`
	FloorID            uint       `gorm:"floor_id"`
	RoomName           *string    `gorm:"room_name"`
	RoomNumber         string     `gorm:"room_number"`
	RoomAddress        string     `gorm:"room_address"`
	ElectricNumber     *string    `gorm:"electric_number"`
	ElectricUserNumber *string    `gorm:"electric_user_number"`
	AmountOfBedRoom    *int32     `gorm:"amount_of_bed_room"`
	AmountOfToiletRoom *int32     `gorm:"amount_of_toilet_room"`
	AmountOfLivingRoom *int32     `gorm:"amount_of_living_room"`
	SizeSQM            float32    `gorm:"size_sqm"`
	TypeOfView         string     `gorm:"type_of_view"` // sea, city, city and sea
	Remark             *string    `gorm:"remark"`
	StatusOfRoom       string     `gorm:"status_of_room"` // Rent, Sale, Rent and Sale, Returned
	IsActive           bool       `gorm:"is_active"`
	CreatedAt          time.Time  `gorm:"created_at"`
	UpdatedAt          *time.Time `gorm:"updated_at"`

	Person models_Person.Person `gorm:"foreignKey:OwnerID"`
	Floor  Floor                `gorm:"foreignKey:FloorID"`
}
