package model

import "gorm.io/gorm"

type DetailTransaksi struct {
	gorm.Model

	TransaksiID int
	Transaksi   Transaksi
	ProdukID    int
	Produk      Produk
	UserID      int
	User        User
	Total       int
	Qty         int
}
