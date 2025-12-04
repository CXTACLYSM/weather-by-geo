package main

import (
	"io"
	"log"
	"net/http"

	"github.com/CXTACLYSM/weather-by-geo/configs"
	"github.com/CXTACLYSM/weather-by-geo/internal/shared/infrastructure/telegram"
	"github.com/CXTACLYSM/weather-by-geo/pkg/connector/clickhouse"
	"github.com/CXTACLYSM/weather-by-geo/pkg/connector/postgres"
)

func main() {
	cfg, err := configs.Create()
	if err != nil {
		log.Fatalf("Failed to load configs: %v", err)
	}

	pgConn, err := postgres.NewConnector(cfg.Postgres)
	if err != nil {
		log.Fatalf("Failed to connect to PostgreSQL: %v", err)
	}
	defer pgConn.Close()

	chConn, err := clickhouse.NewConnector(cfg.ClickHouse)
	if err != nil {
		log.Fatalf("Failed to connect to ClickHouse: %v", err)
	}
	defer chConn.Close()

	tlClient := telegram.NewClient(cfg)
	tlClient.SetWebhook()

	log.Println("Successfully posted webhook URL")

	tlReceiver := telegram.NewReceiver(cfg)
	tlHandler := telegram.NewHandler(cfg)

	http.HandleFunc("/", webhookHandler(tlReceiver, tlHandler, tlClient))

	log.Printf("Starting server on port 8080")
	if err := http.ListenAndServe("0.0.0.0:8080", nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func webhookHandler(receiver *telegram.Receiver, handler *telegram.Handler, client *telegram.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Cannot read body", http.StatusInternalServerError)
			return
		}

		update, err := receiver.Receive(body)
		if err != nil {
			http.Error(w, "Cannot parse body", http.StatusInternalServerError)
			return
		}

		response, err := handler.Handle(update)
		if err != nil {
			http.Error(w, "Cannot handle update", http.StatusInternalServerError)
			return
		}

		err = client.SendMessage(response)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
	}
}
