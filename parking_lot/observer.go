package parking_lot

type Observer struct {
}

type IObserver interface {
	NotifyParkingFull()
	NotifyParkingSpaceAvailable()
}
