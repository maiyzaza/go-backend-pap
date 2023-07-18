package models

import (
	"time"

	models_Person "PattayaAvenueProperty/models/Person"
)

type PersonContract struct {
	ID             uint       `gorm:"primaryKey"`
	PersonID       uint       `gorm:"person_id"`
	RoomContractID uint       `gorm:"room_contract_id"`
	Type           string     `gorm:"type:varchar(255)"` // buyer, seller, owner, salesperson
	IsActive       bool       `gorm:"is_active"`
	CreatedAt      time.Time  `gorm:"created_at"`
	UpdatedAt      *time.Time `gorm:"updated_at"`

	Person       models_Person.Person `gorm:"foreignKey:PersonID"`
	RoomContract RoomContract         `gorm:"foreignKey:RoomContractID"`
}
