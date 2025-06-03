package memory

import (
	"fmt"

	"main.go/internal/weatherapp"
	"main.go/internal/weatherapp/openmeteo"
	weatherAppProducer "main.go/internal/weatherapp/producer"
)

const (
	weatherCodeTag   = "weathercode"
	temperature2mTag = "temperature_2m"
)

var _ weatherAppProducer.Producer = (*Producer)(nil)

type Producer struct {
	openMeteoAPI openmeteo.API
}

func New(openMeteoAPI openmeteo.API) *Producer {
	return &Producer{openMeteoAPI: openMeteoAPI}
}

func (p Producer) Produce(cityInfo weatherapp.ShortCityInfo) (weatherapp.WeatherMsg, error) {
	var msg weatherapp.WeatherMsg

	weatherResp, err := p.openMeteoAPI.GetWeather(
		cityInfo.Latitude, cityInfo.Longitude, weatherCodeTag,
	)
	if err != nil {
		return weatherapp.WeatherMsg{}, fmt.Errorf("error while getting weather: %w", err)
	}

	msg.WeatherCodes = weatherResp.Hourly.WeatherCodes

	weatherResp, err = p.openMeteoAPI.GetWeather(
		cityInfo.Latitude, cityInfo.Longitude, temperature2mTag,
	)
	if err != nil {
		return weatherapp.WeatherMsg{}, fmt.Errorf("error while getting weather: %w", err)
	}

	msg.Temperatures = weatherResp.Hourly.Temperature2m
	msg.CityName = cityInfo.Name

	return msg, nil
}
