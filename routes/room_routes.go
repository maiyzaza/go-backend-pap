package routes

import (
	"PattayaAvenueProperty/controller"

	"github.com/gin-gonic/gin"
)

func setupRoomRoutes(r *gin.RouterGroup, roomController controller.RoomController) {
	r.GET("/places", roomController.GetAllPlace)
	r.POST("/places/:placeName", roomController.CreatePlace)
	r.POST("/building", roomController.CreateBuilding)
	r.POST("/floor", roomController.CreateFloor)
	r.POST("/room", roomController.CreateRoom)
	r.POST("/editroom/:roomID", roomController.EditRoom)
	r.GET("/room/:roomID", roomController.GetRoomByID)
}
