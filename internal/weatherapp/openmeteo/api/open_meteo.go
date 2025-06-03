package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"main.go/internal/weatherapp"
	"main.go/internal/weatherapp/openmeteo"
)

const (
	dateFormat = "2006-01-02"
)

var _ openmeteo.API = (*OpenMeteo)(nil)

type OpenMeteo struct {
	apiURL    string
	startDate string
	endDate   string
}

func NewOpenMeteo(apiURL string, analysisDurationInMonths int) OpenMeteo {
	return OpenMeteo{
		apiURL: apiURL,
		startDate: time.Now().AddDate(
			0, -analysisDurationInMonths, 0).Format(dateFormat),
		endDate: time.Now().Format(dateFormat),
	}
}

func (p OpenMeteo) GetWeather(
	latitude, longitude, weatherTag string,
) (weatherapp.WeatherStats, error) {
	url := fmt.Sprintf("%s?latitude=%s&longitude=%s&start_date=%s&end_date=%s&hourly=%s",
		p.apiURL, latitude, longitude, p.startDate, p.endDate, weatherTag)

	resp, err := http.Get(url)
	if err != nil {
		return weatherapp.WeatherStats{}, fmt.Errorf("error while getting weather stats: %w", err)
	}

	result, err := io.ReadAll(resp.Body)
	if err != nil {
		return weatherapp.WeatherStats{}, fmt.Errorf("error while reading weather response body: %w", err)
	}

	var weatherStats weatherapp.WeatherStats

	err = json.Unmarshal(result, &weatherStats)
	if err != nil {
		return weatherapp.WeatherStats{}, fmt.Errorf("error while unmarshaling weather stats: %w", err)
	}

	return weatherStats, nil
}
