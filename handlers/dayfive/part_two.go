package dayfive

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func DayFivePartTwoHandle(file []byte) int {
	lists := strings.Split(string(file), "\n\n")

	pageOrderingRules := strings.Split(lists[0], "\n")
	nodeMap, err := createNodeMapFromList(pageOrderingRules)
	if err != nil {
		panic(err)
	}

	pageUpdates := strings.Split(lists[1], "\n")
	sumOfMiddleValuesForIncorrectUpdates := 0
	for _, u := range pageUpdates {
		strValues := strings.Split(u, ",")
		values := make([]int, len(strValues))
		for i, v := range strValues {
			value, _ := strconv.Atoi(v)
			values[i] = value
		}

		correct, err := isUpdateSequenceCorrect(*nodeMap, values)
		if err != nil {
			panic(err)
		}

		if correct {
			continue
		}

		values = reorderIncorrectUpdateSequence(*nodeMap, values)

		correct, err = isUpdateSequenceCorrect(*nodeMap, values)
		if err != nil {
			panic(err)
		}

		if !correct {
			panic(fmt.Sprintf("incorrect update sequence %v", values))
		}

		i := int(len(values) / 2)
		middleValue := values[i]
		sumOfMiddleValuesForIncorrectUpdates += middleValue
	}

	return sumOfMiddleValuesForIncorrectUpdates
}

func reorderIncorrectUpdateSequence(nodeMap NodeMap, values []int) []int {
	slices.SortFunc(values, func(a, b int) int {
		na, nb := nodeMap[a], nodeMap[b]
		if na.isBeforeOf(nb.value) {
			return -1
		}

		if nb.isBeforeOf(na.value) {
			return 1
		}

		return 0
	})

	return values
}
