package service

import (
	"PattayaAvenueProperty/repository"
	dto "PattayaAvenueProperty/service/dto"
)

type ContractService struct {
	contractRepo repository.ContractRepo
	personRepo   repository.PersonRepo
	roomRepo     repository.RoomRepo
}

func NewContractService(contractRepo repository.ContractRepo, personRepo repository.PersonRepo, roomRepo repository.RoomRepo) ContractService {
	return ContractService{
		contractRepo: contractRepo,
		personRepo:   personRepo,
	}
}

func (service *ContractService) GetAllContract() ([]dto.ContractResponseDto, error) {
	roomContract, err := service.contractRepo.GetAllRoomContract()
	if err != nil {
		return nil, err
	}
	personContract, err := service.contractRepo.GetAllPersonContract()
	if err != nil {
		return nil, err
	}
	person, err := service.personRepo.FindAll()
	if err != nil {
		return nil, err
	}
	rooms, err := service.roomRepo.GetAllRoom()
	if err != nil {
		return nil, err
	}

	var result []dto.ContractResponseDto
	for _, contract := range roomContract {
		var tenantName string
		var roomNumber string
		startContractDateString := contract.StartContractDate.Format("2006-01-02 15:04:05")
		endContractDateString := contract.EndContractDate.Format("2006-01-02 15:04:05")
		for _, personContract := range personContract {
			if contract.ID == personContract.RoomContractID {
				for _, person := range person {
					if person.ID == personContract.PersonID {
						tenantName = person.FullName
					}
				}
			}
		}
		for _, room := range rooms {
			if contract.RoomID == room.ID {
				roomNumber = room.RoomNumber
			}
		}

		result = append(result, dto.ContractResponseDto{
			ID:                contract.ID,
			RoomID:            contract.RoomID,
			StartContractDate: startContractDateString,
			EndContractDate:   endContractDateString,
			Rental:            contract.Rental,
			Deposit:           contract.Deposit,
			TenantName:        tenantName,
			RoomNumber:        roomNumber,
			ContractStatus:    contract.IsClosed,
		})
	}
	return result, nil
}

func (service *ContractService) GetContractByRoomID(roomID uint) ([]dto.ContractByIDResponseDto, error) {
	roomContract, err := service.contractRepo.GetRoomContractByRoomID(roomID)
	if err != nil {
		return nil, err
	}
	personContract, err := service.contractRepo.GetPersonContractByPersonContractID(roomID)
	if err != nil {
		return nil, err
	}
	person, err := service.personRepo.FindAll()
	if err != nil {
		return nil, err
	}
	contact, err := service.personRepo.GetAllContact()
	if err != nil {
		return nil, err
	}

	var result []dto.ContractByIDResponseDto
	for _, contract := range roomContract {
		var tenantName string
		var contactType string
		startContractDateString := contract.StartContractDate.Format("2006-01-02 15:04:05")
		endContractDateString := contract.EndContractDate.Format("2006-01-02 15:04:05")
		for _, personContract := range personContract {
			if contract.ID == personContract.RoomContractID {
				for _, person := range person {
					for _, contact := range contact {
						if person.ID == contact.PersonID {
							tenantName = person.FullName
							contactType = contact.Type
						}
					}
				}
			}
		}

		var contractStatus string
		if !contract.IsClosed {
			contractStatus = "active"
		} else {
			contractStatus = "inactive"
		}

		result = append(result, dto.ContractByIDResponseDto{
			ID:                contract.ID,
			RoomID:            contract.RoomID,
			StartContractDate: startContractDateString,
			EndContractDate:   endContractDateString,
			Rental:            contract.Rental,
			Deposit:           contract.Deposit,
			TenantName:        tenantName,
			ContactType:       contactType,
			ContractStatus:    contractStatus,
		})
	}
	return result, nil
}

func (service *ContractService) GetRoomContractByID(id uint) (dto.ContractDetailDto, error) {
	roomContract, err := service.contractRepo.GetRoomContractByID(id)
	if err != nil {
		return dto.ContractDetailDto{}, err
	}
	rooms, err := service.roomRepo.GetAllRoom()
	if err != nil {
		return dto.ContractDetailDto{}, err
	}
	person, err := service.personRepo.FindAll()
	if err != nil {
		return dto.ContractDetailDto{}, err
	}
	contact, err := service.personRepo.GetAllContact()
	if err != nil {
		return dto.ContractDetailDto{}, err
	}
	personContract, err := service.contractRepo.GetAllPersonContract()
	if err != nil {
		return dto.ContractDetailDto{}, err
	}

	var tenantName string
	for _, personContract := range personContract {
		if roomContract.ID == personContract.RoomContractID {
			for _, person := range person {
				for _, contact := range contact {
					if person.ID == contact.PersonID {
						tenantName = person.FullName
					}
				}
			}
		}
	}
	var RoomNumber string
	var RoomAddress string
	for _, room := range rooms {
		if roomContract.RoomID == room.ID {
			RoomNumber = room.RoomNumber
			RoomAddress = room.RoomAddress
		}
	}

	result := dto.ContractDetailDto{
		ID:                     roomContract.ID,
		RoomNumber:             RoomNumber,
		RoomAddress:            RoomAddress,
		TenantName:             tenantName,
		StartContractDate:      roomContract.StartContractDate.Format("2006-01-02 15:04:05"),
		EndContractDate:        roomContract.EndContractDate.Format("2006-01-02 15:04:05"),
		Rental:                 roomContract.Rental,
		Deposit:                roomContract.Deposit,
		CheckInDate:            roomContract.CheckInDate.Format("2006-01-02 15:04:05"),
		CheckOutDate:           roomContract.CheckOutDate.Format("2006-01-02 15:04:05"),
		CheckInElectricNumber:  roomContract.CheckInElectricNumber,
		CheckInWaterNumber:     roomContract.CheckInWaterNumber,
		CheckOutElectricNumber: roomContract.CheckOutElectricNumber,
		CheckOutWaterNumber:    roomContract.CheckOutWaterNumber,
		IsClosed:               roomContract.IsClosed,
		UpdatedAt:              roomContract.UpdatedAt.Format("2006-01-02 15:04:05"),
	}

	return result, nil
}
