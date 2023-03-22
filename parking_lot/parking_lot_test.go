package parking_lot

import (
	"testing"

	"Gurukul-0323/parking_lot/mocks"

	"github.com/stretchr/testify/assert"
)

func TestParkingLot_ParkVehicle(t *testing.T) {
	owner := mocks.MockOwner{}
	parking := NewParkingLot(1, &owner)

	// Park a vehicle and check if it is parked
	firstVehicleParkingErr := parking.ParkVehicle("abc")
	assert.Nil(t, firstVehicleParkingErr)
	assert.Equal(t, true, parking.IsVehicleParked("abc"))

	//Park a next vehicle
	SecondVehicleParkingErr := parking.ParkVehicle("def")
	assert.NotNil(t, SecondVehicleParkingErr)
	assert.Equal(t, false, parking.IsVehicleParked("def"))

	//Owner is notified
	// assert.Equal(t, true, owner.isNotified)
	assert.Equal(t, 2, owner.NotifyParkingFullIsCalledTimes)
}

func TestParkingLot_UnParkVehicle(t *testing.T) {
	owner := mocks.MockOwner{}
	parking := NewParkingLot(2, &owner)

	// Park a vehicle and check if it is parked
	firstVehicleParkingErr := parking.ParkVehicle("abc")
	assert.Nil(t, firstVehicleParkingErr)
	assert.Equal(t, true, parking.IsVehicleParked("abc"))

	secondVehicleParkingErr := parking.ParkVehicle("cde")
	assert.Nil(t, secondVehicleParkingErr)
	assert.Equal(t, true, parking.IsVehicleParked("cde"))

	assert.Equal(t, 1, owner.NotifyParkingFullIsCalledTimes)

	// Unpark vehicle and check if it is not Parked
	VehicleUnParkingErr := parking.UnparkVehicle("abc")
	assert.Nil(t, VehicleUnParkingErr)
	assert.Equal(t, false, parking.IsVehicleParked("abc"))
	assert.Equal(t, 1, owner.NotifyParkingSpaceAvailableIsCalledTimes)
}
