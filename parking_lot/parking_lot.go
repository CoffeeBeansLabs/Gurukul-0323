package parking_lot

import "Gurukul-0323/parking_lot/error_handler"

type ParkingLot struct {
	capacity      int
	vehicles      map[string]bool
	observer      []IObserver
	occupiedSlots int
}

func NewParkingLot(capacity int, observers []IObserver) *ParkingLot {
	return &ParkingLot{
		capacity:      capacity,
		vehicles:      make(map[string]bool),
		observer:      observers,
		occupiedSlots: 0,
	}
}

func (p *ParkingLot) ParkVehicle(regNumber string) error {
	if p.IsVehicleParked(regNumber) {
		return error_handler.VehicleAlreadyParked
	}
	if p.IsFull() {
		p.notifyObserversIfParkingLotIsFull()
		return error_handler.ParkingFullError
	}
	p.vehicles[regNumber] = true
	p.occupiedSlots++
	if p.IsFull() {
		p.notifyObserversIfParkingLotIsFull()
	}
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
		p.notifyObserversIfSpaceAvailable()
		return nil
	}

	return error_handler.VehicleNotfoundError
}

func (p *ParkingLot) notifyObserversIfParkingLotIsFull() {
	for _, observer := range p.observer {
		observer.NotifyParkingFull()
	}
}
func (p *ParkingLot) FetchCapacity() int {
	return p.capacity
}

func (p *ParkingLot) notifyObserversIfSpaceAvailable() {
	for _, observer := range p.observer {
		observer.NotifyParkingFull()
	}
}

func (p *ParkingLot) GetAvailableSpace() int {
	return p.capacity - p.occupiedSlots
}
