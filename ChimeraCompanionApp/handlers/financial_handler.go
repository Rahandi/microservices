package handlers

import (
	"ChimeraCompanionApp/internals"
	"ChimeraCompanionApp/models"
	"ChimeraCompanionApp/services"
	"context"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type FinancialHandler struct {
	redis            *internals.Redis
	financialService *services.FinancialService
}

func NewFinancialHandler(financialService *services.FinancialService, redis *internals.Redis) *FinancialHandler {
	return &FinancialHandler{
		redis:            redis,
		financialService: financialService,
	}
}

func (h *FinancialHandler) Handle(ctx context.Context, input *tgbotapi.Message) (*tgbotapi.MessageConfig, error) {
	splitted := strings.Split(input.Command(), "_")
	if len(splitted) != 2 {
		return nil, nil
	}

	service := splitted[0]
	command := splitted[1]

	if service != "fin" {
		return nil, nil
	}

	switch command {
	case "accountcreate":
		return h.AccountCreate(ctx, input)
	}

	return nil, nil
}

func (h *FinancialHandler) AccountCreate(ctx context.Context, input *tgbotapi.Message) (*tgbotapi.MessageConfig, error) {
	splitted := strings.Split(input.CommandArguments(), " ")
	if len(splitted) != 2 {
		output := tgbotapi.NewMessage(input.Chat.ID, "Invalid arguments!.\nUsage: /fin_accountcreate <name> <account_number>")
		return &output, nil
	}

	err := h.financialService.AccountCreate(ctx, &models.AccountCreateInput{
		Name:          splitted[0],
		AccountNumber: splitted[1],
	})
	if err != nil {
		return nil, err
	}

	output := tgbotapi.NewMessage(input.Chat.ID, "Account created successfully!")

	return &output, nil
}
