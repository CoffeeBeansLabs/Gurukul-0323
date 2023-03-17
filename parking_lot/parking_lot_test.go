package parking_lot

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParkingLot_ParkVehicle(t *testing.T) {
	parking := NewParkingLot(1)

	// Park a vehicle and check if it is parked
	firstVehicleParkingErr := parking.ParkVehicle("abc")
	assert.Nil(t, firstVehicleParkingErr)
	assert.Equal(t, true, parking.IsVehicleParked("abc"))

	//Park a next vehicle
	SecondVehicleParkingErr := parking.ParkVehicle("def")
	assert.NotNil(t, SecondVehicleParkingErr)
	assert.Equal(t, false, parking.IsVehicleParked("def"))
}
