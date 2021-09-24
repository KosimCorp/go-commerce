package model

import "gorm.io/gorm"

type Toko struct {
	gorm.Model

	Name   string
	UserID int
	User   User
	Alamat string
}
