package dto

type PlaceDto struct {
	PlaceID   uint          `json:"place_id" gorm:"column:place_id"`
	PlaceName string        `json:"place_name" gorm:"column:place_name"`
	Buildings []BuildingDto `json:"Buildings" gorm:"column:building"`
}

type BuildingDto struct {
	BuildingID   uint       `json:"building_id" gorm:"column:building_id"`
	BuildingName string     `json:"building_name" gorm:"column:building_name"`
	Floors       []FloorDto `json:"floors" gorm:"column:floor"`
}

type CreateBuildingDto struct {
	PlaceID      uint   `json:"place_id"`
	BuildingName string `json:"building_name"`
}

type FloorDto struct {
	FloorID     uint      `json:"floor_id" gorm:"column:floor_id"`
	FloorNumber string    `json:"floor_number" gorm:"column:floor_number"`
	Rooms       []RoomDto `json:"rooms" gorm:"column:room"`
}

type CreateFloorDto struct {
	BuildingID  uint   `json:"building_id"`
	FloorNumber string `json:"floor_number"`
}

type RoomDto struct {
	RoomID       uint    `json:"room_id" gorm:"column:room_id"`
	RoomNumber   string  `json:"room_number" gorm:"column:room_number"`
	RoomSize     float32 `json:"room_size" gorm:"column:room_size"`
	RoomPrice    string  `json:"room_price" gorm:"column:room_price"`
	OwnerName    string  `json:"owner_name" gorm:"column:owner_name"`
	OwnerContact string  `json:"owner_contacts" gorm:"column:owner_contacts"`
	StatusOfRoom string  `json:"status_of_room" gorm:"column:status_of_room"`
}

type CreateRoomDto struct {
	FloorID      uint    `json:"floor_id"`
	RoomNumber   string  `json:"room_number"`
	RoomAddress  string  `json:"room_address"`
	SizeSQM      float32 `json:"size_sqm"`
	OwnerID      uint    `json:"owner_id"`
	StatusOfRoom string  `json:"status_of_room"`
}

// all room name, room id, room address
type RoomNameDto struct {
	RoomID      uint   `json:"room_id" gorm:"column:room_id"`
	RoomNumber  string `json:"room_number" gorm:"column:room_number"`
	RoomAddress string `json:"room_address" gorm:"column:room_address"`
}
