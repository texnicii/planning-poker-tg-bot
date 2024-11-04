package await

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	cmd "planning_pocker_bot/application/action/command"
	"planning_pocker_bot/domain/service"
	"planning_pocker_bot/infrastructure/telegram/messaging"
)

type Team struct{}

func (t Team) Handle(update tgbotapi.Update) messaging.ResponseBag {
	err := service.NewTeam(service.TeamDto{
		Owner: update.Message.Chat.ID,
		Name:  update.Message.Text,
	})

	if err != nil {
		return cmd.MakeErrorResponse(update.Message.Chat.ID)
	}

	return cmd.Menu{}.Handle(update)
}
