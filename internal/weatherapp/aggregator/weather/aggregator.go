package weather

import (
	"main.go/internal/weatherapp"
	"main.go/internal/weatherapp/aggregator"
)

var _ aggregator.WeatherAppAggregator = (*Aggregator)(nil)

type Aggregator struct {
}

func NewAggregator() *Aggregator {
	return &Aggregator{}
}

func (a Aggregator) CountWeatherCode(
	weatherMsg weatherapp.WeatherMsg,
	weatherCode weatherapp.WeatherCode,
) int {
	weatherCodes := weatherMsg.WeatherCodes
	counter := 0

	for _, value := range weatherCodes {
		if value == weatherCode {
			counter++
		}
	}

	return counter
}

func (a Aggregator) CountAverageTemperature(weatherMsg weatherapp.WeatherMsg) float64 {
	var sumTemps float64

	temperatures := weatherMsg.Temperatures
	for _, temp := range temperatures {
		sumTemps += temp
	}

	if len(temperatures) == 0 {
		return 0
	}

	return sumTemps / float64(len(temperatures))
}
