package strategies

import (
	"Gurukul-0323/parking_lot/error_handler"
	"Gurukul-0323/parking_lot/model"
)

type MaxCapacity struct {
}

func (m MaxCapacity) SelectParkingLot(lots []model.IParkingLot) (model.IParkingLot, error) {
	max := 0
	var maxCapacityParkingLot model.IParkingLot
	for _, parkingLot := range lots {
		if !parkingLot.IsFull() && max < parkingLot.FetchCapacity() {
			max = parkingLot.FetchCapacity()
			maxCapacityParkingLot = parkingLot
		}
	}
	if max == 0 {
		return nil, error_handler.ParkingFullError
	}
	return maxCapacityParkingLot, nil
}
