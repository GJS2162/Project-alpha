package models

type Driver struct {
    ID        string  `json:"id"`
    Name      string  `json:"name"`
    VehicleID string  `json:"vehicle_id"`
    Status    string  `json:"status"`
    Latitude  float64 `json:"latitude"`
    Longitude float64 `json:"longitude"`
}

type DriverRequest struct {
    Name      string `json:"name" binding:"required"`
    VehicleID string `json:"vehicle_id" binding:"required"`
}
