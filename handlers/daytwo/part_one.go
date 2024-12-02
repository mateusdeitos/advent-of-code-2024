package daytwo

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

const (
	LOWER_THRESHOLD_LEVEL_SAFETY = 1
	UPPER_THRESHOLD_LEVEL_SAFETY = 3
)

const (
	DIRECTION_INCREASING = 1
	DIRECTION_DECREASING = -1
)

func DayTwoPartOneHandle(file []byte) (int, error) {
	lines := strings.Split(string(file), "\n")
	safeLevelsCount := 0
	for i, line := range lines {
		safe, err := analyseLine(i, line)
		if err != nil {
			return 0, err
		}

		if !safe {
			continue
		}

		safeLevelsCount++
	}

	return safeLevelsCount, nil
}

func analyseLine(index int, line string) (bool, error) {
	line = strings.TrimFunc(line, func(r rune) bool {
		return r == '\r' || r == '\n' || r == ' ' || r == '\t'
	})

	if line == "" {
		return false, nil
	}

	parsedLine := strings.Split(string(line), " ")

	prev, err := strconv.Atoi(parsedLine[0])
	if err != nil {
		return false, fmt.Errorf("Error parsing line %d: %v", index, err)
	}

	var dir int
	for i := 1; i < len(parsedLine); i++ {
		level, err := strconv.Atoi(parsedLine[i])
		if err != nil {
			return false, fmt.Errorf("Error parsing line %d: %v", index, err)
		}

		diff := level - prev
		if abs(diff) < LOWER_THRESHOLD_LEVEL_SAFETY || abs(diff) > UPPER_THRESHOLD_LEVEL_SAFETY {
			return false, nil
		}

		if i == 1 {
			dir = diff / abs(diff)
		} else if (diff > 0 && dir < 0) || (diff < 0 && dir > 0) {
			return false, nil
		}

		prev = level
	}

	return true, nil
}

func abs(a int) int {
	if a < 0 {
		return -a
	}

	return a
}

func isDifferenceSafe(a, b int) bool {
	diff := int(math.Abs(float64(a - b)))
	return diff >= LOWER_THRESHOLD_LEVEL_SAFETY && diff <= UPPER_THRESHOLD_LEVEL_SAFETY
}

func getDirection(a, b int) int {
	if a < b {
		return DIRECTION_INCREASING
	}

	return DIRECTION_DECREASING
}
