package parking_lot

import "errors"

var (
	ParkingFullError = errors.New("parking lot is full")
)

type ParkingLot struct {
	capacity int
	vehicles map[string]bool
}

func NewParkingLot(capacity int) *ParkingLot {
	return &ParkingLot{
		capacity: capacity,
		vehicles: make(map[string]bool),
	}
}

func (p *ParkingLot) ParkVehicle(regNumber string) error {
	if p.IsFull() {
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
