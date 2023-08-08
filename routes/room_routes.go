package routes

import (
	"PattayaAvenueProperty/controller"

	"github.com/gin-gonic/gin"
)

func setupRoomRoutes(r *gin.RouterGroup, customerController controller.RoomController) {
	r.GET("/places", customerController.GetAllPlace)
	r.POST("/places/:placeName", customerController.CreatePlace)
	r.POST("/building", customerController.CreateBuilding)
	r.POST("/floor", customerController.CreateFloor)
	r.POST("/room", customerController.CreateRoom)
	r.POST("/editroom/:roomID", customerController.EditRoom)
}
