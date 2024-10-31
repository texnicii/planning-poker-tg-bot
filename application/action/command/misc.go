package cmd

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"planning_pocker_bot/domain/service"
	"planning_pocker_bot/infrastructure/telegram/messaging"
)

type UnknownCommandHandler struct {
	IsCallback bool
}

func (cmd UnknownCommandHandler) Handle(update tgbotapi.Update) messaging.ResponseBag {
	response := messaging.ResponseBag{}
	if cmd.IsCallback {
		// TODO i18n
		response.AddCallbackResponse(update.CallbackQuery.ID, "Упс... Что-то пошло не так 🙄")
	} else {
		response.AddChatResponse(update.Message.Chat.ID, "👀Неизвестная команда")
	}

	return response
}

type Start struct {
}

func (cmd Start) Handle(update tgbotapi.Update) messaging.ResponseBag {
	// TODO i18n
	menu := Menu{
		Title: "Привет!",
	}

	err := service.SignUp(service.SignUpDto{
		ChatId:   update.Message.Chat.ID,
		Nickname: update.Message.Chat.UserName,
		AltName:  update.Message.Chat.FirstName,
	})
	if err != nil {
		response := messaging.ResponseBag{}
		response.AddChatResponse(update.Message.Chat.ID, "Упс... Что-то пошло не так 🙄")
		// TODO - need logging

		return response
	}

	return menu.Handle(update)
}

type Menu struct {
	Title string
}

func (cmd Menu) Handle(update tgbotapi.Update) messaging.ResponseBag {
	message := update.Message
	var chatBtn tgbotapi.InlineKeyboardButton

	chatBtn = tgbotapi.NewInlineKeyboardButtonData("🦄Новый чат", "new_chat")

	var menuButtons = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			chatBtn,
			tgbotapi.NewInlineKeyboardButtonData("Прочитать правила", "rules"),
		),
	)

	response := messaging.ResponseBag{}
	var title string
	if cmd.Title == "" {
		// TODO i18n
		title = "Меню"
	} else {
		title = cmd.Title
	}

	response.AddChatResponseWithMarkup(message.Chat.ID, title, menuButtons)
	return response
}

type Echo struct {
}

func (cmd Echo) Handle(update tgbotapi.Update) messaging.ResponseBag {
	response := messaging.ResponseBag{}
	response.AddChatResponse(update.Message.Chat.ID, "✅")

	return response
}
