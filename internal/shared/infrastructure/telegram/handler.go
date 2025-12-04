package telegram

import (
	"errors"

	"github.com/CXTACLYSM/weather-by-geo/configs"
)

type Handler struct {
	cfg *configs.Config
}

func NewHandler(cfg *configs.Config) *Handler {
	return &Handler{
		cfg: cfg,
	}
}

func (rec *Handler) Handle(update *Update) (*Response, error) {
	var message string

	if update.Message == nil {
		return nil, errors.New("unexpected telegram error")
	}
	if update.Message.Location == nil {
		message = "send your location"
	} else {
		message = "location received"
	}

	return &Response{
		ChatID: update.Message.Chat.ID,
		Text:   message,
	}, nil
}
