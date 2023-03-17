package main

import (
	"Gurukul-0323/parking_lot"
	"fmt"
)

func main() {
	parkingLot := parking_lot.NewParkingLot(5)
	err := parkingLot.ParkVehicle("abc123")
	if err != nil {
		return
	}
	if parkingLot.IsVehicleParked("abc123") {
		fmt.Println("Vehicle parked successfully")
	}
}
