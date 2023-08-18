package models

import (
	"time"

	models_Contract "PattayaAvenueProperty/models/Contract"
	models_Document "PattayaAvenueProperty/models/Document"
)

type Transaction struct {
	ID             uint       `json:"id" gorm:"primaryKey"`
	RoomContractID *uint      `json:"room_contract_id" gorm:"room_contract_id"`
	DocumentId     uint       `json:"document_id" gorm:"document_id"`
	CategoryType   string     `json:"category_type" gorm:"category_type"` // RENTAL, SELL, ELECTRIC, DEPT, ...
	IsReceive      bool       `json:"is_receive" gorm:"is_receive"`       // did not update database yet
	Description    string     `json:"description" gorm:"description"`
	PaymentMethod  string     `json:"payment_method" gorm:"payment_method"` // CREDIT, CASH, ...
	Amount         float32    `json:"amount" gorm:"amount"`
	Remark         string     `json:"remark" gorm:"remark"`
	Branch         string     `json:"branch" gorm:"branch"` // Branch1, ...
	IsActive       bool       `json:"is_active" gorm:"is_active"`
	CreatedAt      time.Time  `json:"created_at" gorm:"created_at"`
	UpdatedAt      *time.Time `json:"updated_at" gorm:"updated_at"`

	RoomContract models_Contract.RoomContract `json:"room_contract" gorm:"foreignKey:RoomContractID"`
	Document     models_Document.Document     `json:"document" gorm:"foreignKey:DocumentId"`
}
