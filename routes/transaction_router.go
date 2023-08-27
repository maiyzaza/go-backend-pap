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
}

// package routes

// import (
// 	"PattayaAvenueProperty/controller"

// 	"github.com/gin-gonic/gin"
// )

// func setupRoomRoutes(r *gin.RouterGroup, roomController controller.RoomController) {
// 	r.GET("/places", roomController.GetAllPlace)
// 	r.POST("/places/:placeName", roomController.CreatePlace)
// 	r.POST("/building", roomController.CreateBuilding)
// 	r.POST("/floor", roomController.CreateFloor)
// 	r.POST("/room", roomController.CreateRoom)
// 	r.POST("/editroom/:roomID", roomController.EditRoom)
// 	r.GET("/room/:roomID", roomController.GetRoomByID)
// 	r.POST("/editplace/:placeID/:placeName", roomController.ModifyPlace)
// 	r.POST("/editbuilding/:buildingID/:buildingName", roomController.ModifyBuilding)
// 	r.POST("/roomprice", roomController.CreateRoomPrice)
// 	r.POST("/editroomprice/:roomPriceID", roomController.DeleteRoomPrice)
// 	r.POST("/roompicture", roomController.CreateRoomPicture)
// 	r.POST("/editroompicture/:roomPictureID", roomController.DeleteRoomPicture)
// 	r.POST("/roomdocument", roomController.CreateRoomDocument)
// 	r.POST("/editroomdocument/:roomDocumentID", roomController.DeleteRoomDocument)
// }
