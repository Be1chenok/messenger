package defaultEnv

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

var ErrUnknownTimeUnit = errors.New("unknown time unit (available time units: ns, us, ms, s, m, h, d)")

func GetString(key string, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}

	return value
}

func GetInt(key string, defaultValue int) (int, error) {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue, nil
	}

	val, err := strconv.Atoi(value)
	if err != nil {
		return 0, fmt.Errorf("failed to get %s from env: invalid number %s, err: %w", key, value, err)
	}

	return val, nil
}

func GetDuration(key string, defaultValue time.Duration) (time.Duration, error) {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue, nil
	}

	var total time.Duration

	parts := strings.Fields(value)
	for _, part := range parts {
		var numStr string
		var unit string
		for i, r := range part {
			if r >= '0' && r <= '9' || r == '-' {
				numStr += string(r)
				continue
			}
			unit = part[i:]
			break
		}

		num, err := strconv.Atoi(numStr)
		if err != nil {
			return 0, fmt.Errorf("failed to get %s from env: invalid number: %s, err: %w", key, numStr, err)
		}

		switch unit {
		case "ns":
			total += time.Duration(num) * time.Nanosecond
		case "us":
			total += time.Duration(num) * time.Microsecond
		case "ms":
			total += time.Duration(num) * time.Millisecond
		case "s":
			total += time.Duration(num) * time.Second
		case "m":
			total += time.Duration(num) * time.Minute
		case "h":
			total += time.Duration(num) * time.Hour
		case "d":
			total += time.Duration(num) * time.Hour * 24
		default:
			return 0, fmt.Errorf("failed to get %s from env: invalid time unit: %s, err: %w", key, unit, ErrUnknownTimeUnit)
		}
	}

	return total, nil
}
