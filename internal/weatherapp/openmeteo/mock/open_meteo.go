package mock

import (
	"time"

	"main.go/internal/weatherapp"
	"main.go/internal/weatherapp/openmeteo"
)

const mockProcessTime = 100 * time.Millisecond

var _ openmeteo.API = (*OpenMeteo)(nil)

type OpenMeteo struct {
}

func NewOpenMeteo() *OpenMeteo {
	return &OpenMeteo{}
}

func (p *OpenMeteo) GetWeather(_, _, _ string) (weatherapp.WeatherStats, error) {
	var weatherStats weatherapp.WeatherStats
	weatherStats.Hourly.Temperature2m = []float64{1.0, 2.0, 3.0, 4.0, 5.0}
	weatherStats.Hourly.WeatherCodes = []weatherapp.WeatherCode{1, 2, 3, 4, 5}

	time.Sleep(mockProcessTime)

	return weatherStats, nil
}
