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
	DIRECTION_INCREASING = iota
	DIRECTION_DECREASING
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

	var prev *int
	dir := DIRECTION_INCREASING
	for i, strLevel := range parsedLine {
		level, err := strconv.Atoi(strLevel)
		if err != nil {
			return false, fmt.Errorf("Error parsing line %d: %v", index, err)
		}

		if i == 1 {
			dir = getDirection(*prev, level)
		}

		if i > 0 {
			safe := isSafe(*prev, level)
			if !safe {
				return false, nil
			}

			currentDir := getDirection(*prev, level)
			if currentDir != dir {
				return false, nil
			}
		}

		prev = &level
	}

	return true, nil
}

func isSafe(a, b int) bool {
	diff := int(math.Abs(float64(a - b)))
	return diff >= LOWER_THRESHOLD_LEVEL_SAFETY && diff <= UPPER_THRESHOLD_LEVEL_SAFETY
}

func getDirection(a, b int) int {
	if a < b {
		return DIRECTION_INCREASING
	}

	return DIRECTION_DECREASING
}
