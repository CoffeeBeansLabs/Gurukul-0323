package model

type IParkingLotSelector interface {
	SelectParkingLot([]IParkingLot) (IParkingLot, error)
}
