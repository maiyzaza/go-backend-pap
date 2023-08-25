package models

import (
	"time"
)

type RoomPrice struct {
	ID              uint       `json:"ID" gorm:"primaryKey"`
	RoomID          uint       `json:"RoomID" gorm:"room_id"`
	Amount          float32    `json:"Amount" gorm:"amount"`
	UnitType        *string    `json:"UnitType" gorm:"unit_type"`
	MinDuration     *int32     `json:"MinDuration" gorm:"min_duration"`
	MaxDuration     *int32     `json:"MaxDuration" gorm:"max_duration"`
	Type            string     `json:"Type" gorm:"type:varchar(255)"`
	DepositAmount   *float32   `json:"DepositAmount" gorm:"deposit_amount"`
	DepositUnitType *string    `json:"DepositUnitType" gorm:"deposit_unit_type"`
	IsActive        bool       `json:"IsActive" gorm:"is_active"`
	CreatedAt       time.Time  `json:"CreatedAt" gorm:"created_at"`
	UpdatedAt       *time.Time `json:"UpdatedAt" gorm:"updated_at"`

	Room Room `json:"Room" gorm:"foreignKey:RoomID"`
}
