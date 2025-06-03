package stageone

import (
	"github.com/sirupsen/logrus"
	"main.go/internal/weatherapp"
	weatherAppConsumer "main.go/internal/weatherapp/consumer"
	weatherAppProducer "main.go/internal/weatherapp/producer"
	"main.go/internal/weatherapp/runner/stage"
)

var _ stage.Runner = (*Runner)(nil)

type Runner struct {
	producer        weatherAppProducer.Producer
	consumer        weatherAppConsumer.Consumer
	shortCitiesInfo []weatherapp.ShortCityInfo
}

func NewRunner(
	producer weatherAppProducer.Producer,
	consumer weatherAppConsumer.Consumer,
	shortCitiesInfo []weatherapp.ShortCityInfo,
) *Runner {
	return &Runner{
		producer:        producer,
		consumer:        consumer,
		shortCitiesInfo: shortCitiesInfo,
	}
}

func (r *Runner) Run() error {
	for _, shortCityInfo := range r.shortCitiesInfo {
		msg, err := r.producer.Produce(shortCityInfo)
		if err != nil {
			logrus.Errorf("Error during producing msg: %v", err)

			continue
		}

		r.consumer.Consume(msg)
	}

	return nil
}
