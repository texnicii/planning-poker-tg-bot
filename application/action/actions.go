package action

import (
	"planning_pocker_bot/application/action/callback"
	cmd "planning_pocker_bot/application/action/command"
	"planning_pocker_bot/application/action/message"
	"planning_pocker_bot/infrastructure/telegram/handle"
)

type HandlersContainer struct {
	list map[string]handle.Handler
}

func (c HandlersContainer) Find(name string, isCallback bool) handle.Handler {
	if command, ok := c.list[name]; ok {
		return command
	} else {
		return cmd.UnknownCommandHandler{
			IsCallback: isCallback,
		}
	}
}

func NewHandlersContainer() HandlersContainer {
	return HandlersContainer{
		list: map[string]handle.Handler{
			// commands handles
			"/start": cmd.Start{},
			"/menu":  cmd.Menu{},
			"/echo":  cmd.Echo{},
			//"/stop":  cmd.StopCommandHandler{},
			// callbacks handles
			"callback/some": callback.NewGame{},
			// messages handles
			handle.DefaultMessageHandlerAlias: message.ChatMessageHandler{},
		},
	}
}
