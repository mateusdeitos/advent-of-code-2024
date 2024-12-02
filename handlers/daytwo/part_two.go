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
	line = strings.TrimFunc(line, func(r rune) bool {
		return r == '\r' || r == '\n' || r == ' ' || r == '\t'
	})

	if line == "" {
		return false, nil
	}

	parsedLine := strings.Split(string(line), " ")

	badLevelsCombination := [][]int{}
	levels := make([]int, len(parsedLine))

	for i, strLevel := range parsedLine {
		level, err := strconv.Atoi(strLevel)
		if err != nil {
			return false, fmt.Errorf("Error parsing line %d: %v", index, err)
		}

		levels[i] = level
	}

	fmt.Printf("Analyzing line %d: %v\t", index, levels)
	analyseLevels(levels, func(index int) {
		badLevelsCombination = append(badLevelsCombination, createCombinationIgnoringLevel(levels, index))
	})

	if len(badLevelsCombination) == 0 {
		fmt.Printf("RESULT: safe\n")
		return true, nil
	}

	for _, combination := range badLevelsCombination {
		safe := true
		analyseLevels(combination, func(index int) {
			safe = false
		})

		if safe {
			fmt.Printf("RESULT: safe with tolerance\n")
			return true, nil
		}
	}

	fmt.Printf("RESULT: unsafe\n")

	return false, nil
}

func analyseLevels(levels []int, onUnsafe func(index int)) {
	prev := levels[0]
	dir := getDirection(prev, levels[1])
	for i := 1; i < len(levels); i++ {
		level := levels[i]

		safe := isDifferenceSafe(prev, level)
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
