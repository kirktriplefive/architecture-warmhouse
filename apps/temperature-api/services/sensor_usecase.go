package services

type SensorUseCase interface {
	GetSensorTemperature(location string) float64
}
