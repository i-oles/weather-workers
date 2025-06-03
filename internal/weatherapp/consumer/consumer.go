package consumer

import "main.go/internal/weatherapp"

type Consumer interface {
	Consume(weatherMsg weatherapp.WeatherMsg)
}
