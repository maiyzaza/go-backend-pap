package app

import (
	"PattayaAvenueProperty/service"
)

type SetupService struct {
	Repository         SetupRepository
	PersonService      service.PersonService
	RoomService        service.RoomService
	ContractService    service.ContractService
	TransactionService service.TransactionService
}

func (s *SetupService) Setup() {

	s.PersonService = service.NewPersonService(
		s.Repository.PersonRepo,
	)
	s.RoomService = service.NewRoomService(
		s.Repository.RoomRepo,
		s.Repository.PersonRepo,
	)
	s.ContractService = service.NewContractService(
		s.Repository.ContractRepo,
		s.Repository.PersonRepo,
		s.Repository.RoomRepo,
	)
	s.TransactionService = service.NewTransactionService(
		s.Repository.TransactionRepo,
		s.Repository.RoomRepo,
		s.Repository.ContractRepo,
	)
}
