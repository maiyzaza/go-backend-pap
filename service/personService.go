package service

import (
	"PattayaAvenueProperty/repository"
	"fmt"

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
		fmt.Println("-------------------")
		fmt.Println(person)
		fmt.Println(person.BankAccounts)
		fmt.Println("-------------------")
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

func (service *PersonService) GetBuilding() ([]Dto.PlaceDto, error) {
	places, err := service.personRepo.GetAllPlace()
	if err != nil {
		return nil, err
	}
	var result []Dto.PlaceDto
	for _, place := range places {
		var buildings []Dto.BuildingDto
		for _, building := range place.Buildings {
			var floors []Dto.FloorDto
			for _, floor := range building.Floors {
				floors = append(floors, Dto.FloorDto{
					FloorID:     floor.ID,
					FloorNumber: floor.FloorNumber,
				})
			}
			buildings = append(buildings, Dto.BuildingDto{
				BuildingID:   building.ID,
				BuildingName: building.BuildingName,
				Floors:       floors,
			})
			result = append(result, Dto.PlaceDto{
				PlaceID:   place.ID,
				PlaceName: place.PlaceName,
				Buildings: buildings,
			})
		}
	}
	return result, nil
}
