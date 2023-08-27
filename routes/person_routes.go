package routes

import (
	"PattayaAvenueProperty/controller"

	"github.com/gin-gonic/gin"
)

func setupPersonRoutes(r *gin.RouterGroup, personController controller.PersonController) {
	r.GET("/profiles", personController.GetProfiles)
	r.GET("/profiles/bank/:personID", personController.GetProfilesWithBankAccountByPersonID)
	r.POST("/createprofiles", personController.CreatePerson)
	r.POST("/editprofiles", personController.UpdatePerson)
	r.POST("/editcontact", personController.UpdateContact)
	r.POST("/deletecontact/:contactID", personController.DeleteContact)
	r.POST("/editbankaccount", personController.UpdateBankAccount)
	r.POST("/createcontact", personController.CreateContact)
	r.GET("/contact/:personID", personController.GetAllContactByPersonID)
}
