package app

import "PattayaAvenueProperty/repository"

type SetupRepository struct {
	PersonRepo   repository.PersonRepo
	RoomRepo     repository.RoomRepo
	ContractRepo repository.ContractRepo
}

func (s *SetupRepository) Setup() {
	s.PersonRepo = repository.NewPersonRepo()
	s.RoomRepo = repository.NewRoomRepo()
	s.ContractRepo = repository.NewContractRepo()
}
