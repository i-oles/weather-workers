package log

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"main.go/internal/weatherapp"
	"main.go/internal/weatherapp/writer"
)

var _ writer.WeatherAppResultWriter = (*Writer)(nil)

type Writer struct {
	writer writer.WeatherAppResultWriter
}

func New(writer writer.WeatherAppResultWriter) Writer {
	return Writer{writer: writer}
}

func (w Writer) Write(result map[string]weatherapp.Result) error {
	err := w.writer.Write(result)
	if err != nil {
		return fmt.Errorf("failed writing result: %w", err)
	}

	logrus.Infof("Results: %+v", result)

	return nil
}
