package telegram

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/CXTACLYSM/weather-by-geo/configs"
	forecastFormatter "github.com/CXTACLYSM/weather-by-geo/internal/formatters/forecast"
	"github.com/CXTACLYSM/weather-by-geo/internal/integrations/openMeteo"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func UpdateHandler(bot *tgbotapi.BotAPI, weatherClient *openMeteo.Client) {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, _ := bot.GetUpdatesChan(u)

	bot.RemoveWebhook()
	for update := range updates {
		handleUpdate(update, weatherClient, bot)
	}
}

func WebhookHandler(cfg *configs.Config, bot *tgbotapi.BotAPI, weatherClient *openMeteo.Client) {
	log.Printf("Authorized on account %s", bot.Self.UserName)

	webhookUrl := fmt.Sprintf("https://%s:%s", cfg.App.Host, cfg.App.Port)
	_, err := bot.SetWebhook(tgbotapi.NewWebhook(webhookUrl))
	if err != nil {
		log.Fatalf("Failed to set webhook: %v", err)
	}
	info, err := bot.GetWebhookInfo()
	if err != nil {
		log.Fatalf("Failed to get webhook info: %v", err)
	}
	if info.LastErrorDate != 0 {
		log.Printf("Telegram callback failed: %s", info.LastErrorMessage)
	}
	updates := bot.ListenForWebhook("/")

	socketStr := fmt.Sprintf("0.0.0.0:%s", cfg.App.Port)
	go func() {
		err := http.ListenAndServe(socketStr, nil)
		if err != nil {
			log.Fatalf("Failed to listen and serve: %v", err)
		}
	}()

	for update := range updates {
		handleUpdate(update, weatherClient, bot)
	}
}

func handleUpdate(update tgbotapi.Update, weatherClient *openMeteo.Client, bot *tgbotapi.BotAPI) {
	text, _ := getReplyText(update, weatherClient)
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)
	msg.ReplyToMessageID = update.Message.MessageID

	bot.Send(msg)
}

func getReplyText(update tgbotapi.Update, weatherClient *openMeteo.Client) (string, error) {
	if update.Message == nil {
		return "", errors.New("unexpected telegram error: message is nil")
	}
	if update.Message.Location == nil {
		return "send your location", nil
	}

	forecast, err := weatherClient.GetForecast(update.Message.Location.Latitude, update.Message.Location.Longitude)
	if err != nil {
		return "", err
	}

	return forecastFormatter.FormatForecast(forecast), nil
}
