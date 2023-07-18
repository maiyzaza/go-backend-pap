package app

import "PattayaAvenueProperty/repository"

type SetupRepository struct {
	PersonRepo repository.PersonRepo
}

func (s *SetupRepository) Setup() {
	s.PersonRepo = repository.NewPersonRepo()

}
