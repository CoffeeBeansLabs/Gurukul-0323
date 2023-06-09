package parking_lot

import (
	"Gurukul-0323/parking_lot/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

//func TestParkingLot_ParkVehicle(t *testing.T) {
//	Owner := mocks.MockOwner{}
//	parking := NewParkingLot(1, &Owner)
//
//	// Park a vehicle and check if it is parked
//	firstVehicleParkingErr := parking.ParkVehicle("abc")
//	assert.Nil(t, firstVehicleParkingErr)
//	assert.Equal(t, true, parking.IsVehicleParked("abc"))
//
//	//Park a next vehicle
//	SecondVehicleParkingErr := parking.ParkVehicle("def")
//	assert.NotNil(t, SecondVehicleParkingErr)
//	assert.Equal(t, false, parking.IsVehicleParked("def"))
//
//	//Owner is notified
//	// assert.Equal(t, true, Owner.isNotified)
//	assert.Equal(t, 2, Owner.NotifyParkingFullIsCalledTimes)
//}
//
//func TestParkingLot_UnParkVehicle(t *testing.T) {
//	Owner := mocks.MockOwner{}
//	parking := NewParkingLot(2, &Owner)
//
//	// Park a vehicle and check if it is parked
//	firstVehicleParkingErr := parking.ParkVehicle("abc")
//	assert.Nil(t, firstVehicleParkingErr)
//	assert.Equal(t, true, parking.IsVehicleParked("abc"))
//
//	secondVehicleParkingErr := parking.ParkVehicle("cde")
//	assert.Nil(t, secondVehicleParkingErr)
//	assert.Equal(t, true, parking.IsVehicleParked("cde"))
//
//	assert.Equal(t, 1, Owner.NotifyParkingFullIsCalledTimes)
//
//	// Unpark vehicle and check if it is not Parked
//	VehicleUnParkingErr := parking.UnparkVehicle("abc")
//	assert.Nil(t, VehicleUnParkingErr)
//	assert.Equal(t, false, parking.IsVehicleParked("abc"))
//	assert.Equal(t, 1, Owner.NotifyParkingSpaceAvailableIsCalledTimes)
//}

func TestParkingLot_ParkVehicle_Owner_observer(t *testing.T) {
	//Owner := mocks.MockOwner{}
	owner := mocks.MockAttendant{}
	parking := NewParkingLot(1, []IObserver{&owner})

	// Park a vehicle and check if it is parked
	firstVehicleParkingErr := parking.ParkVehicle("abc")
	assert.Nil(t, firstVehicleParkingErr)
	assert.Equal(t, true, parking.IsVehicleParked("abc"))

	//Park a next vehicle
	SecondVehicleParkingErr := parking.ParkVehicle("def")
	assert.NotNil(t, SecondVehicleParkingErr)
	assert.Equal(t, false, parking.IsVehicleParked("def"))

	//Owner is notified
	// assert.Equal(t, true, Owner.isNotified)
	assert.Equal(t, 2, owner.NotifyParkingFullIsCalledTimes)
}

func TestParkingLot_ParkVehicle_Attendant_observer(t *testing.T) {
	//Owner := mocks.MockOwner{}
	attendant := mocks.MockAttendant{}
	parking := NewParkingLot(1, []IObserver{&attendant})

	// Park a vehicle and check if it is parked
	firstVehicleParkingErr := parking.ParkVehicle("abc")
	assert.Nil(t, firstVehicleParkingErr)
	assert.Equal(t, true, parking.IsVehicleParked("abc"))

	//Park a next vehicle
	SecondVehicleParkingErr := parking.ParkVehicle("def")
	assert.NotNil(t, SecondVehicleParkingErr)
	assert.Equal(t, false, parking.IsVehicleParked("def"))

	//Owner is notified
	// assert.Equal(t, true, Owner.isNotified)
	assert.Equal(t, 2, attendant.NotifyParkingFullIsCalledTimes)
}

func TestParkingLot_ParkVehicle_Attendant_Owner_observer(t *testing.T) {
	//Owner := mocks.MockOwner{}
	attendant := mocks.MockAttendant{}
	owner := mocks.MockOwner{}
	parking := NewParkingLot(1, []IObserver{&attendant, &owner})

	// Park a vehicle and check if it is parked
	firstVehicleParkingErr := parking.ParkVehicle("abc")
	assert.Nil(t, firstVehicleParkingErr)
	assert.Equal(t, true, parking.IsVehicleParked("abc"))

	//Park a next vehicle
	SecondVehicleParkingErr := parking.ParkVehicle("def")
	assert.NotNil(t, SecondVehicleParkingErr)
	assert.Equal(t, false, parking.IsVehicleParked("def"))

	//Owner is notified
	// assert.Equal(t, true, Owner.isNotified)
	assert.Equal(t, 2, attendant.NotifyParkingFullIsCalledTimes)
	assert.Equal(t, 2, owner.NotifyParkingFullIsCalledTimes)
}
