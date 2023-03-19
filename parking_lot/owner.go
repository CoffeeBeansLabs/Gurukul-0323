package parking_lot

import "fmt"

type Owner struct {
	Notifications chan bool
}

func NewOwner() *Owner {
	return &Owner{
		Notifications: make(chan bool),
	}
}

func (o *Owner) NotifyParkingFull() {
	o.Notifications <- true
}

func (o *Owner) ListenForNotifications() {
	for {
		<-o.Notifications
		fmt.Println("[OWNER-NOTIFICATION] :Parking lot is full. Please park at the nearest available lot.")
	}
}
