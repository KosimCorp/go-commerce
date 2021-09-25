package repository

import (
	"ecommerce/internal/pkg/model"

	"gorm.io/gorm"
)

type ProdukRepository struct {
	db *gorm.DB
}

func NewProdukRepository(db *gorm.DB) *ProdukRepository {
	return &ProdukRepository{db: db}
}

func (p ProdukRepository) FindAll() (result []model.Produk, err error) {
	err = p.db.Find(&result).Error

	return
}

func (p ProdukRepository) Find(id int) (result model.Produk, err error) {
	err = p.db.First(&result, id).Error

	return
}

func (p ProdukRepository) Create(produk *model.Produk) error {
	return p.db.Create(produk).Error
}
