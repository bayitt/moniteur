package main

import (
	"moniteur/bot"
	"moniteur/handlers"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	handlers.Init(bot.Init())

	// fmt.Printf("Issuer: %s\nExpiry:%v", issuer, expiry.Format(time.RFC850))
}
