package main

import (
	"gogrpc/internal/database"
	"gogrpc/internal/logger"
	"gogrpc/pkg/blog"

	"github.com/rs/zerolog/log"
)

var models = []interface{}{
	&blog.Article{},
	&blog.Category{},
}

func init() {
	logger.SwitchToHumanReadableMode()
}

func main() {
	db, err := database.Connect()
	if err != nil {
		log.Fatal().Err(err).Msg("Error in db connection")
	}

	if errMigrate := db.AutoMigrate(models...); errMigrate != nil {
		log.Fatal().Err(errMigrate).Msg("Error in migration")
	}

	log.Info().Msg("Done")
}
