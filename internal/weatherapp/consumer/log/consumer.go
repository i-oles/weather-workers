package log

import (
	"github.com/sirupsen/logrus"
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

	logrus.Infof("Consumed msg for city: %v", msg.CityName)
}
