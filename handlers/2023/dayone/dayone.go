package dayone

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var regexpExtractNumbers = regexp.MustCompile(`\d`)

func DayonePartOneHandle(file []byte) int {
	lines := strings.Split(string(file), "\n")
	sum := 0
	for i := range lines {
		fmt.Printf("[%d] - line %s ", i, lines[i])
		matches := regexpExtractNumbers.FindAllString(lines[i], -1)
		if len(matches) == 0 {
			continue
		}

		var fd, ld string
		for i := range matches {
			n, err := strconv.Atoi(matches[i])
			if err != nil {
				continue
			}

			if fd == "" {
				fd = strconv.Itoa(n)
			} else {
				ld = strconv.Itoa(n)
			}
		}

		if ld == "" {
			ld = fd
		}

		fmt.Printf("- digits: %s", fd+ld)

		n, err := strconv.Atoi(fd + ld)
		if err != nil {
			continue
		}

		fmt.Printf(" - converted: %d", n)

		sum += n

		fmt.Printf(" - sum: %d\n", sum)
	}

	return sum
}
