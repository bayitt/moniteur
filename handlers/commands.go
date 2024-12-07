package handlers

import (
	"encoding/json"
	"moniteur/jobs"
	"os"
	"path/filepath"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Commands(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	if update.Message.Command() == "start" {
		StartCommand(bot, update)
		return
	}

	message := tgbotapi.NewMessage(update.Message.From.ID, "I am sorry. I do not understand that.")
	message.ReplyToMessageID = update.Message.MessageID
	bot.Send(message)
}

func StartCommand(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	directory, _ := os.Getwd()
	servicePath := filepath.Join(directory, "services.json")
	serviceContent, _ := os.ReadFile(servicePath)

	var serviceData jobs.ServiceData
	json.Unmarshal(serviceContent, &serviceData)

	services := serviceData.Services
	keyboardRows := [][]tgbotapi.InlineKeyboardButton{}
	serviceLength := len(services)

	var isLengthOdd bool = false

	if serviceLength%2 == 1 {
		serviceLength += 2
		isLengthOdd = true
	}

	for i := 2; i <= serviceLength; i += 2 {
		var keyboardRow []tgbotapi.InlineKeyboardButton
		length := serviceLength

		if isLengthOdd {
			length = serviceLength - 2
		}

		if i > length {
			keyboardRow = tgbotapi.NewInlineKeyboardRow(tgbotapi.NewInlineKeyboardButtonData(services[i-2].Name, "check_service/"+services[i-2].Name))
		} else {
			keyboardRow = tgbotapi.NewInlineKeyboardRow(tgbotapi.NewInlineKeyboardButtonData(services[i-2].Name, "check_service/"+services[i-2].Name), tgbotapi.NewInlineKeyboardButtonData(services[i-1].Name, "check_service/"+services[i-1].Name))
		}
		keyboardRows = append(keyboardRows, keyboardRow)
	}

	keyboard := tgbotapi.NewInlineKeyboardMarkup(keyboardRows...)

	message := tgbotapi.NewMessage(update.Message.From.ID, "Hi, what application do you want to check ?")
	message.ReplyMarkup = keyboard
	bot.Send(message)
}
