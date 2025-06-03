package mutexdecorator

import (
	"main.go/internal/weatherapp"
	"main.go/internal/weatherapp/presaver"
	"sync"
)

type PreSaver struct {
	preSaver presaver.WeatherAppPreSaver
	mutex    *sync.Mutex
}

func New(preSaver presaver.WeatherAppPreSaver) *PreSaver {
	return &PreSaver{
		preSaver: preSaver,
		mutex:    &sync.Mutex{},
	}
}

func (s *PreSaver) Save(cityResult weatherapp.CityResult) {
	s.mutex.Lock()
	s.preSaver.Save(cityResult)
	s.mutex.Unlock()
}

func (s *PreSaver) GetResult() weatherapp.Result {
	return s.preSaver.GetResult()
}

func (s *PreSaver) GetName() string {
	return s.preSaver.GetName()
}
