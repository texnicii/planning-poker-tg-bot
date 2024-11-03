package message

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"planning_pocker_bot/application/action/callback"
	"planning_pocker_bot/domain/service"
	"planning_pocker_bot/infrastructure/telegram/messaging"
)

type ChatMessageHandler struct {
}

func (handler ChatMessageHandler) Handle(update tgbotapi.Update) messaging.ResponseBag {
	awaiting, _ := service.GetAwaiting(update.Message.Chat.ID)

	// handle awaiting
	if awaiting != nil {
		switch awaiting.CommandCallback {
		case "callback/new_team":
			return callback.NewTeam{}.PutTeamName(update)
		}
	}

	// nothing to reply
	return messaging.ResponseBag{}
}
