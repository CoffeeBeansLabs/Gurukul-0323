package parking_lot

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

type MockOwner struct {
	mock.Mock
	Notifications chan bool
}

func (o *MockOwner) NotifyParkingFull() {
	fmt.Println("NotifyParkingFull called")
	o.Called()
}

func TestParkingLot_ParkVehicle(t *testing.T) {

	// Create a mocked owner.
	mockOwner := &MockOwner{
		Notifications: make(chan bool),
	}
	mockOwner.On("NotifyParkingFull")

	parking := NewParkingLot(1, &Owner{Notifications: mockOwner.Notifications})

	// Park a vehicle and check if it is parked
	firstVehicleParkingErr := parking.ParkVehicle("abc")
	assert.Nil(t, firstVehicleParkingErr)
	assert.Equal(t, true, parking.IsVehicleParked("abc"))

	//Park a next vehicle
	SecondVehicleParkingErr := parking.ParkVehicle("def")
	assert.NotNil(t, SecondVehicleParkingErr)
	assert.Equal(t, false, parking.IsVehicleParked("def"))

	// Verify that the mocked owner was called when the parking lot was full.
	mockOwner.AssertCalled(t, "NotifyParkingFull")

	// Wait for a notification from the owner.
	select {
	case <-mockOwner.Notifications:
		// Notification received, test passed.
	case <-time.After(1 * time.Second):
		// Timeout, test failed.
		assert.Fail(t, "Timed out waiting for notification")
	}

}

//func TestParkingLot_UnParkVehicle(t *testing.T) {
//	parking := NewParkingLot(1)
//
//	// Park a vehicle and check if it is parked
//	firstVehicleParkingErr := parking.ParkVehicle("abc")
//	assert.Nil(t, firstVehicleParkingErr)
//	assert.Equal(t, true, parking.IsVehicleParked("abc"))
//
//	//Unpark vehicle and check if it is not Parked
//	VehicleUnParkingErr := parking.UnparkVehicle("abc")
//	assert.Nil(t, VehicleUnParkingErr)
//	assert.Equal(t, false, parking.IsVehicleParked("abc"))
//}
