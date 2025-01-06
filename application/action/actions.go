package action

import (
	"planning_pocker_bot/application/action/callback"
	"planning_pocker_bot/application/action/callback/planning_poker"
	cmd "planning_pocker_bot/application/action/command"
	"planning_pocker_bot/application/action/message"
	"planning_pocker_bot/infrastructure/telegram/handle"
)

type HandlersContainer map[string]handle.Handler

func (c HandlersContainer) Find(name string, isCallback bool) handle.Handler {
	if handler, ok := c[name]; ok {
		return handler
	} else {
		return &cmd.UnknownCommandHandler{
			IsCallback: isCallback,
		}
	}
}

func NewHandlersContainer() HandlersContainer {
	return HandlersContainer{
		// commands handles
		"/start": &cmd.Start{},
		"/echo":  &cmd.Echo{},
		"/menu":  &cmd.Menu{},
		"/stop":  &cmd.Stop{},
		// callbacks handles
		"callback/poker/game":          planning_poker.NewGame(),
		"callback/poker/not-supported": &planning_poker.NotSupported{},
		"callback/settings":            callback.NewSettings(),
		// messages handles
		handle.DefaultMessageHandlerAlias: &message.ChatMessageHandler{},
	}
}
