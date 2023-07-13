package main

import (
	"gogrpc/internal/logger"
	"gogrpc/pkg/blog"
	"gogrpc/pkg/db"

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
	db := db.ConnectDB()
	if err := db.AutoMigrate(models...); err != nil {
		log.Fatal().Err(err).Msg("Error in migration")
	}
	log.Info().Msg("Done")
}
