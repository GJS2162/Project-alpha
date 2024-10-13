package models

import "time"

type Booking struct {
    ID              string    `json:"id"`
    UserID          string    `json:"user_id"`
    DriverID        string    `json:"driver_id"`
    VehicleID       string    `json:"vehicle_id"`
    PickupLocation  string    `json:"pickup_location"`
    DropoffLocation string    `json:"dropoff_location"`
    Status          string    `json:"status"`
    Price           float64   `json:"price"`
    CreatedAt       time.Time `json:"created_at"`
    UpdatedAt       time.Time `json:"updated_at"`
}

type BookingRequest struct {
    UserID          string  `json:"user_id" binding:"required"`
    PickupLocation  string  `json:"pickup_location" binding:"required"`
    DropoffLocation string  `json:"dropoff_location" binding:"required"`
    VehicleType     string  `json:"vehicle_type" binding:"required"`
}
