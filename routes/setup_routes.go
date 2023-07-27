package routes

import (
	"PattayaAvenueProperty/app"

	"github.com/gin-gonic/gin"
)

func SetUpRoutes(router *gin.RouterGroup, ctrl *app.SetupController) (*gin.RouterGroup, error) {
	// Set up routes' handlers
	route := router.Group("api")
	{
		personRoute := route.Group("persons")
		{
			setupPersonRoutes(personRoute, ctrl.PersonController)
		}
		roomRoute := route.Group("rooms")
		{
			setupRoomRoutes(roomRoute, ctrl.RoomController)
		}
	}
	return route, nil
}
