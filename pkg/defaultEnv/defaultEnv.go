package defaultEnv

import (
	"errors"
	"fmt"
	"os"
	"regexp"
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
		return 0, fmt.Errorf("env %s: invalid number %s, err: %w", key, value, err)
	}

	return val, nil
}

func GetDuration(key string, defaultValue time.Duration) (time.Duration, error) {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue, nil
	}

	globalSign := 1
	posValue := value
	if strings.HasPrefix(value, "-") {
		globalSign = -1
		posValue = strings.TrimPrefix(value, "-")
	}

	re := regexp.MustCompile(`^(\d+[a-zA-Z]+)+$`)
	if !re.MatchString(posValue) {
		return 0, fmt.Errorf("env %s: invalid duration format: %s", key, value)
	}

	componentRe := regexp.MustCompile(`(\d+)([a-zA-Z]+)`)
	matches := componentRe.FindAllStringSubmatch(posValue, -1)
	if matches == nil {
		return 0, fmt.Errorf("env %s: invalid duration format: %s", key, value)
	}

	var total time.Duration

	for _, match := range matches {
		numStr, unit := match[1], match[2]
		num, err := strconv.Atoi(numStr)
		if err != nil {
			return 0, fmt.Errorf("env %s: invalid integer: %s, err: %w", key, numStr, err)
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
			total += time.Duration(num) * 24 * time.Hour
		default:
			return 0, fmt.Errorf("env %s: invalid time unit: %s, err: %w", key, unit, ErrUnknownTimeUnit)
		}
	}

	if globalSign < 0 {
		total = -total
	}

	return total, nil
}
