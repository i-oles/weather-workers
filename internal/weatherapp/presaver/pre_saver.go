package presaver

import "main.go/internal/weatherapp"

type WeatherAppPreSaver interface {
	GetName() string
	Save(result weatherapp.CityResult)
	GetResult() weatherapp.Result
}
