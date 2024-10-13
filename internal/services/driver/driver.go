package driver

import (
	"logistics-platform/internal/database"
	"logistics-platform/internal/models"

	"github.com/google/uuid"
)

func CreateDriver(req models.DriverRequest) (*models.Driver, error) {
	// Implement driver creation logic here
	// This is a simplified version and doesn't include all necessary checks

	driver := &models.Driver{
		ID:        uuid.New().String(),
		Name:      req.Name,
		VehicleID: req.VehicleID,
		Status:    "AVAILABLE",
	}

	// Here you would typically save the driver to the database

	return driver, nil
}

func GetDriver(id string) (*models.Driver, error) {
	// Implement driver retrieval logic here
	// This is a placeholder implementation

	var driver models.Driver
	err := database.DB.QueryRow("SELECT id, name, vehicle_id, status FROM drivers WHERE id = $1", id).
		Scan(&driver.ID, &driver.Name, &driver.VehicleID, &driver.Status)

	if err != nil {
		return nil, err
	}

	return &driver, nil
}
