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
	RoomID   uint   `json:"room_id" gorm:"column:room_id"`
	RoomName string `json:"room_name" gorm:"column:room_name"`
}

type CreateRoomDto struct {
	FloorID  uint   `json:"floor_id"`
	RoomName string `json:"room_name"`
}
