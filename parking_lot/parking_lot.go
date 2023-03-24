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
	owner    IOwner
	observer []IObserver
}

func NewParkingLot(capacity int, observers []IObserver) *ParkingLot {
	return &ParkingLot{
		capacity: capacity,
		vehicles: make(map[string]bool),
		//owner:    owner,
		observer: observers,
	}
}

type IParkingLot interface {
	ParkVehicle(regNumber string) error
	IsVehicleParked(regNumber string) bool
	UnparkVehicle(regNumber string) error
}

func (p *ParkingLot) ParkVehicle(regNumber string) error {
	if p.IsVehicleParked(regNumber) {
		return VehicleAlreadyParked
	}
	if p.IsFull() {
		//p.owner.NotifyParkingFull()
		p.notifyObservers()
		return ParkingFullError
	}
	p.vehicles[regNumber] = true
	if p.IsFull() {
		//p.owner.NotifyParkingFull()
		p.notifyObservers()
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
		p.owner.NotifyParkingSpaceAvailable()
		return nil
	}

	return VehicleNotfoundError
}

//func (p *ParkingLot) addObserver(observer Observer) []Observer {
//	p.observer = append(p.observer, observer)
//	return p.observer
//}

func (p *ParkingLot) notifyObservers() {
	for _, observer := range p.observer {
		observer.NotifyParkingFull()
	}
}

type IOwner interface {
	NotifyParkingFull()
	NotifyParkingSpaceAvailable()
}
