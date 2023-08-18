package service

import (
	models_Room "PattayaAvenueProperty/models/Room"
	"PattayaAvenueProperty/repository"
	dto "PattayaAvenueProperty/service/dto"
	"encoding/json"
	"fmt"
	"strconv"
)

type RoomService struct {
	roomRepo repository.RoomRepo
}

func NewRoomService(roomRepo repository.RoomRepo) RoomService {
	return RoomService{
		roomRepo: roomRepo,
	}
}

func (service *RoomService) GetAllPlace() ([]dto.PlaceDto, error) {
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
	persons, err := service.roomRepo.GetAllPersons()
	if err != nil {
		return nil, err
	}
	roomPrices, err := service.roomRepo.GetAllRoomPrices()
	if err != nil {
		return nil, err
	}
	personContacts, err := service.roomRepo.GetAllContacts()
	if err != nil {
		return nil, err
	}

	jsonDataContacts, err := json.Marshal(personContacts)
	if err != nil {
		return nil, err
	}
	jsonDataPersons, err := json.Marshal(persons)
	if err != nil {
		return nil, err
	}
	jsonDataRoomPrices, err := json.Marshal(roomPrices)
	if err != nil {
		return nil, err
	}

	var jsonDataContact []map[string]interface{}
	if err := json.Unmarshal([]byte(jsonDataContacts), &jsonDataContact); err != nil {
		fmt.Println("Error:", err)
	}

	var jsonDataPerson []map[string]interface{}
	if err := json.Unmarshal([]byte(jsonDataPersons), &jsonDataPerson); err != nil {
		fmt.Println("Error:", err)
	}

	var jsonDataRoomPrice []map[string]interface{}
	if err := json.Unmarshal([]byte(jsonDataRoomPrices), &jsonDataRoomPrice); err != nil {
		fmt.Println("Error:", err)
	}

	// fmt.Println(jsonDataContact)
	// fmt.Println(jsonDataPerson)

	contactsByPersonID := make(map[float64][]map[string]interface{})

	for _, contact := range jsonDataContact {
		personID := contact["PersonID"].(float64) // Convert to float64
		if _, ok := contactsByPersonID[personID]; !ok {
			contactsByPersonID[personID] = []map[string]interface{}{}
		}
		contactsByPersonID[personID] = append(contactsByPersonID[personID], contact)
	}

	for _, person := range jsonDataPerson {
		personID := person["ID"].(float64) // Convert to float64
		if contactSlice, ok := contactsByPersonID[personID]; ok {
			person["Contacts"] = contactSlice
		}
	}

	roomsByFloorID := make(map[uint][]dto.RoomDto)
	for _, room := range rooms {
		subList := dto.RoomDto{
			RoomID:       room.ID,
			RoomNumber:   room.RoomNumber,
			RoomSize:     room.SizeSQM,
			RoomPrice:    "",
			OwnerName:    "",
			OwnerContact: "",
			StatusOfRoom: room.StatusOfRoom,
		}
		for _, price := range jsonDataRoomPrice {
			if price["UnitType"] == "SHOW" {
				roomIDStr := strconv.FormatUint(uint64(room.ID), 10)                          // Convert room.ID to string
				priceRoomIDStr := strconv.FormatFloat(price["RoomID"].(float64), 'f', -1, 64) // Convert price["RoomID"] to string
				if roomIDStr == priceRoomIDStr {
					subList.RoomPrice = strconv.FormatFloat(price["Amount"].(float64), 'f', -1, 64)
					break
				} else {
					subList.RoomPrice = ""
					break
				}
			}
		}
		for _, person := range jsonDataPerson {
			if room.OwnerID != nil {
				roomOwnerIDFloat := float64(*room.OwnerID)
				if person["ID"] == roomOwnerIDFloat {
					subList.OwnerName = person["FullName"].(string)
					contacts, ok := person["Contacts"].([]map[string]interface{})
					if ok {
						for _, contact := range contacts {
							contactType := contact["Type"].(string)
							subList.OwnerContact = contactType
							break
						}
						break
					}
				} else {
					subList.OwnerName = ""
					subList.OwnerContact = ""
				}
			}
		}
		roomsByFloorID[room.FloorID] = append(roomsByFloorID[room.FloorID], subList)
	}

	var result []dto.PlaceDto
	for _, place := range places {
		var buildingsDto []dto.BuildingDto
		for _, building := range buildings {
			if building.PlaceID == place.ID {
				var floorsDto []dto.FloorDto
				for _, floor := range floors {
					if floor.BuildingID == building.ID {
						if roomsForFloor, ok := roomsByFloorID[floor.ID]; ok {
							floorsDto = append(floorsDto, dto.FloorDto{
								FloorID:     floor.ID,
								FloorNumber: floor.FloorNumber,
								Rooms:       roomsForFloor,
							})
						} else {
							floorsDto = append(floorsDto, dto.FloorDto{
								FloorID:     floor.ID,
								FloorNumber: floor.FloorNumber,
								Rooms:       nil,
							})
						}
					}
				}
				buildingsDto = append(buildingsDto, dto.BuildingDto{
					BuildingID:   building.ID,
					BuildingName: building.BuildingName,
					Floors:       floorsDto,
				})
			}
		}
		result = append(result, dto.PlaceDto{
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

func (service *RoomService) GetRoomByID(roomID uint) (*dto.RoomResponseDto, error) {
	room, err := service.roomRepo.GetRoomByID(roomID)
	if err != nil {
		return nil, err
	}

	// Convert to RoomResponse model
	roomResponse := dto.RoomResponseDto{
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

func (service *RoomService) ModifyRoom(roomID uint, updatedRoom models_Room.Room) (*dto.RoomResponseDto, error) {
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
	modifiedRoomResponse := dto.RoomResponseDto{
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
