package service

import (
	models_Person "PattayaAvenueProperty/models/Person"
	"PattayaAvenueProperty/repository"

	dto "PattayaAvenueProperty/service/dto"
)

type PersonService struct {
	personRepo repository.PersonRepo
}

func NewPersonService(personRepo repository.PersonRepo) PersonService {
	return PersonService{
		personRepo: personRepo,
	}
}

func (service *PersonService) GetProfiles() ([]dto.PersonDto, error) {
	persons, err := service.personRepo.FindAll()
	if err != nil {
		return nil, err
	}
	var result []dto.PersonDto
	for _, person := range persons {
		result = append(result, dto.PersonDto{
			ID:             person.ID,
			FullName:       person.FullName,
			IdentityNumber: person.IdentityNumber,
		})
	}
	return result, nil
}

func (service *PersonService) GetProfilesWithBankAccount() ([]dto.PersonDetailDto, error) {
	persons, err := service.personRepo.FindAll()
	if err != nil {
		return nil, err
	}
	contacts, err := service.personRepo.GetAllContact()
	if err != nil {
		return nil, err
	}

	var result []dto.PersonDetailDto
	for _, person := range persons {
		var bank []dto.BankAccountDto
		for _, bankAccount := range person.BankAccounts {
			bank = append(bank, dto.BankAccountDto{
				ID:            bankAccount.ID,
				PersonId:      bankAccount.PersonID,
				AccountNumber: bankAccount.AccountNumber,
				BankAddress:   bankAccount.BankAddress,
				BankName:      bankAccount.BankName,
				AccountName:   bankAccount.AccountName,
			})
		}
		var contactList []dto.ContactDto
		for _, contact := range contacts {
			if person.ID == contact.PersonID {
				contactList = append(contactList, dto.ContactDto{
					ID:    contact.ID,
					Type:  contact.Type,
					Value: contact.Value,
				})
			}
		}
		result = append(result, dto.PersonDetailDto{
			ID:           person.ID,
			FullName:     person.FullName,
			Contacts:     contactList,
			BankAccounts: bank,
		})
	}
	return result, nil
}

// create person using CreatePersonDto
func (service *PersonService) CreatePerson(personDto dto.CreatePersonDto) (*dto.PersonDto, error) {
	person := models_Person.Person{
		FullName:       personDto.FullName,
		IdentityNumber: personDto.IdentityNumber,
		IsActive:       true,
	}
	personModel, err := service.personRepo.CreatePerson(&person)
	if err != nil {
		return nil, err
	}
	return &dto.PersonDto{
		ID:             personModel.ID,
		FullName:       personModel.FullName,
		IdentityNumber: personModel.IdentityNumber,
	}, nil
}

// create contact using CreatePersonDto
func (service *PersonService) CreateContact(personID uint, typeContact string, valuueContact string) error {
	contact := models_Person.Contact{
		PersonID: personID,
		Type:     typeContact,
		Value:    valuueContact,
		IsActive: true,
	}
	_, err := service.personRepo.CreateContact(&contact)
	if err != nil {
		return err
	}
	return nil
}

// create bank account intitial using only personID
func (service *PersonService) CreateBankAccount(personID uint) error {
	bankAccount := models_Person.BankAccount{
		PersonID: personID,
		IsActive: true,
	}
	_, err := service.personRepo.CreateBankAccount(&bankAccount)
	if err != nil {
		return err
	}
	return nil
}

// edit person using EditPeopleDto
func (service *PersonService) UpdatePerson(personDto dto.EditPeopleDto) (*dto.PersonDto, error) {
	existingPerson, err := service.personRepo.FindPersonById(personDto.PersonID)
	if err != nil {
		return nil, err
	}

	existingPerson.FullName = personDto.FullName
	existingPerson.IdentityNumber = personDto.IdentityNumber

	personModel, err := service.personRepo.UpdatePerson(existingPerson)
	if err != nil {
		return nil, err
	}
	return &dto.PersonDto{
		ID:             personModel.ID,
		FullName:       personModel.FullName,
		IdentityNumber: personModel.IdentityNumber,
	}, nil
}

// edit contact using CreateContactDto
func (service *PersonService) UpdateContact(contactDto dto.EditContactDto) (*dto.ContactDto, error) {
	existingContact, err := service.personRepo.FindContactById(contactDto.ID)
	if err != nil {
		return nil, err
	}

	existingContact.Type = contactDto.TypeContact
	existingContact.Value = contactDto.ValueContact

	contactModel, err := service.personRepo.UpdateContact(existingContact)
	if err != nil {
		return nil, err
	}
	return &dto.ContactDto{
		ID:    contactModel.ID,
		Type:  contactModel.Type,
		Value: contactModel.Value,
	}, nil
}

// delete contact using parameter contactID
func (service *PersonService) DeleteContact(contactID uint) error {
	_, err := service.personRepo.DeleteContact(contactID)
	if err != nil {
		return err
	}
	return nil
}

// edit bank account using BankAccountDto
func (service *PersonService) UpdateBankAccount(bankAccountDto dto.BankAccountDto) (*dto.BankAccountDto, error) {
	exitstingBankAccount, err := service.personRepo.FindBankAccountById(bankAccountDto.ID)
	if err != nil {
		return nil, err
	}

	exitstingBankAccount.BankName = bankAccountDto.BankName
	exitstingBankAccount.BankAddress = bankAccountDto.BankAddress
	exitstingBankAccount.AccountName = bankAccountDto.AccountName
	exitstingBankAccount.AccountNumber = bankAccountDto.AccountNumber
	exitstingBankAccount.SwiftCode = bankAccountDto.SwiftCode

	bankAccountModel, err := service.personRepo.UpdateBankAccount(exitstingBankAccount)
	if err != nil {
		return nil, err
	}
	return &dto.BankAccountDto{
		ID:            bankAccountModel.ID,
		PersonId:      bankAccountModel.PersonID,
		AccountNumber: bankAccountModel.AccountNumber,
		BankName:      bankAccountModel.BankName,
		AccountName:   bankAccountModel.AccountName,
	}, nil
}
