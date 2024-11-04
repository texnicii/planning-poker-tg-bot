package callback

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	cmd "planning_pocker_bot/application/action/command"
	"planning_pocker_bot/domain/entity"
	"planning_pocker_bot/domain/service"
	"planning_pocker_bot/infrastructure/telegram/messaging"
	"time"
)

const AwaitingTeamNameKey entity.CallbackKey = "team_name"

type NewTeam struct{}

func (n NewTeam) Handle(update tgbotapi.Update) messaging.ResponseBag {
	response := messaging.ResponseBag{}
	response.AddChatResponse(update.CallbackQuery.Message.Chat.ID, "What is the name of your team?")
	err := service.StartAwaiting(
		update.CallbackQuery.Message.Chat.ID,
		AwaitingTeamNameKey,
		time.Now().Add(time.Minute*1),
	)
	if err != nil {
		return cmd.MakeErrorResponse(update.CallbackQuery.Message.Chat.ID)
	}

	return response
}
