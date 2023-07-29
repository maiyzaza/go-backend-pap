package service

import (
	models_Room "PattayaAvenueProperty/models/Room"
	"PattayaAvenueProperty/repository"
	"fmt"

	Dto "PattayaAvenueProperty/service/dto"
)

type RoomService struct {
	roomRepo repository.RoomRepo
}

func NewRoomService(roomRepo repository.RoomRepo) RoomService {
	return RoomService{
		roomRepo: roomRepo,
	}
}

func (service *RoomService) GetAllPlace() ([]Dto.PlaceDto, error) {
	places, err := service.roomRepo.GetAllPlace()
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
		}
		result = append(result, Dto.PlaceDto{
			PlaceID:   place.ID,
			PlaceName: place.PlaceName,
			Buildings: buildings,
		})
	}
	return result, nil
}

func (service *RoomService) CreatePlace(placeName string) error {
	placeModel := models_Room.Place{
		PlaceName: placeName,
		Buildings: []models_Room.Building{},
		IsActive:  true,
	}

	_, err := service.roomRepo.CreatePlace(placeModel)

	if err != nil {
		return err
	}
	return nil
}

func (service *RoomService) CreateBuilding(placeID uint, buildingName string) error {
	buildingModel := models_Room.Building{
		BuildingName: buildingName,
		PlaceID:      placeID,
		Floors:       []models_Room.Floor{},
		IsActive:     true,
	}

	fmt.Println(buildingModel)
	_, err := service.roomRepo.CreateBuilding(buildingModel)

	if err != nil {
		return err
	}
	return nil
}
