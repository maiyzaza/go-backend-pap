package routes

import (
	"PattayaAvenueProperty/controller"

	"github.com/gin-gonic/gin"
)

func setupContractRoutes(r *gin.RouterGroup, contractController controller.ContractController) {
	r.GET("/contracts", contractController.GetAllContract)
	r.GET("/contracts/room/:roomID", contractController.GetContractByRoomID)
	r.GET("/contracts/:roomContractID", contractController.GetRoomContractByID)
	r.POST("/createcontracts", contractController.CreateRoomContract)
	r.POST("/closecontracts", contractController.CloseRoomContract)
}
