package main

import (
	"ChimeraCompanionApp/handlers"
	"ChimeraCompanionApp/internals"
	"ChimeraCompanionApp/services"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	config := internals.NewConfig()
	redis := internals.NewRedis(config)
	bot, err := tgbotapi.NewBotAPI(config.Token)
	if err != nil {
		panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	iamService := services.NewIAMService(config, redis)
	iamHandler := handlers.NewIAMHandler(iamService, redis)

	ListHandlers := []handlers.Handler{
		iamHandler,
	}
	handler := handlers.NewMainHandler(bot, ListHandlers)

	for update := range updates {
		if update.Message != nil { // If we got a message
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
			handler.Handle(&update)
		}
	}
}
