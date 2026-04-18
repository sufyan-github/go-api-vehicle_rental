package models

import "gorm.io/gorm"

type Booking struct {
	gorm.Model
	UserID    uint   `json:"user_id"`
	VehicleID uint   `json:"vehicle_id"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
	Status    string `json:"status"`
}