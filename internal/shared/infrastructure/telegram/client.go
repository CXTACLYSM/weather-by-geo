package telegram

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/CXTACLYSM/weather-by-geo/configs"
)

const UrlPattern = "https://%s/bot%s/%s"

const (
	SendMessage = "sendMessage"
	SetWebhook  = "setWebhook"
)

var allowedOperations = map[string]bool{
	SendMessage: true,
	SetWebhook:  true,
}

type Client struct {
	cfg *configs.Config
}

func NewClient(cfg *configs.Config) *Client {
	return &Client{
		cfg: cfg,
	}
}

func (c *Client) SendMessage(response *Response) error {
	url, err := c.url(SendMessage)
	if err != nil {
		return errors.New("failed to send message")
	}

	jsonData, err := json.Marshal(response)
	if err != nil {
		return errors.New("failed to marshal response")
	}

	res, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return errors.New("failed to send message")
	}
	defer res.Body.Close()

	return nil
}

func (c *Client) SetWebhook() {
	url, err := c.url(SetWebhook)
	if err != nil {
		log.Fatalf("Failed to get telegram setWebhook URL: %v", err)
	}
	data := map[string]interface{}{
		"url": c.cfg.App.Url("https"),
	}
	log.Printf("Webhook URL: %s\n", data["url"])

	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Fatalf("Failed to marshal data: %v", err)
	}
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatalf("Failed to post data: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		log.Fatalf("Telegram API error (status %d): %s", resp.StatusCode, string(body))
	}
}

func (c *Client) url(operation string) (string, error) {
	if _, ok := allowedOperations[operation]; !ok {
		return "", fmt.Errorf("operation %q is not supported", operation)
	}
	return fmt.Sprintf(UrlPattern, c.cfg.Telegram.Host, c.cfg.Telegram.Token, operation), nil
}
