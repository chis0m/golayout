package main

import (
	"fmt"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"go-layout/cmd/app"
	"go-layout/config"
	"go-layout/internal/routes"
	"go-layout/store/db"
	"go-layout/utils"
	"gorm.io/gorm"
	"os"
)

func main() {
	log.Logger = zerolog.New(os.Stderr).With().Timestamp().Logger()
	if utils.IsLocal() {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr}).With().Caller().Timestamp().Logger()
	}
	env, err := config.LoadConfig()
	var appDB *gorm.DB
	appDB, err = db.InitDB(env)
	if err != nil {
		log.Fatal().Err(err).Msg("AppDb connection failed")
	}
	appContext, err := app.Initialize(env, appDB, env.Token.SymmetricKey)
	if err != nil {
		log.Fatal().Err(err).Msgf("Failed to initialize application: %v", err)
	}
	router := routes.SetupRoutes(appContext)
	err = router.Run(fmt.Sprintf("%s:%s", env.App.Url, env.App.Port))
	if err != nil {
		log.Fatal().Err(err).Msg("Could not start server")
	}
}
