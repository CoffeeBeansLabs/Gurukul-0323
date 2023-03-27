package parking_lot

import (
	"Gurukul-0323/parking_lot/error_handler"
	"Gurukul-0323/parking_lot/model"
)

type Attendant struct {
	parkingLots []model.IParkingLot
	strategy    model.IParkingLotSelector // fuzzyStrategy, MaxCapacityStrategy
}

func NewAttendant(parkinglots []model.IParkingLot, strategy model.IParkingLotSelector) *Attendant {
	return &Attendant{
		parkingLots: parkinglots,
		strategy:    strategy,
	}
}

type IAttendant interface {
	ParkVehicle(regNumber string) error
	UnparkVehicle(regNumber string) error
}

func (a *Attendant) ParkVehicle(regNumber string) error {
	lot, err := a.strategy.SelectParkingLot(a.parkingLots)
	if err != nil {
		return err
	}
	return lot.ParkVehicle(regNumber)
}

func (a *Attendant) UnparkVehicle(regNumber string) error {
	lot, err := a.fetchParkingLotHavingTheCar(regNumber)
	if err != nil {
		return err
	}
	return lot.ParkVehicle(regNumber)
}

func (a *Attendant) fetchParkingLotHavingSpace() (model.IParkingLot, error) {
	for _, parkingLot := range a.parkingLots {
		if !parkingLot.IsFull() {
			return parkingLot, nil
		}
	}

	return nil, error_handler.ParkingFullError
}

func (a *Attendant) NotifyParkingFull() {
}

func (a *Attendant) NotifyParkingSpaceAvailable() {
}

func (a *Attendant) fetchParkingLotHavingTheCar(regNumber string) (model.IParkingLot, error) {
	for _, parkingLot := range a.parkingLots {
		if parkingLot.IsVehicleParked(regNumber) {
			return parkingLot, nil
		}
	}

	return nil, error_handler.VehicleNotfoundError
}
