package parking_lot

type Observer struct {
}

type IObserver interface {
	NotifyParkingFull()
	NotifyParkingSpaceAvailable()
}

func (o *Observer) NotifyParkingFull() {
}

func (o *Observer) NotifyParkingSpaceAvailable() {
}