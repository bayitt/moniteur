package main

import (
	"log"
	"moniteur/bot"
	"moniteur/handlers"
	"moniteur/jobs"

	"github.com/go-co-op/gocron/v2"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	scheduler, err := gocron.NewScheduler()

	if err != nil {
		log.Panic(err)
	}

	_, jobErr := scheduler.NewJob(gocron.DailyJob(1, gocron.NewAtTimes(gocron.NewAtTime(10, 0, 0))), gocron.NewTask(jobs.Monitor))

	if jobErr != nil {
		log.Panic(err)
	}

	scheduler.Start()
	handlers.Init(bot.Init())
}
