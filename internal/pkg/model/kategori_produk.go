package model

import "gorm.io/gorm"

type KategoriProduk struct {
	gorm.Model

	Name string
}
