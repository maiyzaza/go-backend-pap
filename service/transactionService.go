package service

import (
	models_Document "PattayaAvenueProperty/models/Document"
	models_Transaction "PattayaAvenueProperty/models/Transaction"
	"PattayaAvenueProperty/repository"
	dto "PattayaAvenueProperty/service/dto"
	"time"
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
		var AddSevenHrs = transaction.CreatedAt.Add(7 * time.Hour)

		result = append(result, dto.TransactionResponseDto{
			ID:            transaction.ID,
			CatorgoryType: transaction.CategoryType,
			RoomAddress:   roomAddress,
			PaymentMethod: transaction.PaymentMethod,
			Amount:        transaction.Amount,
			IsReceive:     transaction.IsReceive,
			CreateAt:      AddSevenHrs.String(),
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

	var AddSevenHrs = transaction.CreatedAt.Add(7 * time.Hour)

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
		CreateAt:            AddSevenHrs.String(),
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
		IsActive:      true,
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
		IsActive:    true,
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

		var AddSevenHrs = transaction.CreatedAt.Add(7 * time.Hour)

		result = append(result, dto.TransactionResponseDto{
			ID:            transaction.ID,
			CatorgoryType: transaction.CategoryType,
			RoomAddress:   roomAddress,
			PaymentMethod: transaction.PaymentMethod,
			Amount:        transaction.Amount,
			IsReceive:     transaction.IsReceive,
			CreateAt:      AddSevenHrs.String(),
		})
	}
	return result, nil
}
