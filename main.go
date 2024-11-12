package main

import (
	"moniteur/jobs"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	// bot.Init()
	jobs.Monitor()
}
