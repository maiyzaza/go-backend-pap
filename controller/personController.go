package controller

import (
	"PattayaAvenueProperty/constants"
	"PattayaAvenueProperty/models/handler"
	"PattayaAvenueProperty/service"
	"PattayaAvenueProperty/service/dto"
	"fmt"
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

func (controller *PersonController) GetProfilesWithBankAccount(c *gin.Context) {

	data, err := controller.personService.GetProfilesWithBankAccount()
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

// func (controller *RoomController) CreateRoomDocument(c *gin.Context) {
// 	var body dto.TakeRoomPictureDataDto

// 	if err := c.BindJSON(&body); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"message": "Invalid request body",
// 		})
// 		return
// 	}
// 	err := controller.roomService.CreateRoomDocument(body)
// 	if err != nil {
// 		c.Error(err)
// 		return
// 	}
// 	c.JSON(http.StatusOK, handler.Wrapper{
// 		StatusCode: http.StatusOK,
// 		Message:    constants.SUCCESS,
// 		Data:       nil,
// 	})
// }

// create person
func (controller *PersonController) CreatePerson(c *gin.Context) {
	var body dto.CreatePersonDto

	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request body",
		})
		return
	}

	person, err := controller.personService.CreatePerson(body)
	if err != nil {
		c.Error(err)
		return
	}
	fmt.Println(person)

	//create contact and bank account
	err1 := controller.personService.CreateContact(person.ID, body.TypeContact, body.ValueContact)
	if err != nil {
		c.Error(err1)
		return
	}

	err2 := controller.personService.CreateBankAccount(person.ID)
	if err != nil {
		c.Error(err2)
		return
	}

	c.JSON(http.StatusOK, handler.Wrapper{
		StatusCode: http.StatusOK,
		Message:    constants.SUCCESS,
		Data:       nil,
	})
}
