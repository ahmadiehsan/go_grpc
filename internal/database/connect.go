package database

import (
	"gogrpc/internal/conf"

	"github.com/rs/zerolog/log"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Connect(s conf.Settings) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(s.DBName), &gorm.Config{})
	if err != nil {
		log.Fatal().Err(err).Msg("Error in db connection")
	}

	return db
}
