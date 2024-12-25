package cmd

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"planning_pocker_bot/application/action/common/handler"
	"planning_pocker_bot/infrastructure/telegram/messaging"
)

type Start struct {
	handler.Model
}

func (cmd Start) Handle(update tgbotapi.Update) *messaging.ResponseBag {
	menu := Menu{
		Title: "Welcome",
	}

	return menu.Handle(update)
}
