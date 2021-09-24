package app

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Database() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("database.db"))

	if err != nil {
		panic(err.Error())
	}

	return db
}
