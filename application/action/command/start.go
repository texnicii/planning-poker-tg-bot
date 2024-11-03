package cmd

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"planning_pocker_bot/domain/service"
	"planning_pocker_bot/infrastructure/telegram/messaging"
)

type Start struct {
}

func (cmd Start) Handle(update tgbotapi.Update) messaging.ResponseBag {
	menu := Menu{
		Title: "Welcome to Planning Poker ♥️♦️♠️♣️",
	}

	err := service.SignUp(service.SignUpDto{
		ChatId:   update.Message.Chat.ID,
		Nickname: update.Message.Chat.UserName,
		AltName:  update.Message.Chat.FirstName,
	})
	if err != nil {
		// TODO - need logging
		return MakeErrorResponse(update.Message.Chat.ID)
	}

	return menu.Handle(update)
}
