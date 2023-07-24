package web

import (
	"PattayaAvenueProperty/app"
	"PattayaAvenueProperty/middleware"

	"PattayaAvenueProperty/routes"

	"github.com/gin-gonic/gin"
)

func StartServer(app *app.App) {
	r := gin.Default()
	r.Use(middleware.ErrorHandle())
	r.Use(middleware.CORSMiddleware())
	r.Run(":8080")

	if _, err := routes.SetUpRoutes(&r.RouterGroup, app.Controller); err != nil {
		panic(err)
	}

	if err := r.Run(":8080"); err != nil {
		panic("Failed to start the project")
	}
}
