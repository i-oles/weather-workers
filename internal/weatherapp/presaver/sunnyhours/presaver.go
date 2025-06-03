package sunnyhours

import (
	"main.go/internal/weatherapp"
	"main.go/internal/weatherapp/presaver"
)

const preSaverName = "hours_with_full_sun"

var _ presaver.WeatherAppPreSaver = (*PreSaver)(nil)

type PreSaver struct {
	sunnyHours int
	cityName   string
}

func (s *PreSaver) Save(cityResult weatherapp.CityResult) {
	if cityResult.SunnyHoursCount > s.sunnyHours {
		s.sunnyHours = cityResult.SunnyHoursCount
		s.cityName = cityResult.Name
	}
}

func (s *PreSaver) GetResult() weatherapp.Result {
	return weatherapp.Result{
		CityName: s.cityName,
		Value:    s.sunnyHours,
	}
}

func (s *PreSaver) GetName() string {
	return preSaverName
}
