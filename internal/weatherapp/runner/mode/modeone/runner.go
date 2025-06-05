package modeone

import (
	"log/slog"

	"main.go/internal/weatherapp"
	weatherAppConsumer "main.go/internal/weatherapp/consumer"
	weatherAppProducer "main.go/internal/weatherapp/producer"
	"main.go/internal/weatherapp/runner/mode"
)

var _ mode.Runner = (*Runner)(nil)

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
			slog.Error("error during producing msg", slog.Any("error", err))

			continue
		}

		r.consumer.Consume(msg)
	}

	return nil
}
