package result

import (
	"encoding/json"
	"fmt"
	"os"

	"main.go/internal/weatherapp"
	"main.go/internal/weatherapp/writer"
)

var _ writer.WeatherAppResultWriter = (*Writer)(nil)

type Writer struct {
	resultsFile *os.File
}

func NewWriter(resultsFile *os.File) Writer {
	return Writer{resultsFile: resultsFile}
}

func (w Writer) Write(result map[string]weatherapp.Result) error {
	byteResult, err := json.Marshal(result)
	if err != nil {
		return fmt.Errorf("failed marshaling result: %w", err)
	}

	_, err = w.resultsFile.Write(byteResult)
	if err != nil {
		return fmt.Errorf("failed writing result to file: %w", err)
	}

	return nil
}
