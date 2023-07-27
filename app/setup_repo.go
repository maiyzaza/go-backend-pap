package app

import "PattayaAvenueProperty/repository"

type SetupRepository struct {
	PersonRepo repository.PersonRepo
	RoomRepo   repository.RoomRepo
}

func (s *SetupRepository) Setup() {
	s.PersonRepo = repository.NewPersonRepo()
	s.RoomRepo = repository.NewRoomRepo()
}
