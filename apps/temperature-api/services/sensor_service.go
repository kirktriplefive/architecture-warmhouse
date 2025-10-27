package services

import (
	"math/rand"
	"time"
)

type SensorService struct {
}

func NewSensorService() *SensorService {
	return &SensorService{}
}

func (s *SensorService) GetSensorTemperature(location string) float64 {
	rand.NewSource(time.Now().UnixNano())
	return rand.Float64()
}
