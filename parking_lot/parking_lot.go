package parking_lot

type ParkingLot struct {
	capacity int
	vehicles map[string]bool
	observer []IObserver
}

func NewParkingLot(capacity int, observers []IObserver) *ParkingLot {
	return &ParkingLot{
		capacity: capacity,
		vehicles: make(map[string]bool),
		observer: observers,
	}
}

func (p *ParkingLot) ParkVehicle(regNumber string) error {
	if p.IsVehicleParked(regNumber) {
		return VehicleAlreadyParked
	}
	if p.isFull() {
		p.notifyObserversIfParkingLotIsFull()
		return ParkingFullError
	}
	p.vehicles[regNumber] = true
	if p.isFull() {
		p.notifyObserversIfParkingLotIsFull()
	}
	return nil
}

func (p *ParkingLot) IsVehicleParked(regNumber string) bool {
	return p.vehicles[regNumber]
}

func (p *ParkingLot) isFull() bool {
	return len(p.vehicles) >= p.capacity
}

func (p *ParkingLot) UnparkVehicle(regNumber string) error {

	if p.IsVehicleParked(regNumber) {
		delete(p.vehicles, regNumber)
		p.notifyObserversIfSpaceAvailable()
		return nil
	}

	return VehicleNotfoundError
}

//func (p *ParkingLot) addObserver(observer Observer) []Observer {
//	p.observer = append(p.observer, observer)
//	return p.observer
//}

func (p *ParkingLot) notifyObserversIfParkingLotIsFull() {
	for _, observer := range p.observer {
		observer.NotifyParkingFull()
	}
}

func (p *ParkingLot) notifyObserversIfSpaceAvailable() {
	for _, observer := range p.observer {
		observer.NotifyParkingFull()
	}
}
