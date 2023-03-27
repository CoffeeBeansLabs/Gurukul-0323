package strategies

import (
	"Gurukul-0323/parking_lot/error_handler"
	"Gurukul-0323/parking_lot/model"
)

type FirstComeFirstServe struct {
}

func (FirstComeFirstServe) SelectParkingLot(lots []model.IParkingLot) (model.IParkingLot, error) {
	for _, parkingLot := range lots {
		if !parkingLot.IsFull() {
			return parkingLot, nil
		}
	}

	return nil, error_handler.ParkingFullError
}
