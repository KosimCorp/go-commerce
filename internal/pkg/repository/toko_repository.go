package repository

import (
	"ecommerce/internal/pkg/model"

	"gorm.io/gorm"
)

type TokoRepository struct {
	db *gorm.DB
}

func NewTokoRepository(db *gorm.DB) *TokoRepository {
	return &TokoRepository{db: db}
}

func (t TokoRepository) Create(toko *model.Toko) error {
	return t.db.Create(toko).Error
}

func (t TokoRepository) FindByUser(user model.User) (toko *model.Toko, err error) {
	err = t.db.First(toko, user.ID).Error

	return
}
