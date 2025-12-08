package openMeteo

type Forecast struct {
	Timezone            string              `json:"timezone"`
	Latitude            float64             `json:"latitude"`
	Longitude           float64             `json:"longitude"`
	Elevation           float64             `json:"elevation"`
	CurrentWeather      CurrentWeather      `json:"current_weather"`
	CurrentWeatherUnits CurrentWeatherUnits `json:"current_weather_units"`
}

type CurrentWeatherUnits struct {
	Time          string `json:"time"`
	Temperature   string `json:"temperature"`
	WindSpeed     string `json:"windspeed"`
	WindDirection string `json:"winddirection"`
	IsDay         string `json:"is_day"`
	WeatherCode   string `json:"weathercode"`
}

type CurrentWeather struct {
	Time          string  `json:"time"`
	Temperature   float32 `json:"temperature"`
	WindSpeed     float32 `json:"windspeed"`
	WindDirection float32 `json:"winddirection"`
	IsDay         uint8   `json:"is_day"`
	WeatherCode   uint8   `json:"weathercode"`
}
