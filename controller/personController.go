package controller

import (
	"PattayaAvenueProperty/constants"
	"PattayaAvenueProperty/models/handler"
	"PattayaAvenueProperty/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PersonController struct {
	personService service.PersonService
}

func NewPersonController(personService service.PersonService) PersonController {
	return PersonController{personService: personService}
}

func (controller *PersonController) GetProfiles(c *gin.Context) {

	data, err := controller.personService.GetProfiles()
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
