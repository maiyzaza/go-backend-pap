package app

import (
	"PattayaAvenueProperty/controller"
)

type SetupController struct {
	Service            SetupService
	PersonController   controller.PersonController
	RoomController     controller.RoomController
	ContractController controller.ContractController
}

func (s *SetupController) Setup() {
	s.PersonController = controller.NewPersonController(
		s.Service.PersonService,
	)
	s.RoomController = controller.NewRoomController(
		s.Service.RoomService,
	)
	s.ContractController = controller.NewContractController(
		s.Service.ContractService,
	)
}
