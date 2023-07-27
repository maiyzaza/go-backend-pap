package controller

import (
	"PattayaAvenueProperty/constants"
	"PattayaAvenueProperty/models/handler"
	"PattayaAvenueProperty/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RoomController struct {
	roomService service.RoomService
}

func NewRoomController(roomService service.RoomService) RoomController {
	return RoomController{roomService: roomService}
}

func (controller *RoomController) GetAllPlace(c *gin.Context) {

	data, err := controller.roomService.GetAllPlace()
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, handler.Wrapper{
		StatusCode: http.StatusOK,
		Message:    constants.SUCCESS,
		Data:       data,
	})
}

func (controller *RoomController) CreatePlace(c *gin.Context) {
	placeName := c.Param("placeName")

	err := controller.roomService.CreatePlace(placeName)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, handler.Wrapper{
		StatusCode: http.StatusOK,
		Message:    constants.SUCCESS,
		Data:       nil,
	})
}
