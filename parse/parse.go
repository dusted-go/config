package parse

import (
	"strconv"
	"time"
)

// Bool parses a string into a boolean value.
// If the value cannot be parsed then the default will be returned.
func Bool(val string, def bool) bool {
	b, err := strconv.ParseBool(val)
	if err != nil {
		return def
	}
	return b
}

// Int parses a string into a int value.
// If the value cannot be parsed then the default will be returned.
func Int(val string, def int) int {
	i, err := strconv.Atoi(val)
	if err != nil {
		return def
	}
	return i
}

// Duration parses a string into a duration.
// If the value cannot be parsed then the default will be returned.
func Duration(val string, def time.Duration) time.Duration {
	d, err := time.ParseDuration(val)
	if err != nil {
		return def
	}
	return d
}
