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
		Title: "Welcome to Planning Poker ♥️♦️♠️♣️",
	}

	// FIXME - signup or update new group
	//err := service.SignUp(service.SignUpDto{
	//	GroupId:   update.Message.Chat.ID,
	//	Nickname: update.Message.Chat.UserName,
	//	AltName:  update.Message.Chat.FirstName,
	//})
	//if err != nil {
	//	// TODO - need logging
	//	response.AddCallbackResponse(update.CallbackQuery.ID, "Hm... something goes wrong 🙄")
	//}

	return menu.Handle(update)
}
