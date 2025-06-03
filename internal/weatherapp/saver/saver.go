package saver

import "main.go/internal/weatherapp"

type WeatherAppSaver interface {
	Save(result weatherapp.CityResult)
	GetResults() map[string]weatherapp.Result
}
