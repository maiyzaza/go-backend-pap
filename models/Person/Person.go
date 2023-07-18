package models

import (
	"time"
)

type Person struct {
	ID                  uint          `gorm:"primaryKey"`
	FullName            string        `gorm:"full_name"`
	CitizenDocumentUrl  *string       `gorm:"citizen_doucument_url"`
	PassportDocumentUrl *string       `gorm:"passport_document_url"`
	IsActive            bool          `gorm:"is_active"`
	CreatedAt           time.Time     `gorm:"type:timestamp"`
	UpdatedAt           *time.Time    `gorm:"updated_at"`
	BankAccounts        []BankAccount `gorm:"bank_accounts;foreignKey:person_id"`
}

func (Person) TableName() string {
	return "persons"
}
