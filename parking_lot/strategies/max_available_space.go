package strategies

import (
	"Gurukul-0323/parking_lot/error_handler"
	"Gurukul-0323/parking_lot/model"
)

type MaxAvailableSpace struct {
}

func (m MaxAvailableSpace) SelectParkingLot(lots []model.IParkingLot) (model.IParkingLot, error) {
	max := 0
	var maxAvailableSpaceParkingLot model.IParkingLot
	for _, parkingLot := range lots {
		if !parkingLot.IsFull() && max < parkingLot.GetAvailableSpace() {
			max = parkingLot.GetAvailableSpace()
			maxAvailableSpaceParkingLot = parkingLot
		}
	}
	if max == 0 {
		return nil, error_handler.ParkingFullError
	}
	return maxAvailableSpaceParkingLot, nil
}
