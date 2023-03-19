package parking_lot

import "errors"

var (
	ParkingFullError     = errors.New("parking lot is full")
	VehicleNotfoundError = errors.New("vehicle not found")
	VehicleAlreadyParked = errors.New("vehicle already parked")
)

type ParkingLot struct {
	capacity int
	vehicles map[string]bool
	owner    *Owner
}

func NewParkingLot(capacity int, owner *Owner) *ParkingLot {
	return &ParkingLot{
		capacity: capacity,
		vehicles: make(map[string]bool),
		owner:    owner,
	}
}

func (p *ParkingLot) ParkVehicle(regNumber string) error {
	if p.IsVehicleParked(regNumber) {
		return VehicleAlreadyParked
	}
	if p.IsFull() {
		p.owner.NotifyParkingFull()
		return ParkingFullError
	}
	p.vehicles[regNumber] = true
	return nil
}

func (p *ParkingLot) IsVehicleParked(regNumber string) bool {
	return p.vehicles[regNumber]
}

func (p *ParkingLot) IsFull() bool {
	return len(p.vehicles) >= p.capacity
}

func (p *ParkingLot) UnparkVehicle(regNumber string) error {

	if p.IsVehicleParked(regNumber) {
		delete(p.vehicles, regNumber)
		return nil
	}
	return VehicleNotfoundError
}
