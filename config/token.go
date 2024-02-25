package config

import (
	"go-layout/utils"
	"os"
	"time"
)

type TokenConfig struct {
	AccessDuration  time.Duration
	RefreshDuration time.Duration
	SymmetricKey    string `validate:"require"`
}

func LoadTokenConfig() (TokenConfig, error) {
	return TokenConfig{
		AccessDuration:  utils.GetDuration("ACCESS_TOKEN_DURATION", "15m"),
		RefreshDuration: utils.GetDuration("REFRESH_TOKEN_DURATION", "24h"),
		SymmetricKey:    os.Getenv("TOKEN_SYMMETRIC_KEY"),
	}, nil
}
