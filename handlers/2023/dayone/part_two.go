package dayone

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var regexpExtractNumbersAndNumericStrings = regexp.MustCompile(`\d|one|two|three|four|five|six|seven|eight|nine`)
var regexpExtractNumericStrings = regexp.MustCompile(`one|two|three|four|five|six|seven|eight|nine`)

var conversionDict = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func adjustOverlaps(line string) string {
	numbers := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

	overlaps := map[string]string{}
	for i, n1 := range numbers {
		if i == len(numbers)-1 {
			continue
		}

		for _, n2 := range numbers[i+1:] {
			n1FirstDigit := string(n1[0])
			n2FirstDigit := string(n2[0])
			n1LastDigit := string(n1[len(n1)-1])
			n2LastDigit := string(n2[len(n2)-1])
			if n1FirstDigit == n2LastDigit {
				key := n2 + n1[1:]
				overlaps[key] = n2 + n1
			}

			if n1LastDigit == n2FirstDigit {
				key := n1 + n2[1:]
				overlaps[key] = n1 + n2
			}
		}
	}

	for key, value := range overlaps {
		line = strings.ReplaceAll(line, key, value)
	}

	return line
}

func DayonePartTwoHandle(file []byte) int {
	lines := strings.Split(string(file), "\n")
	sum := 0
	for i := range lines {
		line := lines[i]
		fmt.Printf("[%d] - line %s", i, line)
		line = adjustOverlaps(line)
		if lines[i] != line {
			fmt.Printf(" - adjusted overlap: %s", line)
		}
		matches := regexpExtractNumbersAndNumericStrings.FindAllString(line, -1)
		if len(matches) == 0 {
			continue
		}

		var fd, ld string
		for i := range matches {
			var n int
			if x, ok := conversionDict[matches[i]]; !ok {
				m, err := strconv.Atoi(matches[i])
				if err != nil {
					continue
				}

				n = m
			} else {
				n = x
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

		fmt.Printf(" - digits: %s", fd+ld)

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
