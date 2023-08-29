package dto

type ContractResponseDto struct {
	ID                uint    `json:"id"`
	RoomID            uint    `json:"room_id"`
	StartContractDate string  `json:"start_contract_date"`
	EndContractDate   string  `json:"end_contract_date"`
	Rental            float32 `json:"rental"`
	Deposit           float32 `json:"deposit"`
	TenantName        string  `json:"tenant_name"`
	RoomNumber        string  `json:"room_number"`
	ContractStatus    string  `json:"contract_status"`
}

type ContractByIDResponseDto struct {
	ID                uint    `json:"id"`
	RoomID            uint    `json:"room_id"`
	StartContractDate string  `json:"start_contract_date"`
	EndContractDate   string  `json:"end_contract_date"`
	Rental            float32 `json:"rental"`
	Deposit           float32 `json:"deposit"`
	TenantName        string  `json:"tenant_name"`
	ContactType       string  `json:"contact_type"`
	ContractStatus    string  `json:"contract_status"`
}

type ContractDetailDto struct {
	ID                     uint    `json:"id"`
	RoomNumber             string  `json:"room_number"`
	RoomAddress            string  `json:"room_address"`
	TenantName             string  `json:"tenant_name"`
	IdentifyNumber         string  `json:"identify_number"`
	StartContractDate      string  `json:"start_contract_date"`
	EndContractDate        string  `json:"end_contract_date"`
	Rental                 float32 `json:"rental"`
	Deposit                float32 `json:"deposit"`
	CheckInDate            string  `json:"check_in_date"`
	CheckOutDate           *string `json:"check_out_date"`
	CheckInElectricNumber  *int    `json:"check_in_electric_number"`
	CheckInWaterNumber     *int    `json:"check_in_water_number"`
	CheckOutElectricNumber *int    `json:"check_out_electric_number"`
	CheckOutWaterNumber    *int    `json:"check_out_water_number"`
	IsClosed               bool    `json:"is_closed"`
	UpdatedAt              string  `json:"updated_at"`
}

type CreateRoomContractDto struct {
	RoomID                uint    `json:"room_id"`
	PersonID              uint    `json:"person_id"`
	PersonContractType    string  `json:"person_contract_type"`
	StartContractDate     string  `json:"start_contract_date"`
	EndContractDate       string  `json:"end_contract_date"`
	Rental                float32 `json:"rental"`
	Deposit               float32 `json:"deposit"`
	CheckInDate           string  `json:"check_in_date"`
	CheckInWaterNumber    int     `json:"check_in_water_number"`
	CheckInElectricNumber int     `json:"check_in_electric_number"`
}

type CloseRoomContractDto struct {
	RoomContractID         uint   `json:"room_contract_id"`
	CheckOutDate           string `json:"check_out_date"`
	CheckOutWaterNumber    int    `json:"check_out_water_number"`
	CheckOutElectricNumber int    `json:"check_out_electric_number"`
}
