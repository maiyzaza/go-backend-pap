package service

import (
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
				AccountNumber: bankAccount.AccountNumber,
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
