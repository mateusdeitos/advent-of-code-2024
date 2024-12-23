package dayeleven

import (
	"strconv"
	"strings"
)

func DayelevenPartOneHandle(file []byte) int {
	line := strings.Split(string(file), " ")

	blinks := 25
	for blinks > 0 {
		total := len(line)
		skip := false
		for i := 0; i < total; i++ {
			if skip {
				skip = false
				continue
			}

			if line[i] == "0" {
				line[i] = "1"
				continue
			}

			if len(line[i])%2 == 0 {
				v1, v2 := split(line[i])
				line = append(line[:i], append([]string{v1, v2}, line[i+1:]...)...)
				skip = true
				total++
				continue
			}

			v, err := strconv.Atoi(line[i])
			if err != nil {
				panic(err)
			}

			line[i] = strconv.Itoa(v * 2024)
		}

		blinks--
	}

	return len(line)
}

func split(s string) (string, string) {
	var v1, v2 string
	v1 = s[:len(s)/2]
	v2 = s[len(s)/2:]

	v2Int, err := strconv.Atoi(v2)
	if err != nil {
		panic(err)
	}

	v2 = strconv.Itoa(v2Int)

	return v1, v2
}
