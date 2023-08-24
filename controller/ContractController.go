package controller

import (
	"PattayaAvenueProperty/constants"
	"PattayaAvenueProperty/models/handler"
	"PattayaAvenueProperty/service"
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
