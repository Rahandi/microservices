package handlers

import (
	"ChimeraCompanionApp/models"
	"ChimeraCompanionApp/services"
	"fmt"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type IAMHandler struct {
	iamService *services.IAMService
}

func NewIAMHandler(iamService *services.IAMService) *IAMHandler {
	return &IAMHandler{
		iamService: iamService,
	}
}

func (h *IAMHandler) Handle(input *tgbotapi.Message) (*tgbotapi.MessageConfig, error) {
	splitted := strings.Split(input.Command(), "_")
	if len(splitted) != 2 {
		return nil, nil
	}

	service := splitted[0]
	command := splitted[1]

	if service != "iam" {
		return nil, nil
	}

	switch command {
	case "register":
		return h.Register(input)
	case "login":
		return h.Login(input)
	}

	return nil, nil
}

func (h *IAMHandler) Register(input *tgbotapi.Message) (*tgbotapi.MessageConfig, error) {
	response, err := h.iamService.Register(&models.RegisterInput{
		Name:      input.From.UserName,
		Email:     input.From.UserName,
		AccountId: strconv.FormatInt(input.From.ID, 10),
	})
	if err != nil {
		return nil, err
	}

	message := fmt.Sprintf("Token: %s \nRefresh Token: %s", response.Data.Token, response.Data.RefreshToken)
	output := tgbotapi.NewMessage(input.Chat.ID, message)

	return &output, nil
}

func (h *IAMHandler) Login(input *tgbotapi.Message) (*tgbotapi.MessageConfig, error) {
	response, err := h.iamService.Login(&models.LoginInput{
		Email:     input.From.UserName,
		AccountId: strconv.FormatInt(input.From.ID, 10),
	})
	if err != nil {
		return nil, err
	}

	message := fmt.Sprintf("Token: %s \nRefresh Token: %s", response.Data.Token, response.Data.RefreshToken)
	output := tgbotapi.NewMessage(input.Chat.ID, message)

	return &output, nil
}
