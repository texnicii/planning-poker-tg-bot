package main

import (
	"planning_pocker_bot/application/action"
	"planning_pocker_bot/infrastructure"
	"planning_pocker_bot/infrastructure/config"
	"planning_pocker_bot/infrastructure/di"
	"planning_pocker_bot/infrastructure/telegram"
	"planning_pocker_bot/infrastructure/telegram/handle"
)

func main() {
	infrastructure.Bootstrap()

	controller := handle.NewController(
		di.Get(config.BotClient).(*telegram.BotClient),
		action.NewHandlersContainer(),
	)
	controller.Handle()
}
