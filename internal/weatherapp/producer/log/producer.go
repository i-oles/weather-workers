package log

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"main.go/internal/weatherapp"
	weatherAppProducer "main.go/internal/weatherapp/producer"
)

var _ weatherAppProducer.Producer = (*Producer)(nil)

type Producer struct {
	producer weatherAppProducer.Producer
}

func New(producer weatherAppProducer.Producer) *Producer {
	return &Producer{producer: producer}
}

func (p Producer) Produce(cityInfo weatherapp.ShortCityInfo) (weatherapp.WeatherMsg, error) {
	msg, err := p.producer.Produce(cityInfo)
	if err != nil {
		return weatherapp.WeatherMsg{}, fmt.Errorf("failed to produce weather msg: %w", err)
	}

	logrus.Infof("Produced msg for city: %v", msg.CityName)

	return msg, nil
}
