package memory

import (
	"main.go/internal/weatherapp"
	"main.go/internal/weatherapp/presaver"
	"main.go/internal/weatherapp/saver"
)

var _ saver.WeatherAppSaver = (*Saver)(nil)

type Saver struct {
	preSavers []presaver.WeatherAppPreSaver
}

func NewSaver(preSavers []presaver.WeatherAppPreSaver) Saver {
	return Saver{preSavers: preSavers}
}

func (s Saver) Save(cityResult weatherapp.CityResult) {
	for _, preSaver := range s.preSavers {
		preSaver.Save(cityResult)
	}
}

func (s Saver) GetResults() map[string]weatherapp.Result {
	results := make(map[string]weatherapp.Result, len(s.preSavers))

	for _, preSaver := range s.preSavers {
		cityName := preSaver.GetName()
		result := preSaver.GetResult()
		results[cityName] = result
	}

	return results
}
