package parking_lot

type attendant struct {
	parkingLots []IParkingLot
}

func NewAttendant(parkinglots []IParkingLot) *attendant {
	return &attendant{
		parkingLots: parkinglots,
	}
}

type IAttendant interface {
	ParkVehicle(regNumber string) error
	UnparkVehicle(regNumber string) error
}

func (a *attendant) ParkVehicle(regNumber string) error {
	lot, err := a.fetchParkingLotHavingSpace()
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

func (a *attendant) fetchParkingLotHavingSpace() (IParkingLot, error) {
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

func (a *attendant) fetchParkingLotHavingTheCar(regNumber string) (IParkingLot, error) {
	for _, parkingLot := range a.parkingLots {
		if parkingLot.IsVehicleParked(regNumber) {
			return parkingLot, nil
		}
	}

	return nil, VehicleNotfoundError
}
