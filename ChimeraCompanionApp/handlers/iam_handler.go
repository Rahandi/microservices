package handlers

import (
	"ChimeraCompanionApp/internals"
	"ChimeraCompanionApp/models"
	"ChimeraCompanionApp/services"
	"ChimeraCompanionApp/types"
	"context"
	"fmt"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type IAMHandler struct {
	redis      *internals.Redis
	iamService *services.IAMService
}

func NewIAMHandler(iamService *services.IAMService, redis *internals.Redis) *IAMHandler {
	return &IAMHandler{
		redis:      redis,
		iamService: iamService,
	}
}

func (h *IAMHandler) Handle(ctx context.Context, input *tgbotapi.Message) (*tgbotapi.MessageConfig, error) {
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
		return h.Register(ctx, input)
	case "login":
		return h.Login(ctx, input)
	case "whoami":
		return h.WhoAmI(ctx, input)
	}

	return nil, nil
}

func (h *IAMHandler) Register(ctx context.Context, input *tgbotapi.Message) (*tgbotapi.MessageConfig, error) {
	response, err := h.iamService.Register(ctx, &models.RegisterInput{
		Name:      strings.Join([]string{input.From.FirstName, input.From.LastName}, " "),
		AccountId: strconv.FormatInt(input.From.ID, 10),
		Password:  input.From.UserName + strconv.FormatInt(input.From.ID, 10),
	})
	if err != nil {
		return nil, err
	}

	accountId := ctx.Value(types.AccountIdKey).(string)
	user := models.User{
		ID:           accountId,
		Username:     input.From.UserName,
		Token:        response.Data.Token,
		RefreshToken: response.Data.RefreshToken,
	}
	_, err = h.redis.Client.HSet(ctx, accountId, user).Result()
	if err != nil {
		return nil, err
	}

	message := fmt.Sprintf("Registered successfully!\nHello %s!", input.From.FirstName)
	output := tgbotapi.NewMessage(input.Chat.ID, message)

	return &output, nil
}

func (h *IAMHandler) Login(ctx context.Context, input *tgbotapi.Message) (*tgbotapi.MessageConfig, error) {
	response, err := h.iamService.Login(ctx, &models.LoginInput{
		AccountId: strconv.FormatInt(input.From.ID, 10),
		Password:  input.From.UserName + strconv.FormatInt(input.From.ID, 10),
	})
	if err != nil {
		return nil, err
	}

	accountId := ctx.Value(types.AccountIdKey).(string)
	exists, err := h.redis.Client.HExists(ctx, accountId, "token").Result()
	if err != nil {
		return nil, err
	}
	if !exists {
		user := models.User{
			ID:           accountId,
			Username:     input.From.UserName,
			Token:        response.Data.Token,
			RefreshToken: response.Data.RefreshToken,
		}
		_, err = h.redis.Client.HSet(ctx, accountId, user).Result()
		if err != nil {
			return nil, err
		}
	} else {
		_, err = h.redis.Client.HSet(ctx, accountId, "token", response.Data.Token, "refresh_token", response.Data.RefreshToken).Result()
		if err != nil {
			return nil, err
		}
	}

	message := fmt.Sprintf("Logged in successfully!\nHello %s!", input.From.FirstName)
	output := tgbotapi.NewMessage(input.Chat.ID, message)

	return &output, nil
}

func (h *IAMHandler) WhoAmI(ctx context.Context, input *tgbotapi.Message) (*tgbotapi.MessageConfig, error) {
	response, err := h.iamService.WhoAmI(ctx)
	if err != nil {
		return nil, err
	}

	message := fmt.Sprintf("Name: %s \nPrincipal: %s", response.Data.Name, response.Data.Principal)
	output := tgbotapi.NewMessage(input.Chat.ID, message)

	return &output, nil
}
