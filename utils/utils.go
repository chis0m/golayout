package utils

import (
	"fmt"
	"go-layout/internal/types"
	"os"
	"strings"
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

func GetDuration(variable string, defaultValue ...string) time.Duration {
	durationStr := Getenv(variable, defaultValue...)
	duration, err := time.ParseDuration(durationStr)
	if err != nil {
		fmt.Println("Error parsing duration:", err)
		return 0
	}
	return duration
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

func PointerUint64(i uint64) *uint64 {
	return &i
}

func CapitalizeFirst(s string) string {
	if s == "" {
		return ""
	}
	return strings.ToUpper(string(s[0])) + s[1:]
}
