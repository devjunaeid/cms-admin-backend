package main

import (
	"github.com/devjunaeid/cms-admin-backend/cmd/api"
	"github.com/devjunaeid/cms-admin-backend/config"
	"github.com/devjunaeid/cms-admin-backend/db"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	// Initializing Database.
	database := db.InitDB(config.Env.PsqlConnectionUrl)
	if database == nil {
		log.Fatal().Msg("Failed to initalized Database! Exiting!!")
		return
	}

	// Initializing API.
	api.CreateAPI(database)
}
