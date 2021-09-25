package app

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func mysqlDialect() gorm.Dialector {
	user := "root"
	password := ""
	host := "127.0.0.1"
	port := 3306
	db := "go-commerce"

	return mysql.Open(
		fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, db),
	)
}

func sqliteDialect() gorm.Dialector {
	return sqlite.Open("database.db")
}

func Database() *gorm.DB {
	driverName := "sqlite"
	var driver gorm.Dialector

	if driverName == "mysql" {
		driver = mysqlDialect()
	}
	if driverName == "sqlite" {
		driver = sqliteDialect()
	}

	db, err := gorm.Open(driver)

	if err != nil {
		panic(err.Error())
	}

	return db
}
