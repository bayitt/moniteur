package bot

import (
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Init() *tgbotapi.BotAPI {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("BOT_API_TOKEN"))

	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	return bot
}
