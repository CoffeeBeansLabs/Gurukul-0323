package mocks

type MockAttendant struct {
	NotifyParkingFullIsCalledTimes           int
	NotifyParkingSpaceAvailableIsCalledTimes int
}

func (m *MockAttendant) NotifyParkingFull() {
	m.NotifyParkingFullIsCalledTimes += 1
}

func (m *MockAttendant) NotifyParkingSpaceAvailable() {
	m.NotifyParkingSpaceAvailableIsCalledTimes += 1
}
