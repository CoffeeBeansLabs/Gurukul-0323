package parking_lot

import (
	"Gurukul-0323/parking_lot/mocks"
	"Gurukul-0323/parking_lot/model"
	"errors"
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
	if err != nil {
		t.Errorf("Expected nil, got error: %v", err)
	}
}

// Test parking vehicle when all parking lots are full
func TestAttendant_ParkVehicle_Error(t *testing.T) {
	parkinglots := []model.IParkingLot{
		&mocks.MockParkingLot{IsFullVal: true, IsVehicleParkedVal: false},
		&mocks.MockParkingLot{IsFullVal: true, IsVehicleParkedVal: false},
	}
	attendant := NewAttendant(parkinglots, mocks.MockStrategy{
		ParkingLot:    parkinglots[0],
		ErrorToReturn: ParkingFullError,
	})

	err := attendant.ParkVehicle("ABC123")
	if !errors.Is(err, ParkingFullError) {
		t.Errorf("Expected error: %v, got: %v", ParkingFullError, err)
	}
}

// Test unparking vehicle when the vehicle is parked in one of the lots
func TestAttendant_UnparkVehicle_Success(t *testing.T) {
	parkinglots := []model.IParkingLot{
		&mocks.MockParkingLot{IsFullVal: false, IsVehicleParkedVal: false},
		&mocks.MockParkingLot{IsFullVal: true, IsVehicleParkedVal: true},
	}
	attendant := NewAttendant(parkinglots, mocks.MockStrategy{})

	err := attendant.UnparkVehicle("ABC123")
	if err != nil {
		t.Errorf("Expected nil, got error: %v", err)
	}
}

// Test unparking vehicle when the vehicle is not parked in any of the lots
func TestAttendant_UnparkVehicle_Error(t *testing.T) {
	parkinglots := []model.IParkingLot{
		&mocks.MockParkingLot{IsFullVal: false, IsVehicleParkedVal: false},
		&mocks.MockParkingLot{IsFullVal: false, IsVehicleParkedVal: false},
	}
	attendant := NewAttendant(parkinglots, mocks.MockStrategy{})

	err := attendant.UnparkVehicle("ABC123")
	if !errors.Is(err, VehicleNotfoundError) {
		t.Errorf("Expected error: %v, got: %v", VehicleNotfoundError, err)
	}
}
