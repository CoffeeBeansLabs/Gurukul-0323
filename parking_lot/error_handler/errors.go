package error_handler

import "errors"

var (
	ParkingFullError     = errors.New("parking lot is full")
	VehicleNotfoundError = errors.New("vehicle not found")
	VehicleAlreadyParked = errors.New("vehicle already parked")
)
