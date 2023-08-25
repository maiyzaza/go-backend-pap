package dto

type ContractResponseDto struct {
	ID                uint    `json:"ID"`
	RoomID            uint    `json:"RoomID"`
	StartContractDate string  `json:"StartContractDate"`
	EndContractDate   string  `json:"EndContractDate"`
	Rental            float32 `json:"Rental"`
	Deposit           float32 `json:"Deposit"`
	TenantName        string  `json:"TenantName"`
	RoomNumber        string  `json:"RoomNumber"`
	ContractStatus    bool    `json:"ContractStatus"`
}

type ContractByIDResponseDto struct {
	ID                uint    `json:"ID"`
	RoomID            uint    `json:"RoomID"`
	StartContractDate string  `json:"StartContractDate"`
	EndContractDate   string  `json:"EndContractDate"`
	Rental            float32 `json:"Rental"`
	Deposit           float32 `json:"Deposit"`
	TenantName        string  `json:"TenantName"`
	ContactType       string  `json:"ContactType"`
	ContractStatus    string  `json:"ContractStatus"`
}

type ContractDetailDto struct {
	ID                     uint    `json:"ID"`
	RoomNumber             string  `json:"RoomNumber"`
	RoomAddress            string  `json:"RoomAddress"`
	TenantName             string  `json:"TenantName"`
	StartContractDate      string  `json:"StartContractDate"`
	EndContractDate        string  `json:"EndContractDate"`
	Rental                 float32 `json:"Rental"`
	Deposit                float32 `json:"Deposit"`
	CheckInDate            string  `json:"CheckInDate"`
	CheckOutDate           string  `json:"CheckOutDate"`
	CheckInElectricNumber  *int    `json:"CheckInElectricNumber"`
	CheckInWaterNumber     *int    `json:"CheckInWaterNumber"`
	CheckOutElectricNumber *int    `json:"CheckOutElectricNumber"`
	CheckOutWaterNumber    *int    `json:"CheckOutWaterNumber"`
	IsClosed               bool    `json:"IsClosed"`
	UpdatedAt              string  `json:"UpdatedAt"`
}
