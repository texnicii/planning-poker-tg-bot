package planning_poker

import (
	"fmt"
	tgbotapi "github.com/OvyFlash/telegram-bot-api"
	"golang.org/x/text/message"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"math/rand"
	"planning_pocker_bot/application/action/common/handler"
	"planning_pocker_bot/application/service"
	"planning_pocker_bot/application/service/repository"
	"planning_pocker_bot/domain/entity"
	"planning_pocker_bot/infrastructure/config"
	"planning_pocker_bot/infrastructure/di"
	"planning_pocker_bot/infrastructure/telegram/messaging"
	_ "planning_pocker_bot/internal/translations"
)

type Game struct {
	handler.Model
	voteButtons [][]tgbotapi.InlineKeyboardButton
	db          *gorm.DB
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
		db: di.Get(config.DbClient).(*gorm.DB),
	}
}

func (game Game) Handle(update tgbotapi.Update) *messaging.ResponseBag {
	if update.CallbackQuery.Message.Chat.ID > 0 {
		return NotSupported{}.Handle(update)
	}

	group := repository.Read[entity.Group](update.CallbackQuery.Message.Chat.ID)
	langPrinter := service.InitPrinter(group.Options.Lang)

	response := new(messaging.ResponseBag)
	var output string
	buttons := game.voteButtons

	switch game.Method {
	case "vote":
		game.saveVote(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.From, game.Input.(string))
		votes := game.getVotes(update.CallbackQuery.Message.Chat.ID)
		output = viewVotes(votes)
		if len(votes) > 0 {
			buttons = append(buttons, game.buildRevealButton(langPrinter)...)
		}
	case "reveal":
		votes := game.getVotes(update.CallbackQuery.Message.Chat.ID)
		output = viewReveal(votes)
		buttons = game.buildFinishButton(langPrinter)
	case "next":
		game.ResetVotes(update.CallbackQuery.Message.Chat.ID)
	default:
		votes := game.getVotes(update.CallbackQuery.Message.Chat.ID)
		output = viewVotes(votes)
		if len(votes) > 0 {
			buttons = append(buttons, game.buildRevealButton(langPrinter)...)
		}
	}

	title := langPrinter.Sprintf("Planning Poker") + " ♥️♦️♠️♣️"

	if output != "" {
		output = title + "\n\n" + output
	} else {
		output = title
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

func (game Game) buildRevealButton(printer *message.Printer) [][]tgbotapi.InlineKeyboardButton {
	return [][]tgbotapi.InlineKeyboardButton{
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(
				"🃏"+printer.Sprintf("Reveal"),
				"callback/poker/game@reveal",
			),
		),
	}
}

func (game Game) buildFinishButton(printer *message.Printer) [][]tgbotapi.InlineKeyboardButton {
	return [][]tgbotapi.InlineKeyboardButton{
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(printer.Sprintf("Next ")+"➡️", "callback/poker/game@next"),
		),
	}
}

func chooseSticker() string {
	var stickers = []string{"🧑‍🔧", "🦄", "🧛🏻", "🐝", "🐣", "🦖", "🦋", "👻", "🕵🏻‍♂️", "👩🏻‍🚀", "🧑🏾‍💻"}
	i := rand.Intn(len(stickers) - 1)

	return stickers[i]
}
