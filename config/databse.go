package config

import (
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func GetDB() *gorm.DB {

	if db == nil {
		var err error
		dns := os.Getenv("DB_CONNECTION")

		db, err = gorm.Open(sqlite.Open(dns), &gorm.Config{})
		if err != nil {
			panic("failed to connect database")
		}
	}
	return db

}
