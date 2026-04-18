package models

import "gorm.io/gorm"

type Vehicle struct {
	gorm.Model
	Name        string  `json:"name"`
	Brand       string  `json:"brand"`
	PricePerDay float64 `json:"price_per_day"`
	Available   bool    `json:"available"`
}