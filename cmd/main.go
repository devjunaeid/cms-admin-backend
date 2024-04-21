package main

import (
	"github.com/devjunaeid/cms-admin-backend/cmd/api"
	"github.com/rs/zerolog"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	api.CreateAPI()
}
