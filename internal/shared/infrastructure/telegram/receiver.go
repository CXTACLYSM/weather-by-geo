package telegram

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/CXTACLYSM/weather-by-geo/configs"
)

type Receiver struct {
	cfg *configs.Config
}

func NewReceiver(cfg *configs.Config) *Receiver {
	return &Receiver{
		cfg: cfg,
	}
}

func (rec *Receiver) Receive(body []byte) (*Update, error) {
	var update Update
	if err := json.Unmarshal(body, &update); err != nil {
		return nil, fmt.Errorf("telegram: invalid JSON: %w", err)
	}

	rec.log(update)

	return &update, nil
}

func (rec *Receiver) log(update Update) {
	prettyJSON, err := json.MarshalIndent(update, "", "  ")
	if err != nil {
		log.Printf("Failed to format JSON: %v", err)
	}
	log.Printf("Received:\n%s", string(prettyJSON))
}
