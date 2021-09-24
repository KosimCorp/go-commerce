package model

import "gorm.io/gorm"

type Transaksi struct {
	gorm.Model

	UserID int
	User   User
	Total  int
}
