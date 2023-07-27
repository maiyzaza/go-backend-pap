package service

import (
	"PattayaAvenueProperty/repository"

	Dto "PattayaAvenueProperty/service/dto"
	// Dto "PattayaAvenueProperty/service/dto/person"
)

type PersonService struct {
	personRepo repository.PersonRepo
}

func NewPersonService(personRepo repository.PersonRepo) PersonService {
	return PersonService{
		personRepo: personRepo,
	}
}

func (service *PersonService) GetProfiles() ([]Dto.PersonDto, error) {
	persons, err := service.personRepo.FindAll()
	if err != nil {
		return nil, err
	}
	var result []Dto.PersonDto
	for _, person := range persons {
		var bank []Dto.BankAccountDto
		for _, bankAccount := range person.BankAccounts {
			bank = append(bank, Dto.BankAccountDto{
				ID:            bankAccount.ID,
				AccountNumber: bankAccount.AccountNumber,
				BankName:      bankAccount.BankName,
				AccountName:   bankAccount.AccountName,
			})
		}
		result = append(result, Dto.PersonDto{
			ID:                  person.ID,
			FullName:            person.FullName,
			CitizenDocumentUrl:  person.CitizenDocumentUrl,
			PassportDocumentUrl: person.PassportDocumentUrl,
			BankAccounts:        bank,
		})
	}
	return result, nil
}
