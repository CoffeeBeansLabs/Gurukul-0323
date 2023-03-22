package parking_lot

import "fmt"

type owner struct {
	isNotified bool
}

func NewOwner() *owner {
	return &owner{
		isNotified: false,
	}
}

func (o *owner) NotifyParkingFull() {
	if !o.isNotified {
		// send notification to owner that parking lot is full
		fmt.Println("[Owner-notification] : Parking Lot is full !! Please Look for new Lot")
		o.isNotified = true
	}
}
