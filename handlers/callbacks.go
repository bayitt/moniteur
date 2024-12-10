package handlers

import (
	"encoding/json"
	"moniteur/jobs"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"sync"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Callbacks(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	splits := strings.Split(update.CallbackQuery.Data, "/")
	chatId := update.CallbackQuery.From.ID
	command := splits[0]

	if command != "check_service" {
		message := tgbotapi.NewMessage(chatId, "I am sorry. I do not understand that.")
		bot.Send(message)
		return
	}

	directory, _ := os.Getwd()
	servicePath := filepath.Join(directory, "services.json")
	serviceContent, _ := os.ReadFile(servicePath)

	var serviceData jobs.ServiceData
	json.Unmarshal(serviceContent, &serviceData)

	services := serviceData.Services

	index := slices.IndexFunc(services, func(service jobs.Service) bool { return service.Name == splits[1] })

	if index == -1 {
		message := tgbotapi.NewMessage(chatId, "I am sorry. I do not recognize that application.")
		bot.Send(message)
		return
	}

	service := services[index]

	var wg sync.WaitGroup
	wg.Add(1)

	channel := make(chan jobs.Service, 1)
	jobs.PingService(service, &wg, channel)

	wg.Wait()
	close(channel)

	chatIds := []int{int(chatId)}

	jobs.Alert(channel, chatIds)
}
