package parking_lot

type attendant struct {
	parkingLot IParkingLot
}

func NewAttendant(parkinglot IParkingLot) *attendant {
	return &attendant{
		parkingLot: parkinglot,
	}
}

type IAttendant interface {
	ParkVehicle(regNumber string) error
	UnparkVehicle(regNumber string) error
}

func (a *attendant) ParkVehicle(regNumber string) error {
	return a.parkingLot.ParkVehicle(regNumber)
}

func (a *attendant) UnparkVehicle(regNumber string) error {
	return a.parkingLot.UnparkVehicle(regNumber)
}

func (a *attendant) NotifyParkingFull() {
}

func (a *attendant) NotifyParkingSpaceAvailable() {
}
