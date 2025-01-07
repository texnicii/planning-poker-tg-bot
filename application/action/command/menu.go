package cmd

import (
	tgbotapi "github.com/OvyFlash/telegram-bot-api"
	"planning_pocker_bot/application/action/common/handler"
	"planning_pocker_bot/application/service"
	"planning_pocker_bot/application/service/repository"
	"planning_pocker_bot/domain/entity"
	"planning_pocker_bot/infrastructure/telegram/messaging"
	_ "planning_pocker_bot/internal/translations"
)

type Menu struct {
	handler.Model
	Title   string
	buttons [][]tgbotapi.InlineKeyboardButton
}

func NewMenu() *Menu {
	return &Menu{
		buttons: [][]tgbotapi.InlineKeyboardButton{
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("Planning Poker", "callback/poker/game"),
				tgbotapi.NewInlineKeyboardButtonData("⚙️", "callback/settings"),
			),
		},
	}
}

func (m Menu) Handle(update tgbotapi.Update) *messaging.ResponseBag {
	group := repository.Read[entity.Group](m.readChatId(update))
	p := service.InitPrinter(group.Options.Lang)

	var title string
	if m.Title == "" {
		title = p.Sprintf("Menu")
	} else {
		title = m.Title
	}

	response := new(messaging.ResponseBag)

	switch m.Method {
	case "replace":
		response.AddEditMessageResponseWithMarkup(
			update.CallbackQuery.Message.Chat.ID,
			update.CallbackQuery.Message.MessageID,
			title,
			tgbotapi.NewInlineKeyboardMarkup(m.buttons...),
		)
	default:
		response.AddChatResponseWithMarkup(
			update.Message.Chat.ID,
			update.Message.MessageThreadID,
			title,
			tgbotapi.NewInlineKeyboardMarkup(m.buttons...),
		)
	}

	return response
}

func (m Menu) readChatId(update tgbotapi.Update) int64 {
	if update.CallbackQuery != nil {
		return update.CallbackQuery.Message.Chat.ID
	} else {
		return update.Message.Chat.ID
	}
}
