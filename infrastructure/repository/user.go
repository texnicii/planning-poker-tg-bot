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
