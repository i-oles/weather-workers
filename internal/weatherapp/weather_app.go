package weatherapp

type WeatherCode int

type WeatherStats struct {
	Hourly struct {
		WeatherCodes  []WeatherCode `json:"weathercode"`
		Temperature2m []float64     `json:"temperature_2m"`
	}
}

type WeatherMsg struct {
	CityName     string
	WeatherCodes []WeatherCode
	Temperatures []float64
}

type Result struct {
	CityName string `json:"city_name,omitempty"`
	Value    any    `json:"value,omitempty"`
}
