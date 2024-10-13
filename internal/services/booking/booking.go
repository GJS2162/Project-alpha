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

// CreateBooking creates a new booking in the database
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
	log.Printf("Inserting booking: %+v", booking)
	_, err := database.DB.Exec(query, booking.ID, booking.UserID, booking.PickupLocation, booking.DropoffLocation, booking.Status, booking.CreatedAt, booking.UpdatedAt)
	if err != nil {
		log.Printf("Error inserting booking: %v", err)
		return nil, err
	}

	log.Printf("Booking created successfully with ID: %s", booking.ID)
	return booking, nil
}

// GetBooking retrieves a booking by its ID
func GetBooking(id string) (*models.Booking, error) {
	var booking models.Booking
	var createdAt, updatedAt []byte // Change to byte slice for scanning

	query := `SELECT id, user_id, pickup_location, dropoff_location, status, created_at, updated_at FROM bookings WHERE id = ?`

	log.Printf("Querying booking with ID: %s", id)
	err := database.DB.QueryRow(query, id).Scan(
		&booking.ID,
		&booking.UserID,
		&booking.PickupLocation,
		&booking.DropoffLocation,
		&booking.Status,
		&createdAt, // Scan into byte slice
		&updatedAt, // Scan into byte slice
	)

	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("No booking found with ID: %s", id)
			return nil, fmt.Errorf("Booking not found")
		}
		log.Printf("Error retrieving booking: %v", err)
		return nil, err
	}

	// Parse byte slices to time.Time
	var errCreatedAt, errUpdatedAt error
	booking.CreatedAt, errCreatedAt = time.Parse("2006-01-02 15:04:05", string(createdAt))
	if errCreatedAt != nil {
		log.Printf("Error parsing created_at: %v", errCreatedAt)
		return nil, errCreatedAt
	}

	booking.UpdatedAt, errUpdatedAt = time.Parse("2006-01-02 15:04:05", string(updatedAt))
	if errUpdatedAt != nil {
		log.Printf("Error parsing updated_at: %v", errUpdatedAt)
		return nil, errUpdatedAt
	}

	return &booking, nil
}
