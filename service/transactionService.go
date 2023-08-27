package service

import (
	models_Document "PattayaAvenueProperty/models/Document"
	models_Transaction "PattayaAvenueProperty/models/Transaction"
	"PattayaAvenueProperty/repository"
	dto "PattayaAvenueProperty/service/dto"
)

type TransactionService struct {
	transactionRepo repository.TransactionRepo
	roomRepo        repository.RoomRepo
	contractRepo    repository.ContractRepo
}

func NewTransactionService(transactionRepo repository.TransactionRepo, roomRepo repository.RoomRepo, contractRepo repository.ContractRepo) TransactionService {
	return TransactionService{
		transactionRepo: transactionRepo,
		roomRepo:        roomRepo,
		contractRepo:    contractRepo,
	}
}

// Get all transactions
func (service *TransactionService) GetAllTransaction() ([]dto.TransactionResponseDto, error) {
	transactions, err := service.transactionRepo.GetAllTransaction()
	if err != nil {
		return nil, err
	}
	// get room address from room table, room id = room id from transaction
	rooms, err := service.roomRepo.GetAllRoom()
	if err != nil {
		return nil, err
	}

	var result []dto.TransactionResponseDto
	for _, transaction := range transactions {
		var roomAddress string
		for _, room := range rooms {
			if transaction.RoomID != nil && *transaction.RoomID == room.ID {
				roomAddress = room.RoomAddress
			}
		}
		result = append(result, dto.TransactionResponseDto{
			ID:            transaction.ID,
			CatorgoryType: transaction.CategoryType,
			RoomAddress:   roomAddress,
			PaymentMethod: transaction.PaymentMethod,
			Amount:        transaction.Amount,
			IsReceive:     transaction.IsReceive,
		})
	}
	return result, nil
}

// get transaction by id
func (service *TransactionService) GetTransactionByID(id uint) (*dto.TransactionDetailResponseDto, error) {
	transaction, err := service.transactionRepo.GetTransactionByID(id)
	if err != nil {
		return nil, err
	}
	// get room id from transaction and match with room id from room and get room address
	rooms, err := service.roomRepo.GetAllRoom()
	if err != nil {
		return nil, err
	}

	var roomAddress string
	for _, room := range rooms {
		if transaction.RoomID != nil && *transaction.RoomID == room.ID {
			roomAddress = room.RoomAddress
		}
	}

	transactionDocuments, err := service.transactionRepo.GetAllDocument()
	if err != nil {
		return nil, err
	}

	// one transaction can have one document
	var transactionDocumentList dto.TransactionDocumentResponseDto
	for _, transactionDocument := range transactionDocuments {
		if transaction.DocumentId == transactionDocument.ID {
			transactionDocumentList = dto.TransactionDocumentResponseDto{
				ID:          transactionDocument.ID,
				DocumentUrl: transactionDocument.DocumentUrl,
			}
		}
	}

	// change is receive to "recevied" or "not receive"
	var isReceive string
	if transaction.IsReceive {
		isReceive = "received"
	} else {
		isReceive = "not receive"
	}

	return &dto.TransactionDetailResponseDto{
		ID:                  transaction.ID,
		CatorgoryType:       transaction.CategoryType,
		Amount:              transaction.Amount,
		PaymentMethod:       transaction.PaymentMethod,
		RoomAddress:         roomAddress,
		IsReceive:           isReceive,
		Description:         transaction.Description,
		Remark:              transaction.Remark,
		TransactionDocument: transactionDocumentList,
	}, nil
}

// create transaction
func (service *TransactionService) CreateTransaction(transactionDto dto.CreateTransactionDto, documentID uint) (*dto.TransactionResponseDto, error) {

	transaction := models_Transaction.Transaction{
		RoomID:        &transactionDto.RoomID,
		DocumentId:    documentID,
		CategoryType:  transactionDto.CategoryType,
		IsReceive:     transactionDto.IsReceive,
		Description:   transactionDto.Description,
		PaymentMethod: transactionDto.PaymentMethod,
		Amount:        transactionDto.Amount,
		Remark:        transactionDto.Remark,
	}

	transactionModel, err := service.transactionRepo.CreateTransaction(&transaction)
	if err != nil {
		return nil, err
	}
	return &dto.TransactionResponseDto{
		ID:            transactionModel.ID,
		CatorgoryType: transactionModel.CategoryType,
		// RoomAddress:   transactionDto.RoomAddress,
		PaymentMethod: transactionModel.PaymentMethod,
		Amount:        transactionModel.Amount,
		IsReceive:     transactionModel.IsReceive,
	}, nil
}

// create document
func (service *TransactionService) CreateDocument(documentDto dto.CreateDocumentDto) (*dto.TransactionDocumentResponseDto, error) {
	document := models_Document.Document{
		DocumentUrl: documentDto.DocumentUrl,
	}

	documentModel, err := service.transactionRepo.CreateDocument(&document)
	if err != nil {
		return nil, err
	}
	return &dto.TransactionDocumentResponseDto{
		ID:          documentModel.ID,
		DocumentUrl: documentModel.DocumentUrl,
	}, nil
}

// delete transaction by setting is_active to false
func (service *TransactionService) DeleteTransaction(id uint) error {
	err := service.transactionRepo.DeleteTransaction(id)
	if err != nil {
		return err
	}
	return nil
}

// get all trasanction that is deleted
func (service *TransactionService) GetAllDeletedTransaction() ([]dto.TransactionResponseDto, error) {
	transactions, err := service.transactionRepo.GetAllDeletedTransaction()
	if err != nil {
		return nil, err
	}
	// get room address from room table, room id = room id from transaction
	rooms, err := service.roomRepo.GetAllRoom()
	if err != nil {
		return nil, err
	}

	var result []dto.TransactionResponseDto
	for _, transaction := range transactions {
		var roomAddress string
		for _, room := range rooms {
			if transaction.RoomID != nil && *transaction.RoomID == room.ID {
				roomAddress = room.RoomAddress
			}
		}
		result = append(result, dto.TransactionResponseDto{
			ID:            transaction.ID,
			CatorgoryType: transaction.CategoryType,
			RoomAddress:   roomAddress,
			PaymentMethod: transaction.PaymentMethod,
			Amount:        transaction.Amount,
			IsReceive:     transaction.IsReceive,
		})
	}
	return result, nil
}

// func (service *PersonService) GetProfiles() ([]dto.PersonDto, error) {
// 	persons, err := service.personRepo.FindAll()
// 	if err != nil {
// 		return nil, err
// 	}
// 	var result []dto.PersonDto
// 	for _, person := range persons {
// 		result = append(result, dto.PersonDto{
// 			ID:             person.ID,
// 			FullName:       person.FullName,
// 			IdentityNumber: person.IdentityNumber,
// 		})
// 	}
// 	return result, nil
// }

// func (service *PersonService) GetProfilesWithBankAccount() ([]dto.PersonDetailDto, error) {
// 	persons, err := service.personRepo.FindAll()
// 	if err != nil {
// 		return nil, err
// 	}
// 	contacts, err := service.personRepo.GetAllContact()
// 	if err != nil {
// 		return nil, err
// 	}

// 	var result []dto.PersonDetailDto
// 	for _, person := range persons {
// 		var bank []dto.BankAccountDto
// 		for _, bankAccount := range person.BankAccounts {
// 			bank = append(bank, dto.BankAccountDto{
// 				ID:            bankAccount.ID,
// 				PersonId:      bankAccount.PersonID,
// 				AccountNumber: bankAccount.AccountNumber,
// 				BankAddress:   bankAccount.BankAddress,
// 				BankName:      bankAccount.BankName,
// 				AccountName:   bankAccount.AccountName,
// 			})
// 		}
// 		var contactList []dto.ContactDto
// 		for _, contact := range contacts {
// 			if person.ID == contact.PersonID {
// 				contactList = append(contactList, dto.ContactDto{
// 					ID:    contact.ID,
// 					Type:  contact.Type,
// 					Value: contact.Value,
// 				})
// 			}
// 		}
// 		result = append(result, dto.PersonDetailDto{
// 			ID:           person.ID,
// 			FullName:     person.FullName,
// 			Contacts:     contactList,
// 			BankAccounts: bank,
// 		})
// 	}
// 	return result, nil
// }

// // create person using CreatePersonDto
// func (service *PersonService) CreatePerson(personDto dto.CreatePersonDto) (*dto.PersonDto, error) {
// 	person := models_Person.Person{
// 		FullName:       personDto.FullName,
// 		IdentityNumber: personDto.IdentityNumber,
// 		IsActive:       true,
// 	}
// 	personModel, err := service.personRepo.CreatePerson(&person)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &dto.PersonDto{
// 		ID:             personModel.ID,
// 		FullName:       personModel.FullName,
// 		IdentityNumber: personModel.IdentityNumber,
// 	}, nil
// }

// // create contact using CreatePersonDto
// func (service *PersonService) CreateContact(personID uint, typeContact string, valuueContact string) error {
// 	contact := models_Person.Contact{
// 		PersonID: personID,
// 		Type:     typeContact,
// 		Value:    valuueContact,
// 		IsActive: true,
// 	}
// 	_, err := service.personRepo.CreateContact(&contact)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// // create bank account intitial using only personID
// func (service *PersonService) CreateBankAccount(personID uint) error {
// 	bankAccount := models_Person.BankAccount{
// 		PersonID: personID,
// 		IsActive: true,
// 	}
// 	_, err := service.personRepo.CreateBankAccount(&bankAccount)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// // edit person using EditPeopleDto
// func (service *PersonService) UpdatePerson(personDto dto.EditPeopleDto) (*dto.PersonDto, error) {
// 	existingPerson, err := service.personRepo.FindPersonById(personDto.PersonID)
// 	if err != nil {
// 		return nil, err
// 	}

// 	existingPerson.FullName = personDto.FullName
// 	existingPerson.IdentityNumber = personDto.IdentityNumber

// 	personModel, err := service.personRepo.UpdatePerson(existingPerson)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &dto.PersonDto{
// 		ID:             personModel.ID,
// 		FullName:       personModel.FullName,
// 		IdentityNumber: personModel.IdentityNumber,
// 	}, nil
// }

// // edit contact using CreateContactDto
// func (service *PersonService) UpdateContact(contactDto dto.EditContactDto) (*dto.ContactDto, error) {
// 	existingContact, err := service.personRepo.FindContactById(contactDto.ID)
// 	if err != nil {
// 		return nil, err
// 	}

// 	existingContact.Type = contactDto.TypeContact
// 	existingContact.Value = contactDto.ValueContact

// 	contactModel, err := service.personRepo.UpdateContact(existingContact)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &dto.ContactDto{
// 		ID:    contactModel.ID,
// 		Type:  contactModel.Type,
// 		Value: contactModel.Value,
// 	}, nil
// }

// // delete contact using parameter contactID
// func (service *PersonService) DeleteContact(contactID uint) error {
// 	_, err := service.personRepo.DeleteContact(contactID)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// // edit bank account using BankAccountDto
// func (service *PersonService) UpdateBankAccount(bankAccountDto dto.BankAccountDto) (*dto.BankAccountDto, error) {
// 	exitstingBankAccount, err := service.personRepo.FindBankAccountById(bankAccountDto.ID)
// 	if err != nil {
// 		return nil, err
// 	}

// 	exitstingBankAccount.BankName = bankAccountDto.BankName
// 	exitstingBankAccount.BankAddress = bankAccountDto.BankAddress
// 	exitstingBankAccount.AccountName = bankAccountDto.AccountName
// 	exitstingBankAccount.AccountNumber = bankAccountDto.AccountNumber
// 	exitstingBankAccount.SwiftCode = bankAccountDto.SwiftCode

// 	bankAccountModel, err := service.personRepo.UpdateBankAccount(exitstingBankAccount)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &dto.BankAccountDto{
// 		ID:            bankAccountModel.ID,
// 		PersonId:      bankAccountModel.PersonID,
// 		AccountNumber: bankAccountModel.AccountNumber,
// 		BankName:      bankAccountModel.BankName,
// 		AccountName:   bankAccountModel.AccountName,
// 	}, nil
// }
