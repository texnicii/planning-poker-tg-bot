package repository

import (
	"gorm.io/gorm"
	"planning_pocker_bot/domain/entity"
)

type GroupRepository struct {
	db *gorm.DB
}

func (u *GroupRepository) Create(user *entity.Group) (*entity.Group, error) {
	result := u.db.Create(user)

	return user, result.Error
}

func (u *GroupRepository) Delete(chatId int64) error {
	result := u.db.Delete(&entity.Group{ChatId: chatId})

	return result.Error
}

func (u *GroupRepository) Get(chatId int64) (*entity.Group, error) {
	user := &entity.Group{}
	result := u.db.Find(user, chatId)

	if user.ChatId != 0 {
		return user, result.Error
	}

	return nil, result.Error
}
