package services

import (
	"encoding/json"
	"go-api/config"
	"go-api/models"
	"go-api/repositories"
)

func CreateVehicle(vehicle models.Vehicle) error {
	vehicle.Available = true
	return repositories.CreateVehicle(&vehicle)
}

func GetVehicles() ([]models.Vehicle, error) {

	// Try cache
	cached, err := config.RedisClient.Get(config.Ctx, "vehicles").Result()

	if err == nil {
		var vehicles []models.Vehicle
		json.Unmarshal([]byte(cached), &vehicles)
		return vehicles, nil
	}

	config.RedisClient.Del(config.Ctx, "vehicles")

	// DB fallback
	vehicles, err := repositories.GetAllVehicles()
	if err != nil {
		return vehicles, err
	}

	// Save to cache
	data, _ := json.Marshal(vehicles)
	config.RedisClient.Set(config.Ctx, "vehicles", data, 0)

	return vehicles, nil
}
