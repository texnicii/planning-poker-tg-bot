package handle

import (
	tgbotapi "github.com/OvyFlash/telegram-bot-api"
	"github.com/rs/zerolog/log"
	"planning_pocker_bot/infrastructure/telegram"
	"planning_pocker_bot/infrastructure/telegram/messaging"
	"strings"
)

type Controller struct {
	bot      *telegram.BotClient
	handlers HandlersContainer
}

func NewController(bot *telegram.BotClient, handlersContainer HandlersContainer) Controller {
	return Controller{
		bot:      bot,
		handlers: handlersContainer,
	}
}

func (c *Controller) Handle() {
	updates := c.bot.Api().GetUpdatesChan(c.bot.Config())
	var responseBag *messaging.ResponseBag

	for update := range updates {
		go func(updateItem tgbotapi.Update) {
			if updateItem.Message != nil {
				if updateItem.Message.IsCommand() {
					responseBag = c.HandleCommand(updateItem)
				} else {
					responseBag = c.HandleMessage(updateItem)
				}
			} else if updateItem.CallbackQuery != nil {
				responseBag = c.HandleCallback(updateItem)
			}

			c.send(responseBag)
		}(update)
	}
}

func (c *Controller) HandleCommand(update tgbotapi.Update) *messaging.ResponseBag {
	message := update.Message
	commandName := c.parseCommand(message.Text[message.Entities[0].Offset:message.Entities[0].Length])
	commandHandler := c.handlers.Find(commandName)

	return commandHandler.Handle(update)
}

func (c *Controller) HandleMessage(update tgbotapi.Update) *messaging.ResponseBag {
	msgHandler := c.handlers.Find(DefaultMessageHandlerAlias)

	return msgHandler.Handle(update)
}

func (c *Controller) HandleCallback(update tgbotapi.Update) *messaging.ResponseBag {
	var action string
	var data any = nil

	decodedInput := strings.SplitN(update.CallbackQuery.Data, "@", 2)
	handlerName := decodedInput[0]
	if len(decodedInput) > 1 {
		decodedAction := strings.SplitN(decodedInput[1], ":", 2)
		action = decodedAction[0]
		if len(decodedAction) > 1 {
			data = decodedAction[1]
		}
	}

	callbackHandler := c.handlers.Find(handlerName)
	callbackHandler.SetInput(action, data)
	defer func() {
		// clear input
		callbackHandler.SetInput("", "")
	}()

	return callbackHandler.Handle(update)
}

func (c *Controller) send(responseBag *messaging.ResponseBag) {
	if responseBag == nil {
		// nothing to send
		return
	}

	for _, response := range responseBag.Responses {
		switch response.(type) {
		case messaging.ChatResponse:
			chatResponse := response.(messaging.ChatResponse)
			msg := tgbotapi.NewMessage(chatResponse.ChatId(), chatResponse.Text())
			msg.MessageThreadID = chatResponse.ThreadId()
			msg.ParseMode = tgbotapi.ModeHTML
			markup := chatResponse.Markup()
			if markup != nil {
				msg.ReplyMarkup = markup
			}

			if msg.Text != "" || msg.ReplyMarkup != nil {
				_, err := c.bot.Api().Send(msg)
				if err != nil {
					log.Error().Str("context", "chat response").Msg(err.Error())
				}
			}
		case messaging.CallbackResponse:
			callbackResponse := response.(messaging.CallbackResponse)
			if callbackResponse.Text() != "" {
				callback := tgbotapi.NewCallback(callbackResponse.QueryId(), callbackResponse.Text())
				_, err := c.bot.Api().Request(callback)
				if err != nil {
					log.Error().Str("context", callbackResponse.Text()).Msg(err.Error())
				}
			}
		case messaging.EditMessageResponse:
			chatResponse := response.(messaging.EditMessageResponse)
			msg := tgbotapi.NewEditMessageText(chatResponse.ChatId(), chatResponse.MessageId, chatResponse.Text())
			msg.ParseMode = tgbotapi.ModeHTML
			markup := chatResponse.Markup()
			if markup != nil {
				m := markup.(tgbotapi.InlineKeyboardMarkup)
				msg.BaseEdit.ReplyMarkup = &m
			}

			if msg.Text != "" || msg.ReplyMarkup != nil {
				_, err := c.bot.Api().Send(msg)
				if err != nil {
					log.Error().Str("context", "edit").Msg(err.Error())
				}
			}
		}
	}
}

// Parses command name if command has suffix with @ separator
func (c *Controller) parseCommand(command string) string {
	return strings.Split(command, "@")[0]
}
