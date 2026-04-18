package repositories

import (
	"go-api/config"
	"go-api/models"
)

func CreateBooking(booking *models.Booking) error {
	return config.DB.Create(booking).Error
}

func GetBookingsByUser(userID interface{}) ([]models.Booking, error) {
	var bookings []models.Booking
	err := config.DB.Where("user_id = ?", userID).Find(&bookings).Error
	return bookings, err
}

func GetBookingByID(id string) (models.Booking, error) {
	var booking models.Booking
	err := config.DB.First(&booking, id).Error
	return booking, err
}

func UpdateBooking(booking *models.Booking) error {
	return config.DB.Save(booking).Error
}