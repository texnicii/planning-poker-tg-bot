package handler

import "planning_pocker_bot/infrastructure/telegram/handle"

type Method string

type Model struct {
	Method Method
	Input  any
	handle.Handler
}

// SetInput Sets additional (method and data) for handler
// ! NOTE should be reset per every handler call because a handler is created once
func (a *Model) SetInput(action string, input any) {
	a.Input = input
	a.Method = Method(action)
}
