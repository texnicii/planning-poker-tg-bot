package handle

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"planning_pocker_bot/infrastructure/telegram"
	"planning_pocker_bot/infrastructure/telegram/messaging"
)

type Controller struct {
	bot      *telegram.BotClient
	handlers HandlerContainer
}

func NewController(bot *telegram.BotClient, handlersContainer HandlerContainer) Controller {
	return Controller{
		bot:      bot,
		handlers: handlersContainer,
	}
}

func (conductor *Controller) Handle() {
	updates := conductor.bot.Api().GetUpdatesChan(conductor.bot.Config())
	var responses messaging.ResponseBag

	for update := range updates {
		go func(updateItem tgbotapi.Update) {
			if updateItem.Message != nil {
				if updateItem.Message.IsCommand() {
					responses = conductor.HandleCommand(updateItem)
				} else {
					responses = conductor.HandleMessage(updateItem)
				}
			} else if updateItem.CallbackQuery != nil {
				responses = conductor.HandleCallback(updateItem)
			}

			conductor.send(responses)
		}(update)
	}
}

func (conductor *Controller) HandleCommand(update tgbotapi.Update) messaging.ResponseBag {
	message := update.Message
	messageEntity := message.Entities[0]
	commandName := message.Text[messageEntity.Offset:messageEntity.Length]
	commandHandler := conductor.handlers.Find(commandName, false)

	return commandHandler.Handle(update)
}

func (conductor *Controller) HandleMessage(update tgbotapi.Update) messaging.ResponseBag {
	msgHandler := conductor.handlers.Find(DefaultMessageHandlerAlias, false)

	return msgHandler.Handle(update)
}

func (conductor *Controller) HandleCallback(update tgbotapi.Update) messaging.ResponseBag {
	callbackHandler := conductor.handlers.Find(update.CallbackQuery.Data, true)

	return callbackHandler.Handle(update)
}

func (conductor *Controller) send(responseBag messaging.ResponseBag) {
	for _, response := range responseBag.Responses {
		switch response.(type) {
		case messaging.ChatResponse:
			chatResponse := response.(messaging.ChatResponse)
			msg := tgbotapi.NewMessage(chatResponse.ChatId(), chatResponse.Text())
			markup := chatResponse.Markup()
			if markup != nil {
				msg.ReplyMarkup = markup
			}

			if msg.Text != "" || msg.ReplyMarkup != nil {
				_, err := conductor.bot.Api().Send(msg)
				if err != nil {
					// TODO - need logging
					fmt.Println(err)
				}
			}
		case messaging.CallbackResponse:
			callbackResponse := response.(messaging.CallbackResponse)
			if callbackResponse.Text() != "" {
				callback := tgbotapi.NewCallback(callbackResponse.QueryId(), callbackResponse.Text())
				_, err := conductor.bot.Api().Request(callback)
				if err != nil {
					// TODO - need logging
					fmt.Println(err)
				}
			}
		}
	}
}
