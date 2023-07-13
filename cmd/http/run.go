package main

import (
	"gogrpc/api/http"
	"gogrpc/internal/database"
	"gogrpc/internal/logger"

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

	server := http.NewServer(db)

	errRun := server.Run()
	if errRun != nil {
		log.Fatal().Err(errRun).Msg("Failed to run server")
	}
}
