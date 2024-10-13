package booking

import (
	"database/sql"
	"fmt"
	"log"
	"logistics-platform/internal/database"
	"logistics-platform/internal/models"
	"time"

	"github.com/google/uuid"
)

func CreateBooking(req models.BookingRequest) (*models.Booking, error) {
	booking := &models.Booking{
		ID:              uuid.New().String(),
		UserID:          req.UserID,
		PickupLocation:  req.PickupLocation,
		DropoffLocation: req.DropoffLocation,
		Status:          "PENDING",
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}

	query := `INSERT INTO bookings (id, user_id, pickup_location, dropoff_location, status, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?)`
	_, err := database.DB.Exec(query, booking.ID, booking.UserID, booking.PickupLocation, booking.DropoffLocation, booking.Status, booking.CreatedAt, booking.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return booking, nil
}

func GetBooking(id string) (*models.Booking, error) {
	var booking models.Booking
	query := `SELECT * FROM bookings WHERE id = ?`

	log.Printf("Querying booking with ID: %s", id)
	// Execute the query with the provided id
	err := database.DB.QueryRow(query, id).Scan(&booking.ID, &booking.UserID, &booking.PickupLocation, &booking.DropoffLocation, &booking.Status, &booking.CreatedAt, &booking.UpdatedAt)

	if err != nil {
		// Check if the error is due to no rows found
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("Booking not found")
		}
		return nil, fmt.Errorf("error retrieving booking: %v", err)
	}

	return &booking, nil
}
