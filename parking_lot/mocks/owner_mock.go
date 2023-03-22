package mocks

type MockOwner struct {
	NotifyParkingFullIsCalledTimes           int
	NotifyParkingSpaceAvailableIsCalledTimes int
}

func (m *MockOwner) NotifyParkingFull() {
	m.NotifyParkingFullIsCalledTimes += 1
}

func (m *MockOwner) NotifyParkingSpaceAvailable() {
	m.NotifyParkingSpaceAvailableIsCalledTimes += 1
}
