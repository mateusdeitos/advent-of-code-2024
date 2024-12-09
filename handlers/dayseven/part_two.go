package dayseven

import (
	"strconv"
	"strings"
)

func DaySevenPartTwoHandle(file []byte) int64 {
	paths := strings.Split(string(file), "\n")
	data := make([]Line, len(paths))
	for i := range paths {
		parts := strings.Split(paths[i], ": ")
		result, err := strconv.ParseInt(parts[0], 10, 64)
		if err != nil {
			panic(err)
		}

		vStr := strings.Split(parts[1], " ")
		values := make([]int, len(vStr))
		for j := range vStr {
			v, err := strconv.Atoi(vStr[j])
			if err != nil {
				panic(err)
			}
			values[j] = v
		}

		data[i] = Line{Result: result, Values: values}
	}

	total := int64(0)
	for _, line := range data {
		combs := generateCombinations(line.Values)
		for _, comb := range combs {
			r := comb.Eval()
			if r == line.Result {
				total += r
				break
			}
		}
	}

	return total
}
