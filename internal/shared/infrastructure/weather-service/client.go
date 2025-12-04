package weather_service

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const Host = ""
const OperationGetWeather = ""

type Client struct {
}

func NewClient() *Client {
	return &Client{}
}

func (c *Client) getWeather() (*WeatherResponse, error) {
	url := c.url("https", OperationGetWeather)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var weatherResponse WeatherResponse
	err = json.Unmarshal(body, &weatherResponse)
	if err != nil {
		return nil, err
	}

	return &weatherResponse, nil
}

func (c *Client) url(scheme string, operation string) string {
	return fmt.Sprintf("%s://%s/%s", scheme, Host, operation)
}
