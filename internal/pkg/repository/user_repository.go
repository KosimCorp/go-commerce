package repository

import (
	"ecommerce/internal/pkg/model"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (u UserRepository) FindByCredentials(username string) (user model.User, err error) {
	err = u.db.First(&user, "username = ?", username).Error

	return
}
