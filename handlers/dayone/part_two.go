package dayone

import (
	"fmt"
	"strings"
)

// DayOnePartTwoHandle calculates the similarity score between the two lists
// It compares how many times a number appears on the right list and multiply by its value
// e.g: if 3 appears 2 times in the right list, the similarity score will be 6
func DayOnePartTwoHandle(file []byte) (int, error) {
	score := map[int][2]int{}
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
			score[l] = [2]int{1, 0}
		} else {
			// otherwise:
			// 1 - calculate the previous similarity score
			// 2 - increase the count
			// 3 - subtract the previous score from total
			// 4 - add the new score to the total
			s := score[l]
			prevScore := calcSimilarity(s)
			s[0]++
			score[l] = s
			total -= prevScore
			total += calcSimilarity(s)
		}

		// if the right list value is in the score map
		// 1 - calculate the previous similarity score
		// 2 - add the right list value to the value
		// 3 - subtract the previous score from total
		// 4 - add the new score to the total
		if s, ok := score[r]; ok {
			prevScore := calcSimilarity(s)
			s[1] += r
			score[r] = s
			total -= prevScore
			total += calcSimilarity(s)
		} else {
			// otherwise:
			// add the right list value to the score map with 0 count
			score[r] = [2]int{0, r}
		}
	}

	return total, nil
}

func calcSimilarity(set [2]int) int {
	return set[0] * set[1]
}
