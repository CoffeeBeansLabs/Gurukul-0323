package mocks

// Mock parking lot for testing purposes
type MockParkingLot struct {
	IsFullVal          bool
	IsVehicleParkedVal bool
}

func (mpl MockParkingLot) IsFull() bool {
	return mpl.IsFullVal
}

func (mpl MockParkingLot) IsVehicleParked(regNumber string) bool {
	return mpl.IsVehicleParkedVal
}

func (mpl *MockParkingLot) ParkVehicle(regNumber string) error {
	return nil
}
func (mpl *MockParkingLot) UnparkVehicle(regNumber string) error {
	return nil
}
