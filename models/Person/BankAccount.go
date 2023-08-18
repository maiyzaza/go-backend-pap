package models

import (
	"time"
)

type BankAccount struct {
	ID            uint       `json:"ID" gorm:"primaryKey"`
	PersonID      uint       `json:"PersonID" gorm:"person_id"`
	BankName      string     `json:"BankName" gorm:"bank_name"`
	BankAddress   string     `json:"BankAddress" gorm:"bank_address"`
	AccountName   string     `json:"AccountName" gorm:"account_name"`
	AccountNumber string     `json:"AccountNumber" gorm:"account_number"`
	SwiftCode     *string    `json:"SwiftCode" gorm:"swift_code"`
	IsActive      bool       `json:"IsActive" gorm:"is_active"`
	CreatedAt     time.Time  `json:"CreatedAt" gorm:"created_at"`
	UpdatedAt     *time.Time `json:"UpdatedAt" gorm:"updated_at"`

	Person Person `json:"Person" gorm:"foreignKey:person_id"`
}
