package controller

import (
	"PattayaAvenueProperty/constants"
	"PattayaAvenueProperty/models/handler"
	"PattayaAvenueProperty/service"
	dto "PattayaAvenueProperty/service/dto"
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

func (controller *RoomController) CreateBuilding(c *gin.Context) {
	var body dto.BuildingPostDto

	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request body",
		})
		return
	}

	err := controller.roomService.CreateBuilding(body.PlaceID, body.BuildingName)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, handler.Wrapper{
		StatusCode: http.StatusOK,
		Message:    constants.SUCCESS,
		Data:       nil,
	})
}
