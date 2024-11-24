package handler

import "planning_pocker_bot/infrastructure/telegram/handle"

type Method string

type Model struct {
	Method Method
	Input  any
	handle.Handler
}

func (a *Model) SetInput(action string, input any) {
	a.Input = input
	a.Method = Method(action)
}
