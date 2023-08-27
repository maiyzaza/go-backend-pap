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

func (controller *PersonController) GetProfilesWithBankAccountByPersonID(c *gin.Context) {

	personID := c.Param("personID")

	personIDUint, err := strconv.ParseUint(personID, 10, 64)
	if err != nil {
		c.Error(err)
		return
	}

	data, err := controller.personService.GetProfilesWithBankAccountByPersonID(uint(personIDUint))
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

// update person using EditPeopleDto
func (controller *PersonController) UpdatePerson(c *gin.Context) {
	var body dto.EditPeopleDto

	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request body",
		})
		return
	}

	_, err := controller.personService.UpdatePerson(body)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, handler.Wrapper{
		StatusCode: http.StatusOK,
		Message:    constants.SUCCESS,
		Data:       nil,
	})
}

// create contact using CreateContactDto
func (controller *PersonController) CreateContact(c *gin.Context) {
	var body dto.CreateContactDto

	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request body",
		})
		return
	}

	err := controller.personService.CreateContact(body.PersonID, body.TypeContact, body.ValueContact)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, handler.Wrapper{
		StatusCode: http.StatusOK,
		Message:    constants.SUCCESS,
		Data:       nil,
	})
}

// edit contact using EditContactDto
func (controller *PersonController) UpdateContact(c *gin.Context) {
	var body dto.EditContactDto

	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request body",
		})
		return
	}

	_, err := controller.personService.UpdateContact(body)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, handler.Wrapper{
		StatusCode: http.StatusOK,
		Message:    constants.SUCCESS,
		Data:       nil,
	})
}

// delete contact using parameter contactID
func (controller *PersonController) DeleteContact(c *gin.Context) {
	contactID := c.Param("contactID")

	contactIDUint, err := strconv.ParseUint(contactID, 10, 64)
	if err != nil {
		c.Error(err)
		return
	}
	err1 := controller.personService.DeleteContact(uint(contactIDUint))
	if err1 != nil {
		c.Error(err1)
		return
	}

	c.JSON(http.StatusOK, handler.Wrapper{
		StatusCode: http.StatusOK,
		Message:    constants.SUCCESS,
		Data:       nil,
	})
}

// edit bank account using BankAccountDto
func (controller *PersonController) UpdateBankAccount(c *gin.Context) {
	var body dto.BankAccountDto

	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request body",
		})
		return
	}

	_, err := controller.personService.UpdateBankAccount(body)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, handler.Wrapper{
		StatusCode: http.StatusOK,
		Message:    constants.SUCCESS,
		Data:       nil,
	})
}

// GetAllContact ByPersonID
func (controller *PersonController) GetAllContactByPersonID(c *gin.Context) {
	personID := c.Param("personID")

	personIDUint, err := strconv.ParseUint(personID, 10, 64)
	if err != nil {
		c.Error(err)
		return
	}

	data, err := controller.personService.GetAllContactByPersonID(uint(personIDUint))
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, handler.Wrapper{
		StatusCode: http.StatusOK,
		Message:    constants.SUCCESS,
		Data:       data,
	})
}
