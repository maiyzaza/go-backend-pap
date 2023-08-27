package models

import (
	"time"
)

type InternalTransaction struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Product     string    `json:"product" gorm:"product"`
	Description string    `json:"description" gorm:"description"`
	Quantity    int       `json:"quantity" gorm:"quantity"`
	Amount      float32   `json:"amount" gorm:"amount"`
	Total       float32   `json:"total" gorm:"total"`
	Branch      string    `json:"branch" gorm:"branch"`         //Branch1, ...
	IsReceive   bool      `json:"is_receive" gorm:"is_receive"` // did not update database yet
	IsActive    bool      `json:"is_active" gorm:"is_active"`
	CreatedAt   time.Time `json:"created_at" gorm:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"updated_at"`
}
