package strategies

import (
	"Gurukul-0323/parking_lot"
	"Gurukul-0323/parking_lot/model"
)

type fifoStrategy struct {
}

func (fifoStrategy) selectParkingLot(lots []model.IParkingLot) (model.IParkingLot, error) {
	for _, parkingLot := range lots {
		if !parkingLot.IsFull() {
			return parkingLot, nil
		}
	}

	return nil, parking_lot.ParkingFullError
}
