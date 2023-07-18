package app

import (
	"PattayaAvenueProperty/service"
)

type SetupService struct {
	Repository    SetupRepository
	PersonService service.PersonService
}

func (s *SetupService) Setup() {

	s.PersonService = service.NewPersonService(
		s.Repository.PersonRepo,
	)
}
