package handlers

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "logistics-platform/internal/models"
    "logistics-platform/internal/services/driver"
)

func CreateDriver(c *gin.Context) {
    var driverReq models.DriverRequest
    if err := c.ShouldBindJSON(&driverReq); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    newDriver, err := driver.CreateDriver(driverReq)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, newDriver)
}

func GetDriver(c *gin.Context) {
    id := c.Param("id")

    driverDetails, err := driver.GetDriver(id)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Driver not found"})
        return
    }

    c.JSON(http.StatusOK, driverDetails)
}
