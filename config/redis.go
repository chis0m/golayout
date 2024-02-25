package config

import (
	"go-layout/utils"
)

type RedisConfig struct {
	Host string
	Port string
}

func LoadCacheConfig() (RedisConfig, error) {
	return RedisConfig{
		Host: utils.Getenv("REDIS_HOST", "localhost"),
		Port: utils.Getenv("REDIS_PORT", "5432"),
	}, nil
}
