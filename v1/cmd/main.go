package main

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/CXTACLYSM/weather-by-geo/configs"
	"github.com/CXTACLYSM/weather-by-geo/internal/integrations/openMeteo"
	"github.com/CXTACLYSM/weather-by-geo/internal/integrations/telegram"
)

func main() {
	cfg, err := configs.Create()
	if err != nil {
		log.Fatalf("Failed to load configs: %v", err)
	}

	tlClient := telegram.NewClient(cfg)
	tlClient.SetWebhook()
	log.Println("Successfully posted webhook URL")

	weatherClient := openMeteo.NewClient(cfg.OpenMeteo)
	http.HandleFunc("/", webhookHandler(tlClient, weatherClient))

	socketStr := fmt.Sprintf("0.0.0.0:%s", cfg.App.Port)

	log.Printf("Starting server on port 8080")
	if err := http.ListenAndServe(socketStr, nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func webhookHandler(telegramClient *telegram.Client, weatherClient *openMeteo.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		body, err := io.ReadAll(r.Body)
		if err != nil {
			log.Printf("Failed to read body: %v", err)
			http.Error(w, "Cannot read body", http.StatusInternalServerError)
			return
		}

		update, err := telegram.Receive(body)
		if err != nil {
			log.Printf("Failed to parse body: %v", err)
			http.Error(w, "Cannot parse body", http.StatusInternalServerError)
			return
		}

		response, err := telegram.Handle(update, weatherClient)
		if err != nil {
			log.Printf("Failed to handle request: %v", err)
			http.Error(w, "Cannot handle update", http.StatusInternalServerError)
			return
		}

		err = telegramClient.SendMessage(response)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}
