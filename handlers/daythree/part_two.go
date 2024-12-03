package daythree

import (
	"fmt"
	"regexp"
)

var mulEnablingRegexp = regexp.MustCompile(`(mul\(\d+,\d+\))|(do(n't)?\(\))`)

func DayThreePartTwoHandle(file []byte) (int, error) {
	lines := string(file)
	m := mulEnablingRegexp.FindAllStringSubmatch(lines, -1)
	r := 0
	enabled := true
	for i, matches := range m {
		if len(matches) < 1 {
			return 0, fmt.Errorf("error parsing match %d", i)
		}

		line := matches[0]
		if line == "don't()" {
			enabled = false
			continue
		} else if line == "do()" {
			enabled = true
			continue
		} else if !enabled {
			continue
		}

		v, err := multiply(i, line)
		if err != nil {
			return 0, err
		}

		r += v
	}

	return r, nil

}
