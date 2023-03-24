package parking_lot

import (
	"Gurukul-0323/parking_lot/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAttendant_ParkVehicle(t *testing.T) {
	owner := mocks.MockOwner{}
	parking := NewParkingLot(1, &owner)
	attendant := NewAttendant(*parking)

	// Park a vehicle and check if it is parked
	firstVehicleParkingErr := attendant.parkingLot.ParkVehicle("abc")
	assert.Nil(t, firstVehicleParkingErr)
	assert.Equal(t, true, parking.IsVehicleParked("abc"))

	//Park a next vehicle
	SecondVehicleParkingErr := attendant.parkingLot.ParkVehicle("def")
	assert.NotNil(t, SecondVehicleParkingErr)
	assert.Equal(t, false, attendant.parkingLot.IsVehicleParked("def"))

}

func TestAttendant_UnParkVehicle(t *testing.T) {
	owner := mocks.MockOwner{}
	parking := NewParkingLot(2, &owner)
	attendant := NewAttendant(*parking)

	// Park a vehicle and check if it is parked
	firstVehicleParkingErr := attendant.parkingLot.ParkVehicle("abc")
	assert.Nil(t, firstVehicleParkingErr)
	assert.Equal(t, true, parking.IsVehicleParked("abc"))

	secondVehicleParkingErr := attendant.parkingLot.ParkVehicle("cde")
	assert.Nil(t, secondVehicleParkingErr)
	assert.Equal(t, true, parking.IsVehicleParked("cde"))

	// Unpark vehicle and check if it is not Parked
	VehicleUnParkingErr := attendant.parkingLot.UnparkVehicle("abc")
	assert.Nil(t, VehicleUnParkingErr)
	assert.Equal(t, false, attendant.parkingLot.IsVehicleParked("abc"))
}
