package fromchannel

import (
	"sync"

	"github.com/sirupsen/logrus"
	"main.go/internal/weatherapp"
	weatherAppProducer "main.go/internal/weatherapp/producer"
)

type ProducerRunner struct {
	producer             weatherAppProducer.Producer
	msgChannel           chan weatherapp.WeatherMsg
	shortCityInfoChannel chan weatherapp.ShortCityInfo
}

func NewRunner(
	producer weatherAppProducer.Producer,
	shortCityInfoChannel chan weatherapp.ShortCityInfo,
	msgChannel chan weatherapp.WeatherMsg,
) *ProducerRunner {
	return &ProducerRunner{
		producer:             producer,
		shortCityInfoChannel: shortCityInfoChannel,
		msgChannel:           msgChannel,
	}
}

func (p *ProducerRunner) Produce(
	wg *sync.WaitGroup,
) {
	defer wg.Done()

	for shortCityInfo := range p.shortCityInfoChannel {
		msg, err := p.producer.Produce(shortCityInfo)
		if err != nil {
			logrus.Errorf("error during producing msg: %v", err)

			continue
		}

		p.msgChannel <- msg
	}
}
