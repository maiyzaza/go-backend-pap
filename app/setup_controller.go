package app

import (
	"PattayaAvenueProperty/controller"
)

type SetupController struct {
	Service          SetupService
	PersonController controller.PersonController
	RoomController   controller.RoomController
}

func (s *SetupController) Setup() {
	s.PersonController = controller.NewPersonController(
		s.Service.PersonService,
	)
	s.RoomController = controller.NewRoomController(
		s.Service.RoomService,
	)
}
