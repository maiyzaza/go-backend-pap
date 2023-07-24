package app

import (
	"fmt"
)

type App struct {
	Controller *SetupController
	Repository *SetupRepository
	Service    *SetupService
}

func New() *App {

	repository := &SetupRepository{}
	repository.Setup()

	service := &SetupService{Repository: *repository}
	service.Setup()

	controller := &SetupController{Service: *service}
	controller.Setup()

	app := App{
		Repository: repository,
		Controller: controller,
	}
	fmt.Println("App is running app.go")

	return &app
}
