package app

import (
	"PattayaAvenueProperty/service"
)

type SetupService struct {
	Repository    SetupRepository
	PersonService service.PersonService
	RoomService   service.RoomService
}

func (s *SetupService) Setup() {

	s.PersonService = service.NewPersonService(
		s.Repository.PersonRepo,
	)
	s.RoomService = service.NewRoomService(
		s.Repository.RoomRepo,
		s.Repository.PersonRepo,
	)
}
