package models

import (
	"time"
)

type RoomPrice struct {
	ID              int        `gorm:"primaryKey"`
	RoomID          int        `gorm:"room_id"`
	Amount          float32    `gorm:"amount"`
	UnitType        *string    `gorm:"unit_type"` // day, month, year
	MinDuration     *int32     `gorm:"min_duration"`
	MaxDuration     *int32     `gorm:"max_duration"`
	Type            string     `gorm:"type:varchar(255)"` // rent, sell
	DepositAmount   *float32   `gorm:"deposit_amount"`
	DepositUnitType *string    `gorm:"deposit_unit_type"` // baht, month
	IsActive        bool       `gorm:"is_active"`
	CreatedAt       time.Time  `gorm:"created_at"`
	UpdatedAt       *time.Time `gorm:"updated_at"`

	Room Room `gorm:"foreignKey:RoomID"`
}
