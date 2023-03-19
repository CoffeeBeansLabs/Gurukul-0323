package parking_lot

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParkingLot_ParkVehicle(t *testing.T) {
	owner := NewOwner()
	parking := NewParkingLot(1, owner)

	// Park a vehicle and check if it is parked
	firstVehicleParkingErr := parking.ParkVehicle("abc")
	assert.Nil(t, firstVehicleParkingErr)
	assert.Equal(t, true, parking.IsVehicleParked("abc"))

	//Park a next vehicle
	SecondVehicleParkingErr := parking.ParkVehicle("def")
	assert.NotNil(t, SecondVehicleParkingErr)
	assert.Equal(t, false, parking.IsVehicleParked("def"))

	//Owner is notified
	assert.Equal(t, true, owner.isNotified)
}

func TestParkingLot_UnParkVehicle(t *testing.T) {
	owner := NewOwner()
	parking := NewParkingLot(1, owner)

	// Park a vehicle and check if it is parked
	firstVehicleParkingErr := parking.ParkVehicle("abc")
	assert.Nil(t, firstVehicleParkingErr)
	assert.Equal(t, true, parking.IsVehicleParked("abc"))

	//Unpark vehicle and check if it is not Parked
	VehicleUnParkingErr := parking.UnparkVehicle("abc")
	assert.Nil(t, VehicleUnParkingErr)
	assert.Equal(t, false, parking.IsVehicleParked("abc"))

}
