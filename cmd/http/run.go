package main

import (
	"gogrpc/api/http"
	"gogrpc/internal/blog"
	"gogrpc/internal/conf"
	"gogrpc/internal/database"
	"gogrpc/internal/logger"

	"github.com/rs/zerolog/log"
)

func main() {
	settings := conf.LoadSettings()
	if settings.Debug {
		logger.SwitchToHumanReadableMode()
	}

	db := database.Connect(settings)
	as := &blog.ArticleService{DB: db}
	server := http.NewServer(settings, as)

	err := server.Run()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to run server")
	}
}
