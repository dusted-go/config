package env

import (
	"os"
	"strconv"
	"time"
)

// GetOrDefault returns the value of a given environment variable
// or the fallback value if the environment variable was not present.
func GetOrDefault(key string, fallback string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return fallback
}

// GetBoolOrDefault returns the bool value of a given environment variable
// or the fallback value if the environment variable was not set.
// If the environment variable was set but its value cannot be parsed into
// a bool then the function will panic.
func GetBoolOrDefault(key string, fallback bool) bool {
	if val, ok := os.LookupEnv(key); ok {
		b, err := strconv.ParseBool(val)
		if err != nil {
			panic(err)
		}
		return b
	}
	return fallback
}

// GetIntOrDefault returns the int value of a given environment variable
// or the fallback value if the environment variable was not set.
// If the environment variable was set but its value cannot be parsed
// into an integer then the function will panic.
func GetIntOrDefault(key string, fallback int) int {
	if val, ok := os.LookupEnv(key); ok {
		i, err := strconv.Atoi(val)
		if err != nil {
			panic(err)
		}
		return i
	}
	return fallback
}

// GetDurationOrDefault returns the time.Duration value of a given environment
// variable or the fallback value if the environment variable was not set.
// If the environment variable was set but its value cannot be parsed into a
// time.Duration object then the function will panic.
func GetDurationOrDefault(key string, fallback time.Duration) time.Duration {
	if val, ok := os.LookupEnv(key); ok {
		d, err := time.ParseDuration(val)
		if err != nil {
			panic(err)
		}
		return d
	}
	return fallback
}
