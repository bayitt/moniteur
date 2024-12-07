package handlers

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func Init(bot *tgbotapi.BotAPI) {
	update := tgbotapi.NewUpdate(0)
	update.Timeout = 60

	updates := bot.GetUpdatesChan(update)

	for update := range updates {
		if update.CallbackQuery != nil {
			Callbacks(bot, update)
			return
		}

		if update.Message.IsCommand() {
			Commands(bot, update)
			return
		}

		message := tgbotapi.NewMessage(update.Message.From.ID, "I am sorry. I do not understand that.")
		message.ReplyToMessageID = update.Message.MessageID
		bot.Send(message)
	}
}
