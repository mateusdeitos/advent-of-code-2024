package dayone

import (
	"fmt"
	"strings"
)

// DayOnePartTwoHandle calculates the similarity score between the two lists
// It compares how many times a number appears on the right list and multiply by its value
// e.g: if 3 appears 2 times in the right list, the similarity score will be 6
func DayOnePartTwoHandle(file []byte) (int, error) {
	score := map[int]*[2]int{}
	total := 0

	lines := strings.Split(string(file), "\n")
	for i, line := range lines {
		n, err := parseLineAndExtractNumbers(i, line)
		if err != nil {
			return 0, fmt.Errorf("error parsing line %d: %v", i, err)
		}

		if n == nil {
			continue
		}

		l := n[0]
		r := n[1]

		// if the left list value is not in the score map
		// add to it with count = 1
		if _, ok := score[l]; !ok {
			score[l] = &[2]int{1, 0}
		} else {
			total = incrementTotal(total, 1, l, 0, score)
		}

		if _, ok := score[r]; ok {
			total = incrementTotal(total, r, r, 1, score)
		} else {
			// add the right list value to the score map with 0 count
			score[r] = &[2]int{0, r}
		}
	}

	return total, nil
}

func incrementTotal(total, increment, key, index int, score map[int]*[2]int) int {
	prev := calcSimilarity(score[key])
	score[key][index] += increment
	total -= prev
	total += calcSimilarity(score[key])
	return total
}

func calcSimilarity(set *[2]int) int {
	return set[0] * set[1]
}
