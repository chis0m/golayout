package config

import "go-layout/utils"

type DatabaseConfig struct {
	Host     string `validate:"require"`
	Name     string `validate:"require"`
	Username string `validate:"require"`
	Password string `validate:"require"`
	Port     string `validate:"require"`
}

func LoadDatabaseConfig() (DatabaseConfig, error) {
	return DatabaseConfig{
		Host:     utils.Getenv("DB_HOST", "localhost"),
		Name:     utils.Getenv("DB_NAME", "database"),
		Username: utils.Getenv("DB_USERNAME", "root"),
		Password: utils.Getenv("DB_PASSWORD", "password"),
		Port:     utils.Getenv("DB_PORT", "5434"),
	}, nil
}
