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
			ID:                  person.ID,
			FullName:            person.FullName,
			CitizenDocumentUrl:  person.CitizenDocumentUrl,
			PassportDocumentUrl: person.PassportDocumentUrl,
		})
	}
	return result, nil
}

func (service *PersonService) GetProfilesWithBankAccount() ([]dto.PersonDto, error) {
	persons, err := service.personRepo.FindAll()
	if err != nil {
		return nil, err
	}
	var result []dto.PersonDto
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
		result = append(result, dto.PersonDto{
			ID:                  person.ID,
			FullName:            person.FullName,
			CitizenDocumentUrl:  person.CitizenDocumentUrl,
			PassportDocumentUrl: person.PassportDocumentUrl,
			BankAccounts:        bank,
		})
	}
	return result, nil
}
