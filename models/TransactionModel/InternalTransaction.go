package models

import (
	"time"
)

type InternalTransaction struct {
	ID          uint      `gorm:"primaryKey"`
	Product     string    `gorm:"product"`
	Description string    `gorm:"description"`
	Quantity    int       `gorm:"quantity"`
	Amount      float32   `gorm:"amount"`
	Total       float32   `gorm:"total"`
	Branch      string    `gorm:"branch"`     //Branch1, ...
	IsReceive   bool      `gorm:"is_receive"` // didnot update database yet
	IsActive    bool      `gorm:"is_active"`
	CreatedAt   time.Time `gorm:"created_at"`
	UpdatedAt   time.Time `gorm:"updated_at"`
}
