package dto

type EditRoomInfoDto struct {
	OwnerID            uint    `json:"owner_id" gorm:"column:owner_id"`
	RoomName           string  `json:"room_name" gorm:"column:room_name"`
	RoomNumber         string  `json:"room_number" gorm:"column:room_number"`
	RoomAddress        string  `json:"room_address" gorm:"column:room_address"`
	ElectricNumber     string  `json:"electric_number" gorm:"column:electric_number"`
	ElectricUserNumber string  `json:"electric_user_number" gorm:"column:electric_user_number"`
	AmountOfBedRoom    *int32  `json:"amount_of_bed_room" gorm:"column:amount_of_bed_room"`
	AmountOfToiletRoom *int32  `json:"amount_of_toilet_room" gorm:"column:amount_of_toilet_room"`
	AmountOfLivingRoom *int32  `json:"amount_of_living_room" gorm:"column:amount_of_living_room"`
	SizeSQM            float32 `json:"size_sqm" gorm:"column:size_sqm"`
	TypeOfView         string  `json:"type_of_view" gorm:"column:type_of_view"`
	Remark             string  `json:"remark" gorm:"column:remark"`
	StatusOfRoom       string  `json:"status_of_room" gorm:"column:status_of_room"`
}

type RoomResponseDto struct {
	ID                 uint    `json:"id"`
	OwnerID            *uint   `json:"owner_id"`
	FloorID            uint    `json:"floor_id"`
	RoomName           *string `json:"room_name"`
	RoomNumber         string  `json:"room_number"`
	RoomAddress        string  `json:"room_address"`
	ElectricNumber     *string `json:"electric_number"`
	ElectricUserNumber *string `json:"electric_user_number"`
	AmountOfBedRoom    *int32  `json:"amount_of_bed_room"`
	AmountOfToiletRoom *int32  `json:"amount_of_toilet_room"`
	AmountOfLivingRoom *int32  `json:"amount_of_living_room"`
	SizeSQM            float32 `json:"size_sqm"`
	TypeOfView         string  `json:"type_of_view"`
	Remark             *string `json:"remark"`
	StatusOfRoom       string  `json:"status_of_room"`
	IsActive           bool    `json:"is_active"`
	CreatedAt          string  `json:"CreatedAt"`
	UpdatedAt          string  `json:"UpdatedAt"`
}
