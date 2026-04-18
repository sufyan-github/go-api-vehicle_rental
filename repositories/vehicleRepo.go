package repositories

import (
	"go-api/config"
	"go-api/models"
)

func CreateVehicle(vehicle *models.Vehicle) error {
	return config.DB.Create(vehicle).Error
}

func GetAllVehicles() ([]models.Vehicle, error) {
	var vehicles []models.Vehicle
	err := config.DB.Find(&vehicles).Error
	return vehicles, err
}

func GetVehicleByID(id uint) (models.Vehicle, error) {
	var vehicle models.Vehicle
	err := config.DB.First(&vehicle, id).Error
	return vehicle, err
}

func UpdateVehicle(vehicle *models.Vehicle) error {
	return config.DB.Save(vehicle).Error
}