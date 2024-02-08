package config

import (
	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"go-layout/utils"
	"os"
)

type Config struct {
	AppEnv     string `validate:"require"`
	AppUrl     string
	AppPort    string
	DBHost     string `validate:"require"`
	DbName     string `validate:"require"`
	DbUsername string `validate:"require"`
	DbPassword string `validate:"require"`
	DbPort     string `validate:"require"`
	RedisUrl   string
	RedisPort  string
}

func GetConfig() (*Config, error) {
	log.Logger = zerolog.New(os.Stderr).With().Timestamp().Logger()
	if utils.IsLocal() {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr}).With().Caller().Timestamp().Logger()
	}

	if err := godotenv.Load(".env"); err != nil {
		log.Err(err).Msg("Could not load .env")
	}

	return &Config{
		AppEnv:     utils.Getenv("APP_ENV", "local"),
		AppUrl:     utils.Getenv("APP_URL", "127.0.0.1"),
		AppPort:    utils.Getenv("APP_PORT", "5000"),
		DBHost:     utils.Getenv("DB_HOST", "localhost"),
		DbName:     utils.Getenv("DB_NAME", "database"),
		DbUsername: utils.Getenv("DB_USER", "root"),
		DbPassword: utils.Getenv("DB_PASSWORD", "password"),
		DbPort:     utils.Getenv("DB_PORT", "5434"),
		RedisUrl:   os.Getenv("REDIS_URL"),
		RedisPort:  os.Getenv("REDIS_PORT"),
	}, nil
}
