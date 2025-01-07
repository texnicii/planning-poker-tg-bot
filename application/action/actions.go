package action

import (
	"planning_pocker_bot/application/action/callback"
	"planning_pocker_bot/application/action/callback/planning_poker"
	cmd "planning_pocker_bot/application/action/command"
	"planning_pocker_bot/application/action/message"
	"planning_pocker_bot/infrastructure/telegram/handle"
)

type HandlersContainer map[string]handle.Handler

func (c HandlersContainer) Find(name string) handle.Handler {
	if handler, ok := c[name]; ok {
		return handler
	} else {
		return &cmd.UnknownCommandHandler{}
	}
}

func NewHandlersContainer() HandlersContainer {
	menu := cmd.NewMenu()

	return HandlersContainer{
		// commands handles
		"/start":        &cmd.Start{},
		"/echo":         &cmd.Echo{},
		"/menu":         menu,
		"callback/menu": menu,
		"/stop":         &cmd.Stop{},
		// callbacks handles
		"callback/poker/game": planning_poker.NewGame(),
		"callback/settings":   callback.NewSettings(),
		// messages handles
		handle.DefaultMessageHandlerAlias: &message.ChatMessageHandler{},
	}
}
