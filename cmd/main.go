package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/CXTACLYSM/weather-by-geo/common/connector/clickhouse"
	"github.com/CXTACLYSM/weather-by-geo/common/connector/postgres"
	"github.com/CXTACLYSM/weather-by-geo/config"
)

func main() {
	cfg, err := config.Create()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
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

	setupWebhook(cfg)
	log.Println("Successfully posted webhook URL")

	setupHttp()
	log.Printf("Starting server on port %s", cfg.App.Port)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", cfg.App.Port), nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func setupWebhook(cfg *config.Config) {
	url := fmt.Sprintf("https://%s/bot%s/setWebhook", cfg.Telegram.Host, cfg.Telegram.Token)
	data := map[string]interface{}{
		"url": cfg.Proxy.Url("https"),
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

func setupHttp() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Cannot read body", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		var data map[string]interface{}
		if err := json.Unmarshal(body, &data); err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		prettyJSON, err := json.MarshalIndent(data, "", "  ")
		if err != nil {
			log.Printf("Failed to format JSON: %v", err)
			return
		}
		log.Printf("Received:\n%s", string(prettyJSON))

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
	})
}
