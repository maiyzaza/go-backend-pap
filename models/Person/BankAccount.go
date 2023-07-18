package models

import (
	"time"
)

type BankAccount struct {
	ID            uint       `gorm:"primaryKey"`
	PersonId      uint       `gorm:"person_id"`
	BankName      string     `gorm:"bank_name"`
	BankAddress   string     `gorm:"bank_address"`
	AccountName   string     `gorm:"account_name"`
	AccountNumber string     `gorm:"account_number"`
	SwiftCode     *string    `gorm:"swift_code"`
	IsActive      bool       `gorm:"is_active"`
	CreatedAt     time.Time  `gorm:"created_at"`
	UpdatedAt     *time.Time `gorm:"updated_at"`

	Person Person `gorm:"foreignKey:person_id"`
}
