package utils

import (
	"go-layout/types"
	"os"
	"time"
)

func Getenv(variable string, defaultValue ...string) string {
	env := os.Getenv(variable)
	if env == "" {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		}
		return ""
	}
	return env
}

func IsLocal() bool {
	return os.Getenv("APP_ENV") == "" || os.Getenv("APP_ENV") == types.Local
}

func PointerString(s string) *string {
	return &s
}

func PointerTime(t time.Time) *time.Time {
	return &t
}

func PointerInt(i int) *int {
	return &i
}

func PointerInt64(i int64) *int64 {
	return &i
}
