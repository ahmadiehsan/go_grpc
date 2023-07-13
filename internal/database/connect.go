package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Connect() (*gorm.DB, error) {
	return gorm.Open(sqlite.Open("db.sqlite"), &gorm.Config{}) // TODO read the db name from a config file
}
