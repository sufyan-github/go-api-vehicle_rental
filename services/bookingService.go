package services

import (
	"errors"

	"go-api/models"
	"go-api/repositories"
)

import "go-api/utils"

func CreateBooking(userID uint, booking models.Booking) (models.Booking, error) {

	vehicle, err := repositories.GetVehicleByID(booking.VehicleID)
	if err != nil {
		return booking, errors.New("vehicle not found")
	}

	if !vehicle.Available {
		return booking, errors.New("vehicle not available")
	}

	vehicle.Available = false
	repositories.UpdateVehicle(&vehicle)

	booking.UserID = userID
	booking.Status = "booked"

	err = repositories.CreateBooking(&booking)
	if err == nil {
		utils.PublishBookingEvent(booking)
	}
	return booking, err
}
func GetBookings(userID interface{}) ([]models.Booking, error) {
	return repositories.GetBookingsByUser(userID)
}

func CancelBooking(id string) error {

	booking, err := repositories.GetBookingByID(id)
	if err != nil {
		return errors.New("booking not found")
	}

	booking.Status = "cancelled"
	repositories.UpdateBooking(&booking)

	vehicle, _ := repositories.GetVehicleByID(booking.VehicleID)
	vehicle.Available = true
	repositories.UpdateVehicle(&vehicle)

	return nil
}