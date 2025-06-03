package producer

import "main.go/internal/weatherapp"

type Producer interface {
	Produce(shortCityInfo weatherapp.ShortCityInfo) (weatherapp.WeatherMsg, error)
}
