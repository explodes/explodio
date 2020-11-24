package stand

import (
	"fmt"
	"os"
)

func RequireEnv(key string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		panic(fmt.Sprintf("environment variable %s is required but was not found", key))
	}
	return value
}

func DefaultEnv(key, defaultValue string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		return defaultValue
	}
	return value
}
