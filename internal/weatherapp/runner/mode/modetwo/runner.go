package modetwo

import (
	"sync"

	"main.go/internal/weatherapp"
	weatherAppConsumer "main.go/internal/weatherapp/consumer"
	weatherAppProducer "main.go/internal/weatherapp/producer"
	channelConsumer "main.go/internal/weatherapp/runner/consumer/channel"
	"main.go/internal/weatherapp/runner/mode"
	msgProducer "main.go/internal/weatherapp/runner/producer/msg/fromarray"
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
	var wg sync.WaitGroup

	msgChannel := make(chan weatherapp.WeatherMsg)

	producerRunner := msgProducer.NewRunner(r.producer, msgChannel)
	consumerRunner := channelConsumer.NewRunner(r.consumer, msgChannel)

	wg.Add(1)

	go consumerRunner.Consume(&wg)

	wg.Add(1)

	go producerRunner.Produce(&wg, r.shortCitiesInfo)

	wg.Wait()

	return nil
}
