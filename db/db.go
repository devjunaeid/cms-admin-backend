package db

import (
	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB(dbURL string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Fatal().Msg("Faild to Connect to the Database!")
		return nil
	}

	log.Info().Msg("Database Initialize Successfully!!")
	return db
}
