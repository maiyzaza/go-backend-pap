package app

import (
	"PattayaAvenueProperty/controller"
)

type SetupController struct {
	Service          SetupService
	PersonController controller.PersonController
}

func (s *SetupController) Setup() {
	s.PersonController = controller.NewPersonController(
		s.Service.PersonService,
	)
}
