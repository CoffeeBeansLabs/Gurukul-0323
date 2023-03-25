package mocks

import (
	"Gurukul-0323/parking_lot/model"
)

// Mock parking lot for testing purposes
type MockStrategy struct {
	ParkingLot    model.IParkingLot
	ErrorToReturn error
}

func (mpl MockStrategy) SelectParkingLot([]model.IParkingLot) (model.IParkingLot, error) {
	return mpl.ParkingLot, mpl.ErrorToReturn
}
