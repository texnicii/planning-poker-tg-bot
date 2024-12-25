package cmd

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"planning_pocker_bot/application/action/common/handler"
	"planning_pocker_bot/infrastructure/telegram/messaging"
)

type Menu struct {
	handler.Model
	Title string
}

func (cmd Menu) Handle(update tgbotapi.Update) *messaging.ResponseBag {
	message := update.Message

	pokerLink := "callback/poker/game"
	if message.Chat.ID > 0 {
		pokerLink = "callback/poker/not-supported"
	}

	var menuButtons = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Planning Poker", pokerLink),
			tgbotapi.NewInlineKeyboardButtonData("⚙️", "setting"),
		),
	)

	var title string
	if cmd.Title == "" {
		title = "Menu"
	} else {
		title = cmd.Title
	}

	response := new(messaging.ResponseBag)
	response.AddChatResponseWithMarkup(message.Chat.ID, title, menuButtons)

	return response
}
