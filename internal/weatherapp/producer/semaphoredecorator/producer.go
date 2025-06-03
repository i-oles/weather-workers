package semaphoredecorator

import (
	"context"
	"fmt"

	"golang.org/x/sync/semaphore"
	"main.go/internal/weatherapp"
	weatherAppProducer "main.go/internal/weatherapp/producer"
)

var _ weatherAppProducer.Producer = (*Producer)(nil)

type Producer struct {
	producer  weatherAppProducer.Producer
	semaphore *semaphore.Weighted
}

func New(producer weatherAppProducer.Producer, semaphore *semaphore.Weighted) *Producer {
	return &Producer{producer: producer, semaphore: semaphore}
}

func (p Producer) Produce(cityInfo weatherapp.ShortCityInfo) (weatherapp.WeatherMsg, error) {
	err := p.semaphore.Acquire(context.Background(), 1)
	if err != nil {
		return weatherapp.WeatherMsg{}, fmt.Errorf("error while acquiring semaphore: %w", err)
	}

	msg, err := p.producer.Produce(cityInfo)

	p.semaphore.Release(1)

	return msg, fmt.Errorf("error while producing weather msg: %w", err)
}
