package model

import "gorm.io/gorm"

type Produk struct {
	gorm.Model

	Name             string
	Description      string
	Harga            int
	ImageUrl         string
	KategoriProdukID int
	KategoriProduk   KategoriProduk
	TokoID           int
	Toko             Toko
}
