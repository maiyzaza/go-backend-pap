package models

import (
	"time"

	models_Document "PattayaAvenueProperty/models/Document"
)

type InternalTransactionDocument struct {
	ID                    uint       `gorm:"primaryKey" json:"id"`
	DocumentID            uint       `gorm:"document_id" json:"document_id"`
	InternalTransactionID uint       `gorm:"internal_transaction_id" json:"internal_transaction_id"`
	IsActive              bool       `gorm:"is_active" json:"is_active"`
	CreatedAt             time.Time  `gorm:"created_at" json:"created_at"`
	UpdatedAt             *time.Time `gorm:"updated_at" json:"updated_at"`

	Document            models_Document.Document `json:"document" gorm:"foreignKey:DocumentID"`
	Internaltransaction InternalTransaction      `json:"internal_transaction" gorm:"foreignKey:InternalTransactionID"`
}
