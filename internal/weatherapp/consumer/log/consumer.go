package log

import (
	"log/slog"

	"main.go/internal/weatherapp"
	weatherAppConsumer "main.go/internal/weatherapp/consumer"
)

var _ weatherAppConsumer.Consumer = (*Consumer)(nil)

type Consumer struct {
	consumer weatherAppConsumer.Consumer
}

func New(consumer weatherAppConsumer.Consumer) *Consumer {
	return &Consumer{consumer: consumer}
}

func (c *Consumer) Consume(msg weatherapp.WeatherMsg) {
	c.consumer.Consume(msg)

	slog.Info("consumed msg for city", slog.Any("cityName", msg.CityName))
}
