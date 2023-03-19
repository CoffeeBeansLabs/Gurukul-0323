package main

import (
	"Gurukul-0323/parking_lot"
	"fmt"
)

func main() {
	owner := parking_lot.Owner{}
	parkingLot := parking_lot.NewParkingLot(5, &owner)
	err := parkingLot.ParkVehicle("abc123")
	if err != nil {
		return
	}
	if parkingLot.IsVehicleParked("abc123") {
		fmt.Println("Vehicle parked successfully")
	}
}
