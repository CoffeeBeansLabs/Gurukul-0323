package mocks

type MockObserver struct {
	NotifyParkingFullIsCalledTimes           int
	NotifyParkingSpaceAvailableIsCalledTimes int
}

func (m *MockObserver) NotifyParkingFull() {
	m.NotifyParkingFullIsCalledTimes += 1
}

func (m *MockObserver) NotifyParkingSpaceAvailable() {
	m.NotifyParkingSpaceAvailableIsCalledTimes += 1
}
