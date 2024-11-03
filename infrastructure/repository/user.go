package repository

import (
	"gorm.io/gorm"
	"planning_pocker_bot/domain/entity"
)

type UserRepository struct {
	db *gorm.DB
}

func (u *UserRepository) Create(chatId int64, nickname string, altName string) (*entity.User, error) {
	user := entity.NewUser(chatId, nickname, altName)
	result := u.db.Create(user)

	return user, result.Error
}

func (u *UserRepository) Delete(chatId int64) error {
	result := u.db.Delete(&entity.User{ChatId: chatId})

	return result.Error
}

func (u *UserRepository) Get(chatId int64) (*entity.User, error) {
	user := &entity.User{}
	result := u.db.Find(user, chatId)

	if user.ChatId != 0 {
		return user, result.Error
	}

	return nil, result.Error
}
