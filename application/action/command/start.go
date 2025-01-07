package cmd

import (
	tgbotapi "github.com/OvyFlash/telegram-bot-api"
	"planning_pocker_bot/application/action/common/handler"
	"planning_pocker_bot/infrastructure/telegram/messaging"
)

type Start struct {
	handler.Model
}

func (cmd Start) Handle(update tgbotapi.Update) *messaging.ResponseBag {
	menu := NewMenu()
	menu.Title = "Welcome"

	return menu.Handle(update)
}
