package app

import (
	"fmt"
	"net/http"
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

	port := ":8080"

	// Start the server and listen on the specified port
	fmt.Printf("Server is listening on port %s...\n", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Printf("Error starting the server: %v", err)
	}
	fmt.Println("App is running app.go")

	return &app
}
