package model

type IParkingLot interface {
	ParkVehicle(regNumber string) error
	IsVehicleParked(regNumber string) bool
	UnparkVehicle(regNumber string) error
	IsFull() bool
	FetchCapacity() int
	GetAvailableSpace() int
}
