package dto

type TransactionResponseDto struct {
	ID            uint    `json:"id" gorm:"primaryKey"`
	CatorgoryType string  `json:"catorgory_type" gorm:"catorgory_type"`
	RoomAddress   string  `json:"room_address" gorm:"room_address"`
	PaymentMethod string  `json:"payment_method" gorm:"payment_method"`
	Amount        float32 `json:"amount" gorm:"amount"`
	IsReceive     bool    `json:"is_receive" gorm:"is_receive"`
	CreateAt      string  `json:"create_at" gorm:"create_at"`
}

type TransactionDetailResponseDto struct {
	ID                  uint                           `json:"id" gorm:"primaryKey"`
	CatorgoryType       string                         `json:"catorgory_type" gorm:"catorgory_type"`
	Amount              float32                        `json:"amount" gorm:"amount"`
	PaymentMethod       string                         `json:"payment_method" gorm:"payment_method"`
	RoomAddress         string                         `json:"room_address" gorm:"room_address"`
	IsReceive           string                         `json:"is_receive" gorm:"is_receive"`
	Description         string                         `json:"description" gorm:"description"`
	Remark              string                         `json:"remark" gorm:"remark"`
	TransactionDocument TransactionDocumentResponseDto `json:"transaction_document" gorm:"transaction_document"`
	CreateAt            string                         `json:"create_at" gorm:"create_at"`
}

type TransactionDocumentResponseDto struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	DocumentUrl string `json:"document_url" gorm:"document_url"`
}

type CreateTransactionDto struct {
	RoomID        uint    `json:"room_id"`
	CategoryType  string  `json:"category_type"`
	IsReceive     bool    `json:"is_receive"`
	Description   string  `json:"description"`
	PaymentMethod string  `json:"payment_method"`
	Amount        float32 `json:"amount"`
	Remark        string  `json:"remark"`
	DocumentUrl   string  `json:"document_url"`
}

// create Document
type CreateDocumentDto struct {
	DocumentUrl string `json:"document_url"`
}
