package api

import (
	"logistics-platform/internal/api/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		v1.POST("/bookings", handlers.CreateBooking)
		v1.GET("/bookings/:id", handlers.GetBooking)
		v1.POST("/drivers", handlers.CreateDriver)
		v1.GET("/drivers/:id", handlers.GetDriver)
		// v1.POST("/users", handlers.CreateUser)
		// v1.GET("/users/:id", handlers.GetUser)
		// v1.POST("/vehicles", handlers.CreateVehicle)
		// v1.GET("/vehicles/:id", handlers.GetVehicle)
	}
}
