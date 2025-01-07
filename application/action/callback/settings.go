package callback

import (
	tgbotapi "github.com/OvyFlash/telegram-bot-api"
	"gorm.io/gorm"
	"planning_pocker_bot/application/action/common/handler"
	"planning_pocker_bot/domain/entity"
	"planning_pocker_bot/infrastructure/config"
	"planning_pocker_bot/infrastructure/di"
	"planning_pocker_bot/infrastructure/telegram/messaging"
)

type Settings struct {
	handler.Model
	db      *gorm.DB
	buttons [][]tgbotapi.InlineKeyboardButton
}

func NewSettings() *Settings {
	return &Settings{
		buttons: [][]tgbotapi.InlineKeyboardButton{
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("🇷🇺", "callback/settings@lang:ru-RU"),
				tgbotapi.NewInlineKeyboardButtonData("🇬🇧", "callback/settings@lang:en-GB"),
			),
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("⬅️", "callback/menu@replace"),
			),
		},
		db: di.Get(config.DbClient).(*gorm.DB),
	}
}

func (s Settings) Handle(update tgbotapi.Update) *messaging.ResponseBag {
	response := new(messaging.ResponseBag)

	switch s.Method {
	case "lang":
		group := s.readGroup(update.CallbackQuery.Message.Chat.ID)
		// if the group options are not saved yet
		if group.ChatId == 0 {
			group.ChatId = update.CallbackQuery.Message.Chat.ID
		}
		group.Options.Lang = s.Input.(string)
		s.saveGroup(group)
		response.AddCallbackResponse(update.CallbackQuery.ID, "Saved 👌")
	default:
		response.AddEditMessageResponseWithMarkup(
			update.CallbackQuery.Message.Chat.ID,
			update.CallbackQuery.Message.MessageID,
			"Settings",
			tgbotapi.NewInlineKeyboardMarkup(s.buttons...),
		)
	}

	return response
}

func (s Settings) readGroup(id int64) *entity.Group {
	group := new(entity.Group)
	s.db.Find(&group, id)

	return group
}

func (s Settings) saveGroup(group *entity.Group) {
	s.db.Save(group)
}
