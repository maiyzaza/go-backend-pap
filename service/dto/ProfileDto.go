package dto

type PersonDto struct {
	ID                  uint             `json:"id" gorm:"primaryKey"`
	FullName            string           `json:"full_name" gorm:"full_name"`
	CitizenDocumentUrl  *string          `json:"citizen_document_url" gorm:"citizen_doucument_url"`
	PassportDocumentUrl *string          `json:"passport_document_url" gorm:"passport_document_url"`
	BankAccounts        []BankAccountDto `json:"bank_accounts" gorm:"bank_accounts"`
}

type BankAccountDto struct {
	ID            uint    `json:"id" gorm:"primaryKey"`
	PersonId      uint    `json:"person_id" gorm:"person_id"`
	BankName      string  `json:"bank_name" gorm:"bank_name"`
	BankAddress   string  `json:"bank_address" gorm:"bank_address"`
	AccountName   string  `json:"account_name" gorm:"account_name"`
	AccountNumber string  `json:"account_number" gorm:"account_number"`
	SwiftCode     *string `json:"swift_code" gorm:"swift_code"`
}
