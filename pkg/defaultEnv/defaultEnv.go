package defaultEnv

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func GetString(key string, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}

	return value
}

func GetBool(key string, defaultValue bool) (bool, error) {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue, nil
	}

	val, err := strconv.ParseBool(value)
	if err != nil {
		return false, fmt.Errorf("env %s: invalid bool value %s, err: %w", key, value, err)
	}

	return val, nil
}

func GetInt(key string, defaultValue int) (int, error) {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue, nil
	}

	val, err := strconv.Atoi(value)
	if err != nil {
		return 0, fmt.Errorf("env %s: invalid number %s, err: %w", key, value, err)
	}

	return val, nil
}

func GetDuration(key string, defaultValue time.Duration) (time.Duration, error) {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue, nil
	}

	val, err := time.ParseDuration(value)
	if err != nil {
		return 0, fmt.Errorf("env %s: invalid duration %s, err: %w", key, value, err)
	}

	return val, nil
}
