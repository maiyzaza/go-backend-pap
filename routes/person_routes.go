package routes

import (
	"PattayaAvenueProperty/controller"

	"github.com/gin-gonic/gin"
)

func setupPersonRoutes(r *gin.RouterGroup, customerController controller.PersonController) {
	r.GET("/profiles", customerController.GetProfiles)
}
