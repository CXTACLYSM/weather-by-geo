package telegram

import (
	"errors"
	"fmt"
	"time"

	openMeteo2 "github.com/CXTACLYSM/weather-by-geo/internal/integrations/openMeteo"
)

func Handle(update *Update, weatherClient *openMeteo2.Client) (*Response, error) {
	if update.Message == nil {
		return nil, errors.New("unexpected telegram error: message is nil")
	}
	if update.Message.Location == nil {
		return &Response{
			ChatID: update.Message.Chat.ID,
			Text:   "send your location",
		}, nil
	}

	forecast, err := weatherClient.GetForecast(update.Message.Location.Latitude, update.Message.Location.Longitude)
	if err != nil {
		return nil, err
	}

	return &Response{
		ChatID: update.Message.Chat.ID,
		Text:   formatForecast(forecast),
	}, nil
}

func formatForecast(forecast *openMeteo2.Forecast) string {
	return fmt.Sprintf("Time (%s): %s\nLatitude: %.5f\nLongitude: %.5f\nTemperature: %.1f %s\nElevation: %.0f m\nWindspeed: %.1f %s\nWind Direction: %.0f %s",
		forecast.Timezone,
		formatTime(forecast.CurrentWeather.Time),
		forecast.Latitude, forecast.Longitude,
		forecast.CurrentWeather.Temperature,
		forecast.CurrentWeatherUnits.Temperature,
		forecast.Elevation,
		forecast.CurrentWeather.WindSpeed,
		forecast.CurrentWeatherUnits.WindSpeed,
		forecast.CurrentWeather.WindDirection,
		forecast.CurrentWeatherUnits.WindDirection,
	)
}

func formatTime(raw string) string {
	t, err := time.Parse("2006-01-02T15:04", raw)
	if err != nil {
		return raw
	}
	return t.Format("02.01.2006 15:04")
}
