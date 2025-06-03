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

	logrus.Infof("Produced msg for city: %v", msg.CityName)

	return msg, fmt.Errorf("error while producing weather msg: %w", err)
}
