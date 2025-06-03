package openmeteo

import "main.go/internal/weatherapp"

type API interface {
	GetWeather(latitude, longitude, weatherTag string) (weatherapp.WeatherStats, error)
}
