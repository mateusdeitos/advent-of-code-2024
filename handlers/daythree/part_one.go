package daythree

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var mulRegexp = regexp.MustCompile(`(mul\(\d+,\d+\))`)
var digitsRegex = regexp.MustCompile(`(\d+,\d+)`)

func DayThreePartOneHandle(file []byte) (int, error) {
	lines := string(file)
	m := mulRegexp.FindAllStringSubmatch(lines, -1)
	r := 0
	for i, matches := range m {
		if len(matches) < 1 {
			return 0, fmt.Errorf("error parsing match %d", i)
		}

		line := matches[0]

		v, err := multiply(i, line)
		if err != nil {
			return 0, err
		}

		r += v
	}

	return r, nil

}

func multiply(i int, line string) (int, error) {
	digits := digitsRegex.FindAllString(line, -1)
	if len(digits) != 1 {
		return 0, fmt.Errorf("error parsing digits of match %d: %v", i, line)
	}

	digits = strings.Split(digits[0], ",")
	if len(digits) != 2 {
		return 0, fmt.Errorf("error parsing digits of match %d: %v", i, line)
	}

	d1, d2 := digits[0], digits[1]
	v1, err := strconv.Atoi(d1)
	if err != nil {
		return 0, fmt.Errorf("error parsing line %d: %v", i, line)
	}
	v2, err := strconv.Atoi(d2)
	if err != nil {
		return 0, fmt.Errorf("error parsing line %d: %v", i, line)
	}

	return v1 * v2, nil
}
