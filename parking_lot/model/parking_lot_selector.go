package model

type IParkingLotSelector interface {
	SelectParkingLot(lots []IParkingLot) (IParkingLot, error)
}
