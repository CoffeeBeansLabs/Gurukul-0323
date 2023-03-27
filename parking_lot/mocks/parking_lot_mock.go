package mocks

// Mock parking lot for testing purposes
type MockParkingLot struct {
	IsFullVal          bool
	IsVehicleParkedVal bool
	capacity           int
	availableSpace     int
}

func (mpl *MockParkingLot) IsFull() bool {
	return mpl.IsFullVal
}

func (mpl *MockParkingLot) IsVehicleParked(regNumber string) bool {
	return mpl.IsVehicleParkedVal
}

func (mpl *MockParkingLot) ParkVehicle(regNumber string) error {
	return nil
}
func (mpl *MockParkingLot) UnparkVehicle(regNumber string) error {
	return nil
}

func (mpl *MockParkingLot) FetchCapacity() int {
	return mpl.capacity
}
func (mpl *MockParkingLot) GetAvailableSpace() int {
	return mpl.availableSpace
}
