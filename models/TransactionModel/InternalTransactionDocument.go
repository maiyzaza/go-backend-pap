package models

import (
	"time"

	models_Document "PattayaAvenueProperty/models/Document"
)

type InternalTransactionDocument struct {
	ID                    uint       `gorm:"primaryKey"`
	DocumentID            uint       `gorm:"document_id"`
	InternalTransactionID uint       `gorm:"internal_transaction_id"`
	IsActive              bool       `gorm:"is_active"`
	CreatedAt             time.Time  `gorm:"created_at"`
	UpdatedAt             *time.Time `gorm:"updated_at"`

	Document            models_Document.Document `gorm:"foreignKey:DocumentID"`
	Internaltransaction InternalTransaction      `gorm:"foreignKey:InternalTransactionID"`
}
