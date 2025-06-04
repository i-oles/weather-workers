package writer

import (
	"main.go/internal/weatherapp"
)

type WeatherAppResultWriter interface {
	Write(result map[string]weatherapp.Result) error
}

type PerformanceTestWriter interface {
	Write(executionDuration string) error
}
