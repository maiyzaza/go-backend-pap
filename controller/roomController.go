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
		OwnerID:            &body.OwnerID,            // Convert string to *string
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

func (controller *RoomController) ModifyPlace(c *gin.Context) {
	placeID := c.Param("placeID")
	// Convert placeID to uint
	placeIDUint, err := strconv.ParseUint(placeID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid place ID",
		})
		return
	}

	placeName := c.Param("placeName")

	err = controller.roomService.ModifyPlace(uint(placeIDUint), placeName)
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

// modify building
func (controller *RoomController) ModifyBuilding(c *gin.Context) {
	buildingID := c.Param("buildingID")
	// Convert buildingID to uint
	buildingIDUint, err := strconv.ParseUint(buildingID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid building ID",
		})
		return
	}

	buildingName := c.Param("buildingName")

	err = controller.roomService.ModifyBuilding(uint(buildingIDUint), buildingName)
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

// create RoomPrice and delete RoomPrice by change is_active to 0
func (controller *RoomController) CreateRoomPrice(c *gin.Context) {
	var body dto.TakeRoomPriceDataDto

	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request body",
		})
		return
	}
	err := controller.roomService.CreateRoomPrice(body)
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

func (controller *RoomController) DeleteRoomPrice(c *gin.Context) {
	roomPriceID := c.Param("roomPriceID")
	// Convert roomPriceID to uint
	roomPriceIDUint, err := strconv.ParseUint(roomPriceID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid roomPrice ID",
		})
		return
	}

	err = controller.roomService.DeleteRoomPrice(uint(roomPriceIDUint))
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

// create RoomPicture and delete RoomPicture by change is_active to 0
func (controller *RoomController) CreateRoomPicture(c *gin.Context) {
	var body dto.TakeRoomPictureDataDto

	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request body",
		})
		return
	}
	roomPictureData, err := controller.roomService.CreateRoomPicture(body)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, handler.Wrapper{
		StatusCode: http.StatusOK,
		Message:    constants.SUCCESS,
		Data:       roomPictureData,
	})
}

func (controller *RoomController) DeleteRoomPicture(c *gin.Context) {
	roomPictureID := c.Param("roomPictureID")
	// Convert roomPictureID to uint
	roomPictureIDUint, err := strconv.ParseUint(roomPictureID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid roomPicture ID",
		})
		return
	}

	err = controller.roomService.DeleteRoomPicture(uint(roomPictureIDUint))
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

// create RoomDocument and delete RoomDocument by change is_active to 0
func (controller *RoomController) CreateRoomDocument(c *gin.Context) {
	var body dto.TakeRoomDocumentDataDto

	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request body",
		})
		return
	}

	roomDocumentData, err := controller.roomService.CreateRoomDocument(body)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, handler.Wrapper{
		StatusCode: http.StatusOK,
		Message:    constants.SUCCESS,
		Data:       roomDocumentData,
	})
}

func (controller *RoomController) DeleteRoomDocument(c *gin.Context) {
	roomDocumentID := c.Param("roomDocumentID")
	// Convert roomDocumentID to uint
	roomDocumentIDUint, err := strconv.ParseUint(roomDocumentID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid roomDocument ID",
		})
		return
	}

	err = controller.roomService.DeleteRoomDocument(uint(roomDocumentIDUint))
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
