package channel

import (
	"sync"

	"main.go/internal/weatherapp"
	weatherAppConsumer "main.go/internal/weatherapp/consumer"
)

type ConsumerRunner struct {
	consumer weatherAppConsumer.Consumer
	channel  <-chan weatherapp.WeatherMsg
}

func NewRunner(
	consumer weatherAppConsumer.Consumer,
	channel <-chan weatherapp.WeatherMsg,
) *ConsumerRunner {
	return &ConsumerRunner{consumer: consumer, channel: channel}
}

func (c *ConsumerRunner) Consume(wq *sync.WaitGroup) {
	defer wq.Done()

	for msg := range c.channel {
		c.consumer.Consume(msg)
	}
}
