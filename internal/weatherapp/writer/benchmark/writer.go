package benchmark

import (
	"fmt"
	"os"

	"main.go/internal/weatherapp/writer"
)

var _ writer.PerformanceTestWriter = (*Writer)(nil)

type Writer struct {
	resultsFile *os.File
}

func NewWriter(resultsFile *os.File) Writer {
	return Writer{
		resultsFile: resultsFile,
	}
}

func (w Writer) Write(executionDuration string) error {
	_, err := w.resultsFile.WriteString(executionDuration)
	if err != nil {
		return fmt.Errorf("error writing to results file: %w", err)
	}

	return nil
}
