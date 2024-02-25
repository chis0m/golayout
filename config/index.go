package config

import (
	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"go-layout/utils"
	"os"
)

type Config struct {
	App   AppConfig
	Token TokenConfig
	Db    DatabaseConfig
	Mail  MailConfig
	Cache RedisConfig
}

func LoadConfig() (*Config, error) {
	log.Logger = zerolog.New(os.Stderr).With().Timestamp().Logger()
	if utils.IsLocal() {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr}).With().Caller().Timestamp().Logger()
	}

	if err := godotenv.Load(".env"); err != nil {
		log.Err(err).Msg("Could not load .env file")
		return nil, err
	}

	appConfig, err := LoadAppConfig()
	if err != nil {
		log.Err(err).Msg("Failed to load App config")
		return nil, err
	}

	dbConfig, err := LoadDatabaseConfig()
	if err != nil {
		log.Err(err).Msg("Failed to load Database config")
		return nil, err
	}

	cacheConfig, err := LoadCacheConfig()
	if err != nil {
		log.Err(err).Msg("Failed to load Cache config")
		return nil, err
	}

	mailConfig, err := LoadMailConfig()
	if err != nil {
		log.Err(err).Msg("Failed to load mail config")
		return nil, err
	}

	tokenConfig, err := LoadTokenConfig()
	if err != nil {
		log.Err(err).Msg("Failed to load Auth Token config")
		return nil, err
	}

	return &Config{
		App:   appConfig,
		Db:    dbConfig,
		Cache: cacheConfig,
		Mail:  mailConfig,
		Token: tokenConfig,
	}, nil
}
