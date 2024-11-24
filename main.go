package main

import (
	"moniteur/bot"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	bot.Init()
	// jobs.Monitor()

	// fmt.Printf("Issuer: %s\nExpiry:%v", issuer, expiry.Format(time.RFC850))
}
