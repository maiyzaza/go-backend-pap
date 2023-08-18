package models

import (
	"time"

	models_Person "PattayaAvenueProperty/models/Person"
)

type PersonContract struct {
	ID             uint       `json:"ID" gorm:"primaryKey"`
	PersonID       uint       `json:"PersonID" gorm:"person_id"`
	RoomContractID uint       `json:"RoomContractID" gorm:"room_contract_id"`
	Type           string     `json:"Type" gorm:"type:varchar(255)"` // buyer, seller, owner, salesperson
	IsActive       bool       `json:"IsActive" gorm:"is_active"`
	CreatedAt      time.Time  `json:"CreatedAt" gorm:"created_at"`
	UpdatedAt      *time.Time `json:"UpdatedAt" gorm:"updated_at"`

	Person       models_Person.Person `json:"Person" gorm:"foreignKey:PersonID"`
	RoomContract RoomContract         `json:"RoomContract" gorm:"foreignKey:RoomContractID"`
}
