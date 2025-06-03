package avgtemp

import (
	"main.go/internal/weatherapp"
	"main.go/internal/weatherapp/presaver"
)

const preSaverName = "highest_avg_temp"

var _ presaver.WeatherAppPreSaver = (*PreSaver)(nil)

type PreSaver struct {
	highestAvgTemp *float64
	cityName       string
}

func (s *PreSaver) Save(cityResult weatherapp.CityResult) {
	if s.highestAvgTemp == nil {
		s.highestAvgTemp = &cityResult.TempAverage
		s.cityName = cityResult.Name

		return
	}

	if cityResult.TempAverage > *s.highestAvgTemp {
		s.highestAvgTemp = &cityResult.TempAverage
		s.cityName = cityResult.Name
	}
}

func (s *PreSaver) GetResult() weatherapp.Result {
	return weatherapp.Result{
		CityName: s.cityName,
		Value:    *s.highestAvgTemp,
	}
}

func (s *PreSaver) GetName() string {
	return preSaverName
}
