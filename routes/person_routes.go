package routes

import (
	"PattayaAvenueProperty/controller"

	"github.com/gin-gonic/gin"
)

func setupPersonRoutes(r *gin.RouterGroup, personController controller.PersonController) {
	r.GET("/profiles", personController.GetProfiles)
}
