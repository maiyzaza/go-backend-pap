package dto

// type PersonDto struct {
// 	ID                  uint             `json:"id" gorm:"primaryKey"`
// 	FullName            string           `json:"full_name" gorm:"full_name"`
// 	IdentityNumber      string           `json:"identity_number" gorm:"identity_number"`
// 	CitizenDocumentUrl  *string          `json:"citizen_document_url" gorm:"citizen_doucument_url"`
// 	PassportDocumentUrl *string          `json:"passport_document_url" gorm:"passport_document_url"`
// 	BankAccounts        []BankAccountDto `json:"bank_accounts" gorm:"bank_accounts"`
// }

type PersonDto struct {
	ID             uint             `json:"id" gorm:"primaryKey"`
	FullName       string           `json:"full_name" gorm:"full_name"`
	IdentityNumber string           `json:"identity_number" gorm:"identity_number"`
	BankAccounts   []BankAccountDto `json:"bank_accounts" gorm:"bank_accounts"`
}

type PersonDetailDto struct {
	ID             uint             `json:"id" gorm:"primaryKey"`
	FullName       string           `json:"full_name" gorm:"full_name"`
	IdentityNumber string           `json:"identity_number" gorm:"identity_number"`
	Contacts       []ContactDto     `json:"contacts" gorm:"contacts"`
	BankAccounts   []BankAccountDto `json:"bank_accounts" gorm:"bank_accounts"`
}

type ContactDto struct {
	ID    uint   `json:"id" gorm:"primaryKey"`
	Type  string `json:"type" gorm:"type"`
	Value string `json:"value" gorm:"value"`
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

type CreatePersonDto struct {
	FullName       string `json:"full_name" gorm:"full_name"`
	IdentityNumber string `json:"identity_number" gorm:"identity_number"`
	TypeContact    string `json:"type_contact" gorm:"type_contact"`
	ValueContact   string `json:"value_contact" gorm:"value_contact"`
}

type EditPeopleDto struct {
	PersonID       uint   `json:"person_id" gorm:"person_id"`
	FullName       string `json:"full_name" gorm:"full_name"`
	IdentityNumber string `json:"identity_number" gorm:"identity_number"`
}

type CreateContactDto struct {
	PersonID     uint   `json:"person_id" gorm:"person_id"`
	TypeContact  string `json:"type_contact" gorm:"type_contact"`
	ValueContact string `json:"value_contact" gorm:"value_contact"`
}

type EditContactDto struct {
	ID           uint   `json:"id" gorm:"primaryKey"`
	TypeContact  string `json:"type_contact" gorm:"type_contact"`
	ValueContact string `json:"value_contact" gorm:"value_contact"`
}
