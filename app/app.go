package app

import (
	"fmt"

	"github.com/gin-gonic/gin"
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

	router := gin.Default()
	err := router.Run(":9888")
	if err != nil {
		panic("[Error] failed to start Gin server due to: " + err.Error())
	}
	fmt.Println("Gin server is running on port 9888")
	fmt.Println("App is running app.go")

	return &app
}
