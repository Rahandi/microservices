package handlers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Handler interface {
	Handle(input *tgbotapi.Message) (*tgbotapi.MessageConfig, error)
}

type MainHandler struct {
	bot      *tgbotapi.BotAPI
	handlers []Handler
}

func NewMainHandler(bot *tgbotapi.BotAPI, handlers []Handler) *MainHandler {
	return &MainHandler{
		bot:      bot,
		handlers: handlers,
	}
}

func (h *MainHandler) Handle(input *tgbotapi.Update) error {
	var output *tgbotapi.MessageConfig
	var err error

	for _, handler := range h.handlers {
		output, err = handler.Handle(input.Message)

		if err != nil {
			outputMsg := tgbotapi.NewMessage(input.Message.Chat.ID, err.Error())
			output = &outputMsg
		}

		if output == nil {
			continue
		}

		_, err = h.bot.Send(output)
		if err != nil {
			return err
		}
	}

	return nil
}
