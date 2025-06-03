package aggregator

import "main.go/internal/weatherapp"

type WeatherAppAggregator interface {
	CountAverageTemperature(weatherMsg weatherapp.WeatherMsg) float64
	CountWeatherCode(weatherMsg weatherapp.WeatherMsg, weatherCode weatherapp.WeatherCode) int
}
