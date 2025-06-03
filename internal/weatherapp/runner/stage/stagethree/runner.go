package stagethree

import (
	"sync"

	"main.go/internal/weatherapp"
	weatherAppConsumer "main.go/internal/weatherapp/consumer"
	weatherAppProducer "main.go/internal/weatherapp/producer"
	channelConsumer "main.go/internal/weatherapp/runner/consumer/channel"
	msgProducer "main.go/internal/weatherapp/runner/producer/msg/fromarray"
	"main.go/internal/weatherapp/runner/stage"
)

var _ stage.Runner = (*Runner)(nil)

type Runner struct {
	producer        weatherAppProducer.Producer
	consumer        weatherAppConsumer.Consumer
	shortCitiesInfo []weatherapp.ShortCityInfo
	consumerNumber  int
}

func NewRunner(
	producer weatherAppProducer.Producer,
	consumer weatherAppConsumer.Consumer,
	shortCitiesInfo []weatherapp.ShortCityInfo,
	consumerNumber int,
) *Runner {
	return &Runner{
		producer:        producer,
		consumer:        consumer,
		shortCitiesInfo: shortCitiesInfo,
		consumerNumber:  consumerNumber,
	}
}

func (r *Runner) Run() error {
	var wg sync.WaitGroup

	msgChannel := make(chan weatherapp.WeatherMsg)

	producerRunner := msgProducer.NewRunner(r.producer, msgChannel)
	consumerRunner := channelConsumer.NewRunner(r.consumer, msgChannel)

	wg.Add(r.consumerNumber)

	for i := 0; i < r.consumerNumber; i++ {
		go consumerRunner.Consume(&wg)
	}

	wg.Add(1)

	go producerRunner.Produce(&wg, r.shortCitiesInfo)
	wg.Wait()

	return nil
}
