package repo

import (
	"alog/internal/model"
	"gorm.io/gorm"
)

type UserRepo struct {
}

func NewUserRepo() *UserRepo {
	return &UserRepo{}
}

func (UserRepo) Update(db *gorm.DB, user *model.User) (err error) {
	return db.Save(user).Error
}

func (UserRepo) Get(db *gorm.DB, user *model.User) (err error) {
	return db.First(user).Error
}
