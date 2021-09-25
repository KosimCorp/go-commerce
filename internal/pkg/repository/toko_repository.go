package repository

import "gorm.io/gorm"

type TokoRepository struct {
	db *gorm.DB
}

func NewTokoRepository(db *gorm.DB) *TokoRepository {
	return &TokoRepository{db: db}
}
