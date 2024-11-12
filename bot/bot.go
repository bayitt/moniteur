package bot

import (
	"fmt"
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Init() *tgbotapi.BotAPI {
	fmt.Print(os.Getenv("BOT_API_TOKEN"))
	bot, err := tgbotapi.NewBotAPI(os.Getenv("BOT_API_TOKEN"))

	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	update := tgbotapi.NewUpdate(0)
	update.Timeout = 60

	updates := bot.GetUpdatesChan(update)

	for update := range updates {
		fmt.Printf(update.Message.Text)
		message := tgbotapi.NewMessage(update.Message.From.ID, "I am very very hungry")
		message.ReplyToMessageID = update.Message.MessageID
		bot.Send(message)
	}

	return bot
}
