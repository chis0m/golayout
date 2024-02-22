package main

import (
	"fmt"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"go-layout/config"
	"go-layout/internal/app"
	"go-layout/internal/routes"
	"go-layout/storage/db"
	"go-layout/utils"
	"gorm.io/gorm"
	"os"
)

func main() {
	log.Logger = zerolog.New(os.Stderr).With().Timestamp().Logger()
	if utils.IsLocal() {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr}).With().Caller().Timestamp().Logger()
	}
	env, err := config.GetConfig()
	var appDB *gorm.DB
	appDB, err = db.InitDB(env)
	if err != nil {
		log.Fatal().Err(err).Msg("AppDb connection failed")
	}
	appContext, err := app.Initialize(appDB, env.TokenSymmetricKey)
	if err != nil {
		log.Fatal().Err(err).Msgf("Failed to initialize application: %v", err)
	}
	router := routes.SetupRoutes(appContext)
	err = router.Run(fmt.Sprintf("%s:%s", env.AppUrl, env.AppPort))
	if err != nil {
		log.Fatal().Err(err).Msg("Could not start server")
	}
}
