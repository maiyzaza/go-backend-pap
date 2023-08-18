package models

import (
	"time"

	models_Room "PattayaAvenueProperty/models/Room"
)

type RoomContract struct {
	ID                     uint       `json:"ID" gorm:"primaryKey"`
	RoomID                 uint       `json:"RoomID" gorm:"room_id"`
	ContractName           string     `json:"ContractName" gorm:"contract_name"`
	StartContractDate      time.Time  `json:"StartContractDate" gorm:"start_contract_date"`
	EndContractDate        time.Time  `json:"EndContractDate" gorm:"end_contract_date"`
	Rental                 float32    `json:"Rental" gorm:"rental"`
	Deposit                float32    `json:"Deposit" gorm:"deposit"`
	CheckInElectricNumber  *int       `json:"CheckInElectricNumber" gorm:"check_in_electric_number"`
	CheckInWaterNumber     *int       `json:"CheckInWaterNumber" gorm:"check_in_water_number"`
	CheckOutElectricNumber *int       `json:"CheckOutElectricNumber" gorm:"check_out_electric_number"`
	CheckOutWaterNumber    *int       `json:"CheckOutWaterNumber" gorm:"check_out_water_number"`
	CheckInDate            *time.Time `json:"CheckInDate" gorm:"check_in_date"`
	CheckOutDate           *time.Time `json:"CheckOutDate" gorm:"check_out_date"`
	IsClosed               bool       `json:"IsClosed" gorm:"is_closed"`
	IsActive               bool       `json:"IsActive" gorm:"is_active"`
	CreatedAt              time.Time  `json:"CreatedAt" gorm:"created_at"`
	UpdatedAt              *time.Time `json:"UpdatedAt" gorm:"updated_at"`

	Room models_Room.Room `json:"Room" gorm:"foreignKey:RoomID"`
}
