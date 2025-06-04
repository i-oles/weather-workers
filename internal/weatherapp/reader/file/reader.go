package file

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"main.go/internal/weatherapp"
)

type Reader struct {
}

func New() Reader {
	return Reader{}
}

func (r Reader) Read(filePath string) ([]weatherapp.CityInfo, error) {
	sourceFile, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}

	byteSource, err := io.ReadAll(sourceFile)
	if err != nil {
		return nil, fmt.Errorf("failed while io.ReadAll: %w", err)
	}

	var citiesInfo []weatherapp.CityInfo

	err = json.Unmarshal(byteSource, &citiesInfo)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal file: %w", err)
	}

	return citiesInfo, nil
}
