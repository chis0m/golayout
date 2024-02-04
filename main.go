package main

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"go-layout/config"
	"go-layout/storage/db"
	"go-layout/utils"
	"os"
)

func main() {
	log.Logger = zerolog.New(os.Stderr).With().Timestamp().Logger()
	if utils.IsLocal() {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr}).With().Caller().Timestamp().Logger()
	}
	env, err := config.GetConfig()
	err = db.InitDB(env)
	if err != nil {
		log.Fatal().Err(err).Msg("DB connection failed")
	}
}
