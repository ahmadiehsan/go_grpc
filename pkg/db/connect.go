package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("db.sqlite"), &gorm.Config{}) // TODO read the db name from a config file
	if err != nil {
		//	TODO error handling
	}

	return db
}
