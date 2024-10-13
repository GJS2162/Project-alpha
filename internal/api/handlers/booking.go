package handlers

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "logistics-platform/internal/models"
    "logistics-platform/internal/services/booking"
)

func CreateBooking(c *gin.Context) {
    var bookingReq models.BookingRequest
    if err := c.ShouldBindJSON(&bookingReq); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    newBooking, err := booking.CreateBooking(bookingReq)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, newBooking)
}

func GetBooking(c *gin.Context) {
    id := c.Param("id")

    bookingDetails, err := booking.GetBooking(id)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Booking not found"})
        return
    }

    c.JSON(http.StatusOK, bookingDetails)
}
