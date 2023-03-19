package parking_lot

import "fmt"

type Owner struct {
	isNotified bool
}

func NewOwner() *Owner {
	return &Owner{
		isNotified: false,
	}
}

type IOwner interface {
	NotifyParkingFull()
}

func (o *Owner) NotifyParkingFull() {
	if !o.isNotified {
		// send notification to owner that parking lot is full
		fmt.Println("[Owner-notification] : Parking Lot is full !! Please Look for new Lot")
		o.isNotified = true
	}
}
