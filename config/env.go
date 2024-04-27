package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

type envVars struct {
	PsqlConnectionUrl string
}

var Env = initEnv()

func initEnv() *envVars {
	err := godotenv.Load()
	if err != nil {
		log.Fatal().Msg("Faild To load .env")
	}

	return &envVars{
		PsqlConnectionUrl: os.Getenv("PG_CONN_URL"),
	}
}
