package main

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"go-layout/config"
	"go-layout/internal/routes"
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
		log.Fatal().Err(err).Msg("AppDb connection failed")
	}

	router := gin.Default()
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}
	router.Use(cors.New(corsConfig))
	routes.SetupRoutes(router)
	err = router.Run(fmt.Sprintf("%s:%s", env.AppUrl, env.AppPort))
	if err != nil {
		log.Fatal().Err(err).Msg("Could not start server")
	}
}
