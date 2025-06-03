package weatherapp

type CityInfo struct {
	City             string `json:"city,omitempty"`
	Lat              string `json:"lat,omitempty"`
	Lng              string `json:"lng,omitempty"`
	Country          string `json:"country,omitempty"`
	Iso2             string `json:"iso2,omitempty"`
	AdminName        string `json:"admin_name,omitempty"`
	Capital          string `json:"capital,omitempty"`
	Population       string `json:"population,omitempty"`
	PopulationProper string `json:"population_proper,omitempty"`
}

type ShortCityInfo struct {
	Name      string `json:"city"`
	Latitude  string `json:"lat"`
	Longitude string `json:"lng"`
}

type CityResult struct {
	Name            string
	FoggyHoursCount int
	TempAverage     float64
	SunnyHoursCount int
}

func GetShortCityInfo(city CityInfo) ShortCityInfo {
	return ShortCityInfo{
		Name:      city.City,
		Latitude:  city.Lat,
		Longitude: city.Lng,
	}
}
