package models

import (
	"time"
)

type Person struct {
	ID                  uint          `json:"ID" gorm:"primaryKey"`
	FullName            string        `json:"FullName" gorm:"full_name"`
	CitizenDocumentUrl  *string       `json:"CitizenDocumentUrl" gorm:"citizen_doucument_url"`
	PassportDocumentUrl *string       `json:"PassportDocumentUrl" gorm:"passport_document_url"`
	IsActive            bool          `json:"IsActive" gorm:"is_active"`
	CreatedAt           time.Time     `json:"CreatedAt" gorm:"type:timestamp"`
	UpdatedAt           *time.Time    `json:"UpdatedAt" gorm:"updated_at"`
	BankAccounts        []BankAccount `json:"BankAccounts" gorm:"bank_accounts;foreignKey:person_id"`
}

func (Person) TableName() string {
	return "persons"
}
