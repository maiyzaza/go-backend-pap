package service

import (
	models "PattayaAvenueProperty/models/Room"
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
	buildings, err := service.roomRepo.GetAllBuilding()
	if err != nil {
		return nil, err
	}
	floors, err := service.roomRepo.GetAllFloor()
	if err != nil {
		return nil, err
	}
	rooms, err := service.roomRepo.GetAllRoom()
	if err != nil {
		return nil, err
	}

	// Create a map to store rooms indexed by floor ID
	roomsByFloorID := make(map[uint][]Dto.RoomDto)
	for _, room := range rooms {
		roomsByFloorID[room.FloorID] = append(roomsByFloorID[room.FloorID], Dto.RoomDto{
			RoomID:   room.ID,
			RoomName: *room.RoomName,
		})
	}
	fmt.Println(roomsByFloorID)

	var result []Dto.PlaceDto
	for _, place := range places {
		var buildingsDto []Dto.BuildingDto
		for _, building := range buildings {
			if building.PlaceID == place.ID {
				var floorsDto []Dto.FloorDto
				for _, floor := range floors {
					fmt.Println(floor.BuildingID, building.ID)
					if floor.BuildingID == building.ID {
						// Include floors only if there are rooms associated with them
						if roomsForFloor, ok := roomsByFloorID[floor.ID]; ok {
							floorsDto = append(floorsDto, Dto.FloorDto{
								FloorID:     floor.ID,
								FloorNumber: floor.FloorNumber,
								Rooms:       roomsForFloor,
							})
						} else {
							floorsDto = append(floorsDto, Dto.FloorDto{
								FloorID:     floor.ID,
								FloorNumber: floor.FloorNumber,
								Rooms:       nil,
							})
						}
					}
				}
				fmt.Println(floorsDto)
				buildingsDto = append(buildingsDto, Dto.BuildingDto{
					BuildingID:   building.ID,
					BuildingName: building.BuildingName,
					Floors:       floorsDto,
				})
			}
		}
		result = append(result, Dto.PlaceDto{
			PlaceID:   place.ID,
			PlaceName: place.PlaceName,
			Buildings: buildingsDto,
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

func (service *RoomService) CreateFloor(buildingID uint, floorNumber string) error {
	floorModel := models_Room.Floor{
		FloorNumber: floorNumber,
		BuildingID:  buildingID,
		IsActive:    true,
	}

	_, err := service.roomRepo.CreateFloor(floorModel)

	if err != nil {
		return err
	}
	return nil
}

func (service *RoomService) CreateRoom(floorID uint, roomName string) error {
	roomModel := models_Room.Room{
		RoomName: &roomName,
		FloorID:  floorID,
		IsActive: true,
	}

	_, err := service.roomRepo.CreateRoom(roomModel)

	if err != nil {
		return err
	}
	return nil
}

func (service *RoomService) GetRoomByID(roomID uint) (*Dto.RoomResponseDto, error) {
	room, err := service.roomRepo.GetRoomByID(roomID)
	if err != nil {
		return nil, err
	}

	// Convert to RoomResponse model
	roomResponse := Dto.RoomResponseDto{
		ID:                 room.ID,
		OwnerID:            room.OwnerID,
		FloorID:            room.FloorID,
		RoomName:           room.RoomName,
		RoomNumber:         room.RoomNumber,
		RoomAddress:        room.RoomAddress,
		ElectricNumber:     room.ElectricNumber,
		ElectricUserNumber: room.ElectricUserNumber,
		AmountOfBedRoom:    room.AmountOfBedRoom,
		AmountOfToiletRoom: room.AmountOfToiletRoom,
		AmountOfLivingRoom: room.AmountOfLivingRoom,
		SizeSQM:            room.SizeSQM,
		TypeOfView:         room.TypeOfView,
		Remark:             room.Remark,
		StatusOfRoom:       room.StatusOfRoom,
		IsActive:           room.IsActive,
		CreatedAt:          room.CreatedAt.String(),
		UpdatedAt:          room.UpdatedAt.String(),
	}

	return &roomResponse, nil
}

func (service *RoomService) ModifyRoom(roomID uint, updatedRoom models.Room) (*Dto.RoomResponseDto, error) {
	existingRoom, err := service.roomRepo.GetRoomByID(roomID)
	if err != nil {
		return nil, err
	}

	// Update the existing room fields here...
	existingRoom.RoomName = updatedRoom.RoomName
	existingRoom.RoomNumber = updatedRoom.RoomNumber
	existingRoom.RoomAddress = updatedRoom.RoomAddress
	existingRoom.ElectricNumber = updatedRoom.ElectricNumber
	existingRoom.ElectricUserNumber = updatedRoom.ElectricUserNumber
	existingRoom.AmountOfBedRoom = updatedRoom.AmountOfBedRoom
	existingRoom.AmountOfToiletRoom = updatedRoom.AmountOfToiletRoom
	existingRoom.AmountOfLivingRoom = updatedRoom.AmountOfLivingRoom
	existingRoom.SizeSQM = updatedRoom.SizeSQM
	existingRoom.TypeOfView = updatedRoom.TypeOfView
	existingRoom.Remark = updatedRoom.Remark
	existingRoom.StatusOfRoom = updatedRoom.StatusOfRoom
	// ... Update other fields as needed

	// Save the modified room back to the database
	modifiedRoom, err := service.roomRepo.ModifyRoom(*existingRoom)
	if err != nil {
		return nil, err
	}

	// Convert to RoomResponse model
	modifiedRoomResponse := Dto.RoomResponseDto{
		ID:                 modifiedRoom.ID,
		OwnerID:            modifiedRoom.OwnerID,
		FloorID:            modifiedRoom.FloorID,
		RoomName:           modifiedRoom.RoomName,
		RoomNumber:         modifiedRoom.RoomNumber,
		RoomAddress:        modifiedRoom.RoomAddress,
		ElectricNumber:     modifiedRoom.ElectricNumber,
		ElectricUserNumber: modifiedRoom.ElectricUserNumber,
		AmountOfBedRoom:    modifiedRoom.AmountOfBedRoom,
		AmountOfToiletRoom: modifiedRoom.AmountOfToiletRoom,
		AmountOfLivingRoom: modifiedRoom.AmountOfLivingRoom,
		SizeSQM:            modifiedRoom.SizeSQM,
		TypeOfView:         modifiedRoom.TypeOfView,
		Remark:             modifiedRoom.Remark,
		StatusOfRoom:       modifiedRoom.StatusOfRoom,
		IsActive:           modifiedRoom.IsActive,
		CreatedAt:          modifiedRoom.CreatedAt.String(),
		UpdatedAt:          modifiedRoom.UpdatedAt.String(),
	}

	return &modifiedRoomResponse, nil
}
