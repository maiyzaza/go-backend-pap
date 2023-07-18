package models

import (
	"time"

	models_Contract "PattayaAvenueProperty/models/Contract"
	models_Document "PattayaAvenueProperty/models/Document"
)

type Transaction struct {
	ID             uint       `gorm:"primaryKey"`
	RoomContractID *uint      `gorm:"room_contract_id"`
	DocumentId     uint       `gorm:"document_id"`
	CategoryType   string     `gorm:"category_type"` // RENTAL, SELL, ELECTRIC, DEPT, ...
	IsReceive      bool       `gorm:"is_receive"`    // didnot update database yet
	Description    string     `gorm:"description"`
	PaymentMethod  string     `gorm:"payment_method"` // CREDIT, CASH, ...
	Amount         float32    `gorm:"amount"`
	Remark         string     `gorm:"remark"`
	Branch         string     `gorm:"branch"` //Branch1, ...
	IsActive       bool       `gorm:"is_active"`
	CreatedAt      time.Time  `gorm:"created_at"`
	UpdatedAt      *time.Time `gorm:"updated_at"`

	RoomContract models_Contract.RoomContract `gorm:"foreignKey:RoomContractID"`
	Document     models_Document.Document     `gorm:"fo reignKey:DocumentId"`
}
