package main

import (
	"gogrpc/api/http"
	"gogrpc/internal/database"
	"gogrpc/internal/logger"
	"gogrpc/pkg/blog"

	"github.com/rs/zerolog/log"
)

func init() {
	logger.SwitchToHumanReadableMode()
}

func main() {
	db, err := database.Connect()
	if err != nil {
		log.Fatal().Err(err).Msg("Error in db connection")
	}

	as := &blog.ArticleService{DB: db}
	server := http.NewServer(as)

	errRun := server.Run()
	if errRun != nil {
		log.Fatal().Err(errRun).Msg("Failed to run server")
	}
}
