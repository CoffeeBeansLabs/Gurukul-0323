package parking_lot

import (
	"Gurukul-0323/parking_lot/error_handler"
	"Gurukul-0323/parking_lot/mocks"
	"Gurukul-0323/parking_lot/model"
	"Gurukul-0323/parking_lot/strategies"
	"github.com/stretchr/testify/assert"
	"testing"
)

// Test parking vehicle when there is space available
func TestAttendant_ParkVehicle_Success(t *testing.T) {
	parkinglots := []model.IParkingLot{
		&mocks.MockParkingLot{IsFullVal: false, IsVehicleParkedVal: false},
		&mocks.MockParkingLot{IsFullVal: true, IsVehicleParkedVal: false},
	}
	attendant := NewAttendant(parkinglots, mocks.MockStrategy{
		ParkingLot: parkinglots[0],
	})

	err := attendant.ParkVehicle("ABC123")
	assert.Nil(t, err)

}

// Test parking vehicle when all parking lots are full
func TestAttendant_ParkVehicle_Error(t *testing.T) {
	parkinglots := []model.IParkingLot{
		&mocks.MockParkingLot{IsFullVal: true, IsVehicleParkedVal: false},
		&mocks.MockParkingLot{IsFullVal: true, IsVehicleParkedVal: false},
	}
	attendant := NewAttendant(parkinglots, mocks.MockStrategy{
		ParkingLot:    parkinglots[0],
		ErrorToReturn: error_handler.ParkingFullError,
	})

	err := attendant.ParkVehicle("ABC123")
	assert.Equal(t, error_handler.ParkingFullError, err)
}

// Test unparking vehicle when the vehicle is parked in one of the lots
func TestAttendant_UnparkVehicle_Success(t *testing.T) {
	parkinglots := []model.IParkingLot{
		&mocks.MockParkingLot{IsFullVal: false, IsVehicleParkedVal: false},
		&mocks.MockParkingLot{IsFullVal: true, IsVehicleParkedVal: true},
	}
	attendant := NewAttendant(parkinglots, mocks.MockStrategy{})

	err := attendant.UnparkVehicle("ABC123")
	assert.Nil(t, err)

}

// Test unparking vehicle when the vehicle is not parked in any of the lots
func TestAttendant_UnparkVehicle_Error(t *testing.T) {
	parkinglots := []model.IParkingLot{
		&mocks.MockParkingLot{IsFullVal: false, IsVehicleParkedVal: false},
		&mocks.MockParkingLot{IsFullVal: false, IsVehicleParkedVal: false},
	}
	attendant := NewAttendant(parkinglots, mocks.MockStrategy{})

	err := attendant.UnparkVehicle("ABC123")
	assert.Equal(t, error_handler.VehicleNotfoundError, err)

}

func TestAttendant_SelectParkingLot_MaxCapacity(t *testing.T) {
	parkingLot1 := NewParkingLot(3, nil)
	parkingLot2 := NewParkingLot(2, nil)
	parkingLot3 := NewParkingLot(4, nil)

	attendant := NewAttendant(
		[]model.IParkingLot{parkingLot1, parkingLot2, parkingLot3},
		strategies.MaxCapacity{},
	)

	// Fill up all parking lots
	err := attendant.ParkVehicle("KA-01-HH-1234")
	assert.Nil(t, err)

	// Select parking lot with max capacity
	lot, err := strategies.MaxCapacity{}.SelectParkingLot([]model.IParkingLot{parkingLot1, parkingLot2, parkingLot3})
	assert.Nil(t, err)
	assert.Equal(t, lot, parkingLot3)

	assert.Equal(t, true, parkingLot3.IsVehicleParked("KA-01-HH-1234"))
	assert.Equal(t, false, parkingLot1.IsVehicleParked("KA-01-HH-1234"))
	assert.Equal(t, false, parkingLot2.IsVehicleParked("KA-01-HH-1234"))
}

func TestAttendant_SelectParkingLot_MaxFreeSpace(t *testing.T) {
	parkingLot1 := NewParkingLot(3, nil)
	parkingLot2 := NewParkingLot(6, nil)

	attendant := NewAttendant(
		[]model.IParkingLot{parkingLot1, parkingLot2},
		strategies.MaxAvailableSpace{},
	)

	// Fill up all parking lots
	err := attendant.ParkVehicle("KA-01-HH-1234")
	assert.Nil(t, err)

	// Select parking lot with max capacity
	lot, err := strategies.MaxAvailableSpace{}.SelectParkingLot([]model.IParkingLot{parkingLot1, parkingLot2})
	assert.Nil(t, err)
	assert.Equal(t, lot, parkingLot2)

	assert.Equal(t, false, parkingLot1.IsVehicleParked("KA-01-HH-1234"))
	assert.Equal(t, true, parkingLot2.IsVehicleParked("KA-01-HH-1234"))
}

func TestAttendant_SelectParkingLot_FirstComeFirstServe(t *testing.T) {
	parkingLot1 := NewParkingLot(3, nil)
	parkingLot2 := NewParkingLot(6, nil)

	attendant := NewAttendant(
		[]model.IParkingLot{parkingLot1, parkingLot2},
		strategies.FirstComeFirstServe{},
	)

	// Fill up all parking lots
	err := attendant.ParkVehicle("KA-01-HH-1234")
	assert.Nil(t, err)

	// Select parking lot with max capacity
	lot, err := strategies.FirstComeFirstServe{}.SelectParkingLot([]model.IParkingLot{parkingLot1, parkingLot2})
	assert.Nil(t, err)
	assert.Equal(t, lot, parkingLot1)

	assert.Equal(t, true, parkingLot1.IsVehicleParked("KA-01-HH-1234"))
	assert.Equal(t, false, parkingLot2.IsVehicleParked("KA-01-HH-1234"))
}
