package openMeteo

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"

	"github.com/CXTACLYSM/weather-by-geo/configs/integrations/openMeteo"
)

const OperationForecast = "forecast"

var allowedOperations = map[string]bool{
	OperationForecast: true,
}

type Client struct {
	cfg openMeteo.Config
}

func NewClient(cfg openMeteo.Config) *Client {
	return &Client{
		cfg: cfg,
	}
}

func (c *Client) GetForecast(latitude float64, longitude float64) (*Forecast, error) {
	forecastUrl, err := c.url(OperationForecast, latitude, longitude)
	if err != nil {
		return nil, fmt.Errorf("error building forecastUrl: %w", err)
	}

	resp, err := http.Get(forecastUrl)
	if err != nil {
		return nil, fmt.Errorf("error fetching data from %s: %w", forecastUrl, err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body from forecastUrl %s: %w", forecastUrl, err)
	}

	var forecast Forecast
	err = json.Unmarshal(body, &forecast)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling json: %w", err)
	}

	return &forecast, nil
}

func (c *Client) url(operation string, latitude float64, longitude float64) (string, error) {
	if _, ok := allowedOperations[operation]; !ok {
		return "", fmt.Errorf("operation %q is not supported", operation)
	}

	u := &url.URL{
		Scheme: "https",
		Host:   c.cfg.Host,
		Path:   fmt.Sprintf("/v%s/%s", c.cfg.Version, operation),
	}

	q := u.Query()
	q.Set("latitude", fmt.Sprintf("%.5f", latitude))
	q.Set("longitude", fmt.Sprintf("%.5f", longitude))
	q.Set("current_weather", "true")
	q.Set("timezone", "auto")

	u.RawQuery = q.Encode()

	log.Printf("Forecast api url: %s", u.String())

	return u.String(), nil
}
