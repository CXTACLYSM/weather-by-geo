package main

import (
	"log"

	"github.com/CXTACLYSM/weather-by-geo/configs"
	"github.com/CXTACLYSM/weather-by-geo/internal/handlers/telegram"
	"github.com/CXTACLYSM/weather-by-geo/internal/integrations/openMeteo"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	cfg, err := configs.Create()
	if err != nil {
		log.Fatalf("Failed to load configs: %v", err)
	}

	weatherClient := openMeteo.NewClient(cfg.OpenMeteo)

	bot, err := tgbotapi.NewBotAPI(cfg.Telegram.Token)
	if err != nil {
		log.Fatalf("Failed to create bot: %v", err)
	}
	log.Printf("Authorized on account %s", bot.Self.UserName)

	switch true {
	case cfg.Telegram.IsWebhook:
		telegram.WebhookHandler(cfg, bot, weatherClient)
	case !cfg.Telegram.IsWebhook:
		telegram.UpdateHandler(bot, weatherClient)
	}
}
