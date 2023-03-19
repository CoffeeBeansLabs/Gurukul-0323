package main

import (
	"Gurukul-0323/parking_lot"
	"fmt"
)

func main() {

	owner := parking_lot.NewOwner()
	parkingLot := parking_lot.NewParkingLot(3, owner)

	go owner.ListenForNotifications()

	firstVehicleParkingErr := parkingLot.ParkVehicle("ABC-123")
	if firstVehicleParkingErr != nil {
		fmt.Println(firstVehicleParkingErr)
	}

	SecondVehicleParkingErr := parkingLot.ParkVehicle("DEF-456")
	if SecondVehicleParkingErr != nil {
		fmt.Println(SecondVehicleParkingErr)
	}

	//ThirdVehicleParkingErr := parkingLot.ParkVehicle("ABC-123")
	//if ThirdVehicleParkingErr != nil {
	//	fmt.Println(ThirdVehicleParkingErr)
	//}
}
