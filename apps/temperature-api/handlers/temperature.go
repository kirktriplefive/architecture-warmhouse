package handlers

import (
	"fmt"
	"net/http"
	"time"

	"temperature-api/services"

	"github.com/gin-gonic/gin"
)

// SensorHandler handles sensor-related requests
type SensorHandler struct {
	Sensor services.SensorUseCase
}

// NewSensorHandler creates a new SensorHandler
func NewTemperatureHandler(service services.SensorUseCase) *SensorHandler {
	return &SensorHandler{
		Sensor: service,
	}
}

// RegisterRoutes registers the sensor routes
func (h *SensorHandler) RegisterRoutes(router *gin.RouterGroup) {
	router.GET("/temperature", h.GetTemperatureByLocation)
	router.GET("/temperature/:sensor", h.GetTemperatureBySensor)
}

// GetTemperatureByLocation handles GET /api/v1/sensors/temperature/:location
func (h *SensorHandler) GetTemperatureByLocation(c *gin.Context) {
	location := c.Request.URL.Query().Get("location")
	if location == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Location is required"})
		return
	}

	var sensorID string

	switch location {
	case "Living Room":
		sensorID = "1"
	case "Bedroom":
		sensorID = "2"
	case "Kitchen":
		sensorID = "3"
	default:
		sensorID = "0"
	}

	temperature := h.Sensor.GetSensorTemperature(location)

	// Return the temperature data
	c.JSON(http.StatusOK, gin.H{
		"location":    location,
		"value":       temperature,
		"unit":        sensorID,
		"status":      location,
		"timestamp":   time.Now(),
		"description": location,
	})
}

// GetTemperatureByLocation handles GET /api/v1/sensors/temperature/:location
func (h *SensorHandler) GetTemperatureBySensor(c *gin.Context) {
	sensorID := c.Param("sensor")
	if sensorID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "sensorID is required"})
		return
	}

	var location string

	switch sensorID {
	case "1":
		location = "Living Room"
	case "2":
		location = "Bedroom"
	case "3":
		location = "Kitchen"
	default:
		location = "Unknown"
	}

	temperature := h.Sensor.GetSensorTemperature(location)

	fmt.Println(temperature)

	// Return the temperature data
	c.JSON(http.StatusOK, gin.H{
		"location":    location,
		"value":       temperature,
		"unit":        sensorID,
		"status":      location,
		"timestamp":   time.Now(),
		"description": location,
	})
}
