package foggyhours

import (
	"main.go/internal/weatherapp"
	"main.go/internal/weatherapp/presaver"
)

const preSaverName = "hours_with_fog"

var _ presaver.WeatherAppPreSaver = (*PreSaver)(nil)

type PreSaver struct {
	foggyHours int
	cityName   string
}

func (s *PreSaver) Save(cityResult weatherapp.CityResult) {
	if cityResult.FoggyHoursCount > s.foggyHours {
		s.foggyHours = cityResult.FoggyHoursCount
		s.cityName = cityResult.Name
	}
}

func (s *PreSaver) GetResult() weatherapp.Result {
	return weatherapp.Result{
		CityName: s.cityName,
		Value:    s.foggyHours,
	}
}

func (s *PreSaver) GetName() string {
	return preSaverName
}
