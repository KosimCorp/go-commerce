package model

import "gorm.io/gorm"

type Produk struct {
	gorm.Model

	Name             string `binding:"required"`
	Description      string `binding:"required"`
	Harga            int    `binding:"required"`
	ImageUrl         string
	KategoriProdukID int `binding:"required"`
	KategoriProduk   KategoriProduk
	TokoID           int
	Toko             Toko
}
