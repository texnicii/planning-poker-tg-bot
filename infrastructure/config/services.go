package config

import "planning_pocker_bot/infrastructure/di"

const (
	BotClient      = di.ServiceKey("bot_client")
	DbClient       = di.ServiceKey("db_client")
	UserRepository = di.ServiceKey("user_repository")
)
