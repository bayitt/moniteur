package handlers

import (
	"fmt"
	"log"
	"net/http"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Init(bot *tgbotapi.BotAPI) {
	var updates tgbotapi.UpdatesChannel

	fmt.Println(os.Getenv("APP_ENV"))

	if os.Getenv("APP_ENV") != "prod" {
		update := tgbotapi.NewUpdate(0)
		update.Timeout = 60

		updates = bot.GetUpdatesChan(update)
	} else {
		webhook, _ := tgbotapi.NewWebhookWithCert(os.Getenv("APP_URL")+"/"+bot.Token, tgbotapi.FilePath("cert.pem"))

		_, err := bot.Request(webhook)

		if err != nil {
			log.Fatal(err)
		}

		info, err := bot.GetWebhookInfo()

		if err != nil {
			log.Fatal(err)
		}

		if info.LastErrorDate != 0 {
			log.Printf("Telegram callback failed: %s", info.LastErrorMessage)
		}

		updates = bot.ListenForWebhook("/" + bot.Token)
		go http.ListenAndServe("0.0.0.0:8080", nil)
	}

	for update := range updates {
		if update.CallbackQuery != nil {
			Callbacks(bot, update)
			continue
		}

		if update.Message.IsCommand() {
			Commands(bot, update)
			continue
		}

		message := tgbotapi.NewMessage(update.Message.From.ID, "I am sorry. I do not understand that.")
		message.ReplyToMessageID = update.Message.MessageID
		bot.Send(message)
	}
}
