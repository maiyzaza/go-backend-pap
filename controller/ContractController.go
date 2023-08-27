package controller

import (
	"PattayaAvenueProperty/constants"
	"PattayaAvenueProperty/models/handler"
	"PattayaAvenueProperty/service"
	"PattayaAvenueProperty/service/dto"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ContractController struct {
	contractService service.ContractService
}

func NewContractController(contractService service.ContractService) ContractController {
	return ContractController{contractService: contractService}
}

func (controller *ContractController) GetAllContract(c *gin.Context) {
	data, err := controller.contractService.GetAllContract()
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

func (controller *ContractController) GetContractByRoomID(c *gin.Context) {
	roomID := c.Param("roomID")
	roomIDUint, err := strconv.ParseUint(roomID, 10, 64)
	if err != nil {
		_ = c.Error(err)
		return
	}
	data, err := controller.contractService.GetContractByRoomID(uint(roomIDUint))
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

func (controller *ContractController) GetRoomContractByID(c *gin.Context) {
	roomContractID := c.Param("roomContractID")
	roomContractIDUint, err := strconv.ParseUint(roomContractID, 10, 64)
	if err != nil {
		_ = c.Error(err)
		return
	}
	data, err := controller.contractService.GetRoomContractByID(uint(roomContractIDUint))
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

// create room contract then create person contract by using room contract id
func (controller *ContractController) CreateRoomContract(c *gin.Context) {
	var requestDto dto.CreateRoomContractDto
	if err := c.ShouldBindJSON(&requestDto); err != nil {
		_ = c.Error(err)
		return
	}

	roomContract, err := controller.contractService.CreateRoomContract(requestDto)
	if err != nil {
		_ = c.Error(err)
		return
	}

	err = controller.contractService.CreatePersonContract(roomContract.ID, requestDto.PersonID, requestDto.PersonContractType)
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

// close room contract by using room contract id and this data CloseRoomContractDto
func (controller *ContractController) CloseRoomContract(c *gin.Context) {
	var requestDto dto.CloseRoomContractDto
	if err := c.ShouldBindJSON(&requestDto); err != nil {
		_ = c.Error(err)
		return
	}

	err := controller.contractService.UpdateRoomContract(requestDto.RoomContractID, requestDto)
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
