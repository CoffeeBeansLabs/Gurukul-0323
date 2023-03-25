package parking_lot

import "Gurukul-0323/parking_lot/model"

type attendant struct {
	parkingLots []model.IParkingLot
	strategy    model.IParkingLotSelector // fuzzyStrategy, MaxCapacityStrategy
}

func NewAttendant(parkinglots []model.IParkingLot, strategy model.IParkingLotSelector) *attendant {
	return &attendant{
		parkingLots: parkinglots,
		strategy:    strategy,
	}
}

type IAttendant interface {
	ParkVehicle(regNumber string) error
	UnparkVehicle(regNumber string) error
}

func (a *attendant) ParkVehicle(regNumber string) error {
	lot, err := a.strategy.SelectParkingLot(a.parkingLots)
	if err != nil {
		return err
	}
	return lot.ParkVehicle(regNumber)
}

func (a *attendant) UnparkVehicle(regNumber string) error {
	lot, err := a.fetchParkingLotHavingTheCar(regNumber)
	if err != nil {
		return err
	}
	return lot.ParkVehicle(regNumber)
}

func (a *attendant) fetchParkingLotHavingSpace() (model.IParkingLot, error) {
	for _, parkingLot := range a.parkingLots {
		if !parkingLot.IsFull() {
			return parkingLot, nil
		}
	}

	return nil, ParkingFullError
}

func (a *attendant) NotifyParkingFull() {
}

func (a *attendant) NotifyParkingSpaceAvailable() {
}

func (a *attendant) fetchParkingLotHavingTheCar(regNumber string) (model.IParkingLot, error) {
	for _, parkingLot := range a.parkingLots {
		if parkingLot.IsVehicleParked(regNumber) {
			return parkingLot, nil
		}
	}

	return nil, VehicleNotfoundError
}
