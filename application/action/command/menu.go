package cmd

import (
	tgbotapi "github.com/OvyFlash/telegram-bot-api"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"planning_pocker_bot/application/action/common/handler"
	"planning_pocker_bot/infrastructure/telegram/messaging"
)

type Menu struct {
	handler.Model
	Title string
}

func (cmd Menu) Handle(update tgbotapi.Update) *messaging.ResponseBag {
	lang := language.Russian // FIXME - ?
	p := message.NewPrinter(lang)

	tgMessage := update.Message

	pokerLink := "callback/poker/game"
	if tgMessage.Chat.ID > 0 {
		pokerLink = "callback/poker/not-supported"
	}

	var menuButtons = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Planning Poker", pokerLink),
			tgbotapi.NewInlineKeyboardButtonData("⚙️", "callback/settings"),
		),
	)

	var title string
	if cmd.Title == "" {
		title = p.Sprintf("Menu")
	} else {
		title = cmd.Title
	}

	response := new(messaging.ResponseBag)
	response.AddChatResponseWithMarkup(tgMessage.Chat.ID, tgMessage.MessageThreadID, title, menuButtons)

	return response
}
