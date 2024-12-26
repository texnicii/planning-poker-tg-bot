package telegram

import tgbotapi "github.com/OvyFlash/telegram-bot-api"

type BotClient struct {
	api    *tgbotapi.BotAPI
	config tgbotapi.UpdateConfig
}

func NewBotClient(token string) (BotClient, error) {
	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 60
	api, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return BotClient{}, err
	}

	return BotClient{
		api:    api,
		config: updateConfig,
	}, nil
}

func (bot *BotClient) Api() *tgbotapi.BotAPI {
	return bot.api
}

func (bot *BotClient) Config() tgbotapi.UpdateConfig {
	return bot.config
}
