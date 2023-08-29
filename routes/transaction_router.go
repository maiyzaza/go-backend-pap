package routes

import (
	"PattayaAvenueProperty/controller"

	"github.com/gin-gonic/gin"
)

func setupTransactionRoutes(r *gin.RouterGroup, transactionController controller.TransactionController) {
	r.GET("/transaction", transactionController.GetAllTransaction)
	r.GET("/transaction/:transactionID", transactionController.GetTransactionByID)
	r.POST("/transaction", transactionController.CreateTransaction)
	r.POST("/transaction/:transactionID", transactionController.DeleteTransaction)
	r.GET("/deletedtransaction", transactionController.GetAllDeletedTransaction)
	r.POST("/deletedtransaction/:transactionID", transactionController.DeleteTransaction)
}
