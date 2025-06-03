package fromarray

import (
	"sync"

	"github.com/sirupsen/logrus"
	"main.go/internal/weatherapp"
	weatherAppProducer "main.go/internal/weatherapp/producer"
)

type ProducerRunner struct {
	producer weatherAppProducer.Producer
	channel  chan<- weatherapp.WeatherMsg
}

func NewRunner(
	producer weatherAppProducer.Producer,
	channel chan<- weatherapp.WeatherMsg,
) *ProducerRunner {
	return &ProducerRunner{producer: producer, channel: channel}
}

func (p *ProducerRunner) Produce(
	wg *sync.WaitGroup,
	shortCitiesInfo []weatherapp.ShortCityInfo,
) {
	defer wg.Done()

	for _, shortCityInfo := range shortCitiesInfo {
		msg, err := p.producer.Produce(shortCityInfo)
		if err != nil {
			logrus.Errorf("error during producing msg: %v", err)

			continue
		}

		p.channel <- msg
	}

	close(p.channel)
}
