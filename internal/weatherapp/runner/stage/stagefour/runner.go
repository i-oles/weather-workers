package stagefour

import (
	"sync"

	"main.go/internal/weatherapp"
	weatherAppConsumer "main.go/internal/weatherapp/consumer"
	weatherAppProducer "main.go/internal/weatherapp/producer"
	channelConsumer "main.go/internal/weatherapp/runner/consumer/channel"
	msgProducer "main.go/internal/weatherapp/runner/producer/msg/fromchannel"
	"main.go/internal/weatherapp/runner/stage"
)

var _ stage.Runner = (*Runner)(nil)

type Runner struct {
	producer        weatherAppProducer.Producer
	consumer        weatherAppConsumer.Consumer
	shortCitiesInfo []weatherapp.ShortCityInfo
	consumerNumber  int
	producerNumber  int
}

func NewRunner(
	producer weatherAppProducer.Producer,
	consumer weatherAppConsumer.Consumer,
	shortCitiesInfo []weatherapp.ShortCityInfo,
	consumerNumber int,
	producerNumber int,
) *Runner {
	return &Runner{
		producer:        producer,
		consumer:        consumer,
		shortCitiesInfo: shortCitiesInfo,
		consumerNumber:  consumerNumber,
		producerNumber:  producerNumber,
	}
}

func (r *Runner) Run() error {
	var consumerWg sync.WaitGroup

	var producerWg sync.WaitGroup

	msgChannel := make(chan weatherapp.WeatherMsg)
	shortCityInfoChannel := make(chan weatherapp.ShortCityInfo)

	producerRunner := msgProducer.NewRunner(
		r.producer, shortCityInfoChannel, msgChannel,
	)

	consumerRunner := channelConsumer.NewRunner(r.consumer, msgChannel)

	go func() {
		for _, cityInfo := range r.shortCitiesInfo {
			shortCityInfoChannel <- cityInfo
		}

		close(shortCityInfoChannel)
	}()

	consumerWg.Add(r.consumerNumber)

	for i := 0; i < r.consumerNumber; i++ {
		go consumerRunner.Consume(&consumerWg)
	}

	producerWg.Add(r.producerNumber)

	for i := 0; i < r.producerNumber; i++ {
		go producerRunner.Produce(&producerWg)
	}

	producerWg.Wait()
	close(msgChannel)
	consumerWg.Wait()

	return nil
}
