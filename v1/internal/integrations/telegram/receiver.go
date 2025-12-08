package telegram

import (
	"encoding/json"
	"fmt"
	"log"
)

func Receive(body []byte) (*Update, error) {
	var update Update
	if err := json.Unmarshal(body, &update); err != nil {
		return nil, fmt.Errorf("telegram: invalid JSON: %w", err)
	}

	logData(update)

	return &update, nil
}

func logData(update Update) {
	prettyJSON, err := json.MarshalIndent(update, "", "  ")
	if err != nil {
		log.Printf("Failed to format JSON: %v", err)
	}
	log.Printf("Received:\n%s", string(prettyJSON))
}
