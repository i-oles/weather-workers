package memory

import (
	"main.go/internal/weatherapp"
	"main.go/internal/weatherapp/aggregator"
	weatherAppConsumer "main.go/internal/weatherapp/consumer"
	"main.go/internal/weatherapp/saver"
)

const (
	fogWeatherCode      = 45
	sunnyDayWeatherCode = 0
)

var _ weatherAppConsumer.Consumer = (*Consumer)(nil)

type Consumer struct {
	aggregator aggregator.WeatherAppAggregator
	saver      saver.WeatherAppSaver
}

func New(aggregator aggregator.WeatherAppAggregator, saver saver.WeatherAppSaver) *Consumer {
	return &Consumer{aggregator: aggregator, saver: saver}
}

func (c Consumer) Consume(msg weatherapp.WeatherMsg) {
	foggyDayCount := c.aggregator.CountWeatherCode(msg, fogWeatherCode)
	sunnyDayCount := c.aggregator.CountWeatherCode(msg, sunnyDayWeatherCode)
	averageTemp := c.aggregator.CountAverageTemperature(msg)

	cityResult := weatherapp.CityResult{
		Name:            msg.CityName,
		FoggyHoursCount: foggyDayCount,
		TempAverage:     averageTemp,
		SunnyHoursCount: sunnyDayCount,
	}

	c.saver.Save(cityResult)
}
