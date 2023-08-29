package service

import (
	models_Document "PattayaAvenueProperty/models/Document"
	models_Room "PattayaAvenueProperty/models/Room"
	"PattayaAvenueProperty/repository"
	dto "PattayaAvenueProperty/service/dto"
	"encoding/json"
	"fmt"
	"strconv"
)

type RoomService struct {
	roomRepo   repository.RoomRepo
	personRepo repository.PersonRepo
}

func NewRoomService(roomRepo repository.RoomRepo, personRepo repository.PersonRepo) RoomService {
	return RoomService{
		roomRepo:   roomRepo,
		personRepo: personRepo,
	}
}

// Helper function to convert interface{} to *string
func getStringPtr(value interface{}) *string {
	if str, ok := value.(string); ok {
		return &str
	}
	return nil
}

// Helper function to convert interface{} to *int32
func getInt32Ptr(value interface{}) *int32 {
	if intValue, ok := value.(float64); ok {
		intValue32 := int32(intValue)
		return &intValue32
	}
	return nil
}

// Helper function to convert interface{} to *float32
func getFloat32Ptr(value interface{}) *float32 {
	if floatValue, ok := value.(float64); ok {
		floatValue32 := float32(floatValue)
		return &floatValue32
	}
	return nil
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
			if price["Type"] == "SHOW" {
				roomIDStr := strconv.FormatUint(uint64(room.ID), 10)                          // Convert room.ID to string
				priceRoomIDStr := strconv.FormatFloat(price["RoomID"].(float64), 'f', -1, 64) // Convert price["RoomID"] to string
				if roomIDStr == priceRoomIDStr {
					subList.RoomPrice = strconv.FormatFloat(price["Amount"].(float64), 'f', -1, 64)
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

func (service *RoomService) CreateRoom(floorID uint, roomNumber string) error {
	roomModel := models_Room.Room{
		RoomNumber: roomNumber,
		FloorID:    floorID,
		IsActive:   true,
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
	price, err := service.roomRepo.GetRoomPriceByRoomID(roomID)
	if err != nil {
		return nil, err
	}
	picture, err := service.roomRepo.GetRoomPictureByRoomID(roomID)
	if err != nil {
		return nil, err
	}
	roomDocument, err := service.roomRepo.GetRoomDocumentByRoomID(roomID)
	if err != nil {
		return nil, err
	}

	jsonDataRoomPrices, err := json.Marshal(price)
	if err != nil {
		return nil, err
	}
	jsonDataRoomPictures, err := json.Marshal(picture)
	if err != nil {
		return nil, err
	}

	var jsonDataRoomPicture []map[string]interface{}
	if err := json.Unmarshal([]byte(jsonDataRoomPictures), &jsonDataRoomPicture); err != nil {
		fmt.Println("Error:", err)
	}
	var jsonDataRoomPrice []map[string]interface{}
	if err := json.Unmarshal([]byte(jsonDataRoomPrices), &jsonDataRoomPrice); err != nil {
		fmt.Println("Error:", err)
	}

	var ownerName *string

	if room.OwnerID != nil {
		person, err := service.personRepo.FindPersonById(*room.OwnerID)
		if err != nil {
			return nil, err
		}
		jsonDataPerson, err := json.Marshal(person)
		if err != nil {
			return nil, err
		}
		var jsonDataPersons map[string]interface{}
		if err := json.Unmarshal(jsonDataPerson, &jsonDataPersons); err != nil {
			fmt.Println("Error:", err)
		}
		if fullName, ok := jsonDataPersons["FullName"].(string); ok {
			ownerName = &fullName
		} else {
			ownerName = nil
		}
	}

	// Convert to RoomPictureResponse model
	var roomPictures []dto.RoomPictureResponseDto
	for _, picture := range jsonDataRoomPicture {
		roomPicture := dto.RoomPictureResponseDto{
			ID:             uint(picture["ID"].(float64)), // Type assertion to float64
			RoomPictureUrl: picture["RoomPictureUrl"].(string),
		}
		roomPictures = append(roomPictures, roomPicture)
	}

	// Convert to RoomPriceResponse model
	var roomPrices []dto.RoomPriceResponseDto
	for _, price := range jsonDataRoomPrice {
		roomPrice := dto.RoomPriceResponseDto{
			ID:              uint(price["ID"].(float64)), // Type assertion to float64
			Amount:          float32(price["Amount"].(float64)),
			UnitType:        getStringPtr(price["UnitType"]),
			MinDuration:     getInt32Ptr(price["MinDuration"]),
			MaxDuration:     getInt32Ptr(price["MaxDuration"]),
			Type:            price["Type"].(string),
			DepositAmount:   getFloat32Ptr(price["DepositAmount"]),
			DepositUnitType: getStringPtr(price["DepositUnitType"]),
		}
		roomPrices = append(roomPrices, roomPrice)
	}

	// Convert to RoomDocumentResponse model
	var roomDocuments []dto.RoomDocumentResponseDto
	for _, document := range roomDocument {
		roomDocument := dto.RoomDocumentResponseDto{
			ID:           document.ID,
			RoomDocument: document.DocumentUrl,
		}
		roomDocuments = append(roomDocuments, roomDocument)
	}

	// Convert to RoomResponse model
	// ownerName := jsonDataPersons["FullName"].(string)
	roomResponse := dto.RoomResponseDto{
		ID:                 room.ID,
		OwnerID:            room.OwnerID,
		OwnerName:          nil,
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
		RoomPrices:         roomPrices,
		RoomPictures:       roomPictures,
		RoomDocuments:      roomDocuments,
	}
	roomResponse.OwnerName = ownerName

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
	existingRoom.OwnerID = updatedRoom.OwnerID
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
	}

	return &modifiedRoomResponse, nil
}

func (service *RoomService) ModifyPlace(placeID uint, placeName string) error {
	existingPlace, err := service.roomRepo.GetPlaceByID(placeID)
	if err != nil {
		return err
	}

	// Update the existing place fields here...
	existingPlace.PlaceName = placeName
	// ... Update other fields as needed

	// Save the modified place back to the database
	_, err = service.roomRepo.EditPlace(*existingPlace)
	if err != nil {
		return err
	}

	return nil
}

// modify building
func (service *RoomService) ModifyBuilding(buildingID uint, buildingName string) error {
	existingBuilding, err := service.roomRepo.GetBuildingByID(buildingID)
	if err != nil {
		return err
	}

	// Update the existing building fields here...
	existingBuilding.BuildingName = buildingName
	// ... Update other fields as needed

	// Save the modified building back to the database
	_, err = service.roomRepo.EditBuilding(*existingBuilding)
	if err != nil {
		return err
	}

	return nil
}

// create RoomPrice and delete RoomPrice by change is_active to 0
func (service *RoomService) CreateRoomPrice(data dto.TakeRoomPriceDataDto) (*models_Room.RoomPrice, error) {
	roomPriceModel := models_Room.RoomPrice{
		RoomID:          data.RoomID,
		Amount:          data.Amount,
		UnitType:        data.UnitType,
		MinDuration:     data.MinDuration,
		MaxDuration:     data.MaxDuration,
		Type:            data.Type,
		DepositUnitType: data.DepositUnitType,
		DepositAmount:   data.DepositAmount,
		IsActive:        true,
	}

	roomPrice, err := service.roomRepo.CreateRoomPrice(roomPriceModel)

	if err != nil {
		return &models_Room.RoomPrice{}, err
	}
	return roomPrice, nil
}

func (service *RoomService) DeleteRoomPrice(roomPriceID uint) error {
	err := service.roomRepo.DeleteRoomPrice(roomPriceID)
	if err != nil {
		return err
	}
	return nil
}

// create RoomPicture and delete RoomPicture by change is_active to 0
func (service *RoomService) CreateRoomPicture(roomPictureData dto.TakeRoomPictureDataDto) (dto.RoomPictureResponseDto, error) {
	roomPictureModel := models_Room.RoomPicture{
		RoomID:         roomPictureData.RoomID,
		RoomPictureUrl: roomPictureData.RoomPictureUrl,
		IsActive:       true,
	}

	roomPicture, err := service.roomRepo.CreateRoomPicture(roomPictureModel)
	if err != nil {
		return dto.RoomPictureResponseDto{}, err
	}
	return dto.RoomPictureResponseDto{
		ID:             roomPicture.ID,
		RoomPictureUrl: roomPicture.RoomPictureUrl,
	}, nil
}

func (service *RoomService) DeleteRoomPicture(roomPictureID uint) error {
	err := service.roomRepo.DeleteRoomPicture(roomPictureID)
	if err != nil {
		return err
	}
	return nil
}

// create RoomDocument and delete RoomDocument by change is_active to 0
func (service *RoomService) CreateRoomDocument(TakeRoomDocumentData dto.TakeRoomDocumentDataDto) (dto.RoomDocumentResponseDto, error) {
	roomDocumentModel := models_Document.RoomDocument{
		RoomID:      TakeRoomDocumentData.RoomID,
		DocumentUrl: TakeRoomDocumentData.RoomDocument,
		IsActive:    true,
	}

	roomDocument, err := service.roomRepo.CreateRoomDocument(roomDocumentModel)
	if err != nil {
		return dto.RoomDocumentResponseDto{}, err
	}
	return dto.RoomDocumentResponseDto{
		ID:           roomDocument.ID,
		RoomDocument: roomDocument.DocumentUrl,
	}, nil
}

func (service *RoomService) DeleteRoomDocument(roomDocumentPictureID uint) error {
	err := service.roomRepo.DeleteRoomDocument(roomDocumentPictureID)
	if err != nil {
		return err
	}
	return nil
}

// get all room number and room id and room address
func (service *RoomService) GetAllRoomName() ([]dto.RoomNameDto, error) {
	rooms, err := service.roomRepo.GetAllRoom()
	if err != nil {
		return nil, err
	}
	var result []dto.RoomNameDto
	for _, room := range rooms {
		result = append(result, dto.RoomNameDto{
			RoomID:      room.ID,
			RoomNumber:  room.RoomNumber,
			RoomAddress: room.RoomAddress,
		})
	}
	return result, nil
}
