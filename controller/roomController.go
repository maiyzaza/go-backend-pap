package controller

import (
	"PattayaAvenueProperty/constants"
	models "PattayaAvenueProperty/models/Room"
	"PattayaAvenueProperty/models/handler"
	"PattayaAvenueProperty/service"
	"PattayaAvenueProperty/service/dto"
	"net/http"
	"strconv"

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
	var body dto.CreateBuildingDto

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

func (controller *RoomController) CreateFloor(c *gin.Context) {
	var body dto.CreateFloorDto

	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request body",
		})
		return
	}

	err := controller.roomService.CreateFloor(body.BuildingID, body.FloorNumber)
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

func (controller *RoomController) CreateRoom(c *gin.Context) {
	var body dto.CreateRoomDto

	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request body",
		})
		return
	}

	err := controller.roomService.CreateRoom(body.FloorID, body.RoomNumber)
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

func (controller *RoomController) EditRoom(c *gin.Context) {
	var body dto.EditRoomInfoDto

	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request body",
		})
		return
	}

	roomID := c.Param("roomID")
	// Convert roomID to uint
	roomIDUint, err := strconv.ParseUint(roomID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid room ID",
		})
		return
	}

	updatedRoom := models.Room{
		RoomName:           &body.RoomName, // Convert string to *string
		RoomNumber:         body.RoomNumber,
		RoomAddress:        body.RoomAddress,
		ElectricNumber:     &body.ElectricNumber,     // Convert string to *string
		ElectricUserNumber: &body.ElectricUserNumber, // Convert string to *string
		AmountOfBedRoom:    body.AmountOfBedRoom,
		AmountOfToiletRoom: body.AmountOfToiletRoom,
		AmountOfLivingRoom: body.AmountOfLivingRoom,
		SizeSQM:            body.SizeSQM,
		TypeOfView:         body.TypeOfView,
		Remark:             &body.Remark, // Convert string to *string
		StatusOfRoom:       body.StatusOfRoom,
	}

	modifiedRoom, err := controller.roomService.ModifyRoom(uint(roomIDUint), updatedRoom)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, handler.Wrapper{
		StatusCode: http.StatusOK,
		Message:    constants.SUCCESS,
		Data:       modifiedRoom,
	})
}

func (controller *RoomController) GetRoomByID(c *gin.Context) {
	roomID := c.Param("roomID")
	// Convert roomID to uint
	roomIDUint, err := strconv.ParseUint(roomID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid room ID",
		})
		return
	}

	room, err := controller.roomService.GetRoomByID(uint(roomIDUint))
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, handler.Wrapper{
		StatusCode: http.StatusOK,
		Message:    constants.SUCCESS,
		Data:       room,
	})
}
