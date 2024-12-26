package planning_poker

import (
	"fmt"
	tgbotapi "github.com/OvyFlash/telegram-bot-api"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"math/rand"
	"planning_pocker_bot/application/action/common/handler"
	"planning_pocker_bot/domain/entity"
	"planning_pocker_bot/infrastructure/config"
	"planning_pocker_bot/infrastructure/di"
	"planning_pocker_bot/infrastructure/telegram/messaging"
)

type Game struct {
	handler.Model
	voteButtons  [][]tgbotapi.InlineKeyboardButton
	revealButton [][]tgbotapi.InlineKeyboardButton
	finishButton [][]tgbotapi.InlineKeyboardButton
	db           *gorm.DB
	title        string
}

func NewGame() *Game {
	return &Game{
		voteButtons: [][]tgbotapi.InlineKeyboardButton{
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("1", "callback/poker/game@vote:1"),
				tgbotapi.NewInlineKeyboardButtonData("2", "callback/poker/game@vote:2"),
				tgbotapi.NewInlineKeyboardButtonData("3", "callback/poker/game@vote:3"),
				tgbotapi.NewInlineKeyboardButtonData("5", "callback/poker/game@vote:5"),
			),
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("8", "callback/poker/game@vote:8"),
				tgbotapi.NewInlineKeyboardButtonData("13", "callback/poker/game@vote:13"),
				tgbotapi.NewInlineKeyboardButtonData("21", "callback/poker/game@vote:21"),
				tgbotapi.NewInlineKeyboardButtonData("❓", "callback/poker/game@vote:?"),
			),
		},
		revealButton: [][]tgbotapi.InlineKeyboardButton{
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("🃏Reveal", "callback/poker/game@reveal"),
			),
		},
		finishButton: [][]tgbotapi.InlineKeyboardButton{
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("Next ➡️", "callback/poker/game@next"),
			),
		},
		db:    di.Get(config.DbClient).(*gorm.DB),
		title: "Planning Poker ♥️♦️♠️♣️",
	}
}

func (game Game) Handle(update tgbotapi.Update) *messaging.ResponseBag {
	response := new(messaging.ResponseBag)
	var output string
	buttons := game.voteButtons

	switch game.Method {
	case "vote":
		game.saveVote(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.From, game.Input.(string))
		votes := game.getVotes(update.CallbackQuery.Message.Chat.ID)
		output = viewVotes(votes)
		if len(votes) > 0 {
			buttons = append(buttons, game.revealButton...)
		}
	case "reveal":
		votes := game.getVotes(update.CallbackQuery.Message.Chat.ID)
		output = viewReveal(votes)
		buttons = game.finishButton
	case "next":
		game.ResetVotes(update.CallbackQuery.Message.Chat.ID)
	default:
		votes := game.getVotes(update.CallbackQuery.Message.Chat.ID)
		output = viewVotes(votes)
		if len(votes) > 0 {
			buttons = append(buttons, game.revealButton...)
		}
	}

	if output != "" {
		output = game.title + "\n\n" + output
	} else {
		output = game.title
	}

	response.AddEditMessageResponseWithMarkup(
		update.CallbackQuery.Message.Chat.ID,
		update.CallbackQuery.Message.MessageID,
		output,
		tgbotapi.NewInlineKeyboardMarkup(buttons...),
	)

	return response
}

func (game Game) saveVote(groupId int64, from *tgbotapi.User, value string) {
	vote := entity.Vote{
		GroupId:   groupId,
		UserId:    from.ID,
		Name:      from.UserName,
		FirstName: from.FirstName,
		Icon:      chooseSticker(),
		Value:     value,
	}

	game.db.Clauses(
		clause.OnConflict{DoUpdates: clause.AssignmentColumns([]string{"value"})},
	).Create(&vote)
}

func (game Game) getVotes(groupId int64) []entity.Vote {
	var votes []entity.Vote
	result := game.db.Where("group_id = ?", groupId).Find(&votes)
	if result.Error != nil {
		fmt.Println(result.Error)
		return nil
	}

	return votes
}

func (game Game) ResetVotes(groupId int64) {
	game.db.Delete(&entity.Vote{}, "group_id = ?", groupId)
}

func chooseSticker() string {
	var stickers = []string{"🧑‍🔧", "🦄", "🧛🏻", "🐝", "🐣", "🦖", "🦋", "👻", "🕵🏻‍♂️", "👩🏻‍🚀", "🧑🏾‍💻"}
	i := rand.Intn(len(stickers) - 1)

	return stickers[i]
}
