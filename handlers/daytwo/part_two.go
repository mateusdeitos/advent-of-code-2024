package daytwo

import (
	"fmt"
	"strconv"
	"strings"
)

func DayTwoPartTwoHandle(file []byte) (int, error) {
	lines := strings.Split(string(file), "\n")
	safeLevelsCount := 0
	for i, line := range lines {
		safe, err := analyseLineWithTolerance(i, line)
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

func analyseLineWithTolerance(index int, line string) (bool, error) {
	levels, err := parseLevels(line)
	if err != nil {
		return false, fmt.Errorf("Error parsing line %d: %v", index, err)
	}

	if levels == nil {
		return false, nil
	}

	if isSafe(levels) {
		return true, nil
	}

	for i := 0; i < len(levels); i++ {
		levelsWithTolerance := createCombinationIgnoringLevel(levels, i)
		if isSafe(levelsWithTolerance) {
			return true, nil
		}
	}

	return false, nil
}

func isSafe(levels []int) bool {
	prev := levels[0]
	dir := getDirection(prev, levels[1])
	for i := 1; i < len(levels); i++ {
		level := levels[i]

		diff := level - prev
		safe := abs(diff) >= LOWER_THRESHOLD_LEVEL_SAFETY && abs(diff) <= UPPER_THRESHOLD_LEVEL_SAFETY
		if !safe {
			return false
		}

		currentDir := getDirection(prev, level)
		if currentDir != dir {
			return false
		}

		prev = level
	}

	return true
}

func parseLevels(line string) ([]int, error) {
	line = strings.TrimFunc(line, func(r rune) bool {
		return r == '\r' || r == '\n' || r == ' ' || r == '\t'
	})

	if line == "" {
		return nil, nil
	}

	parsedLine := strings.Split(string(line), " ")

	levels := make([]int, len(parsedLine))

	for i, strLevel := range parsedLine {
		level, err := strconv.Atoi(strLevel)
		if err != nil {
			return nil, fmt.Errorf("Error parsing line %s: %v", line, err)
		}

		levels[i] = level
	}

	return levels, nil
}

func analyseLevels(levels []int, onUnsafe func(index int)) {
	prev := levels[0]
	dir := getDirection(prev, levels[1])
	for i := 1; i < len(levels); i++ {
		level := levels[i]

		diff := level - prev
		safe := abs(diff) >= LOWER_THRESHOLD_LEVEL_SAFETY && abs(diff) <= UPPER_THRESHOLD_LEVEL_SAFETY
		if !safe {
			onUnsafe(i)
			onUnsafe(i - 1)
			continue
		}

		currentDir := getDirection(prev, level)
		if currentDir != dir {
			for j := i; j >= 0; j-- {
				onUnsafe(j)
			}
			continue
		}

		prev = level
	}

}

func createCombinationIgnoringLevel(levels []int, indexToIgnore int) []int {
	comb := []int{}

	for i, level := range levels {
		if i == indexToIgnore {
			continue
		}

		comb = append(comb, level)
	}

	return comb
}
