package config

import (
	"go-layout/utils"
)

type AppConfig struct {
	Name        string
	Environment string `validate:"require"`
	Url         string
	Port        string
}

func LoadAppConfig() (AppConfig, error) {
	return AppConfig{
		Name:        utils.Getenv("APP_NAME", "GoLayoutApp"),
		Environment: utils.Getenv("APP_ENV", "local"),
		Url:         utils.Getenv("APP_URL", "localhost"),
		Port:        utils.Getenv("APP_PORT", "5000"),
	}, nil
}
