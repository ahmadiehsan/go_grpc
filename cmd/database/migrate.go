package main

import (
	"gogrpc/internal/blog"
	"gogrpc/internal/conf"
	"gogrpc/internal/database"
	"gogrpc/internal/logger"

	"github.com/rs/zerolog/log"
)

var models = []interface{}{
	&blog.Article{},
	&blog.Category{},
}

func main() {
	settings := conf.LoadSettings()
	if settings.Debug {
		logger.SwitchToHumanReadableMode()
	}

	db := database.Connect(settings)

	if err := db.AutoMigrate(models...); err != nil {
		log.Fatal().Err(err).Msg("Error in migration")
	}

	log.Info().Msg("Done")
}
