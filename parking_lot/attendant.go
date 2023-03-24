package parking_lot

type Attendant struct {
	parkingLot ParkingLot
}

func NewAttendant(parkinglot ParkingLot) *Attendant {
	return &Attendant{
		parkingLot: parkinglot,
	}
}

type IAttendant interface {
	ParkVehicle(regNumber string) error
	UnparkVehicle(regNumber string) error
}

func (p *Attendant) ParkVehicle(regNumber string) error {
	if p.parkingLot.IsVehicleParked(regNumber) {
		return VehicleAlreadyParked
	}
	if p.parkingLot.IsFull() {
		p.parkingLot.owner.NotifyParkingFull()
		return ParkingFullError
	}
	p.parkingLot.vehicles[regNumber] = true
	if p.parkingLot.IsFull() {
		p.parkingLot.owner.NotifyParkingFull()
	}
	return nil
}

func (p *Attendant) UnparkVehicle(regNumber string) error {

	if p.parkingLot.IsVehicleParked(regNumber) {
		delete(p.parkingLot.vehicles, regNumber)
		p.parkingLot.owner.NotifyParkingSpaceAvailable()
		return nil
	}

	return VehicleNotfoundError
}
