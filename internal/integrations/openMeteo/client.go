package openMeteo

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/CXTACLYSM/weather-by-geo/configs/integrations/openMeteo"
)

const UrlPattern = "https://%s/v%s/%s"

const OperationForecast = "forecast"

type Client struct {
	cfg openMeteo.Config
}

func NewClient(cfg openMeteo.Config) *Client {
	return &Client{
		cfg: cfg,
	}
}

func (c *Client) GetForecast(latitude float64, longitude float64) (*Forecast, error) {
	url, err := c.url(OperationForecast, latitude, longitude)
	if err != nil {
		return nil, fmt.Errorf("error building url: %w", err)
	}

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error fetching data from %s: %w", url, err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body from url %s: %w", url, err)
	}

	var forecast Forecast
	err = json.Unmarshal(body, &forecast)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling json: %w", err)
	}

	return &forecast, nil
}

func (c *Client) url(operation string, latitude float64, longitude float64) (string, error) {
	switch operation {
	case OperationForecast:
		url := fmt.Sprintf(UrlPattern, c.cfg.Host, c.cfg.Version, operation)
		url = fmt.Sprintf("%s?latitude=%f&longitude=%f&timezone=auto&current_weather=true", url, latitude, longitude)
		fmt.Println("sending url:", url)
		return url, nil
	default:
		return "", errors.New("unknown operation")
	}
}
