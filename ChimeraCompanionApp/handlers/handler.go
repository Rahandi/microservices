package handlers

import (
	"ChimeraCompanionApp/types"
	"context"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Handler interface {
	Handle(ctx context.Context, input *tgbotapi.Message) (*tgbotapi.MessageConfig, error)
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

	accountId := strconv.FormatInt(input.Message.From.ID, 10)
	ctx := context.WithValue(context.Background(), types.AccountIdKey, accountId)

	for _, handler := range h.handlers {
		output, err = handler.Handle(ctx, input.Message)

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
