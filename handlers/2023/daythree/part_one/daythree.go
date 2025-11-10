package partone

import (
	"regexp"
	"strconv"
	"strings"
)

const DOT = "."

var numberRegex = regexp.MustCompile(`\d`)

func DaythreePartOneHandle(file []byte) int {
	lines := strings.Split(string(file), "\n")
	schematic := make([][]string, len(lines))
	for i := range lines {
		lines[i] = strings.Trim(lines[i], "\t")
		columns := strings.Split(lines[i], "")
		schematic[i] = columns
	}

	ns := []int{}
	for y, line := range schematic {
		for x := range schematic[y] {
			c := line[x]
			if !isSymbol(c) {
				continue
			}

			tl := getNumberTopLeft(schematic, x, y)
			if tl > 0 {
				ns = append(ns, tl)
			}

			t := getNumberAbove(schematic, x, y)
			if tl == 0 && t > 0 {
				ns = append(ns, t)
			}

			tr := getNumberTopRight(schematic, x, y)
			if t == 0 && tr > 0 {
				ns = append(ns, tr)
			}

			if l := getNumberToTheLeftOf(schematic, x, y); l > 0 {
				ns = append(ns, l)
			}

			if r := getNumberToTheRightOf(schematic, x, y); r > 0 {
				ns = append(ns, r)
			}

			bl := getNumberBottomLeft(schematic, x, y)
			if bl > 0 {
				ns = append(ns, bl)
			}

			b := getNumberBelow(schematic, x, y)
			if bl == 0 && b > 0 {
				ns = append(ns, b)
			}

			br := getNumberBottomRight(schematic, x, y)
			if b == 0 && br > 0 {
				ns = append(ns, br)
			}
		}
	}

	sum := 0
	for _, n := range ns {
		sum += n
	}

	return sum
}

func isSymbol(char string) bool {
	return char != DOT && !numberRegex.Match([]byte(char))
}

func getNumberToTheRightOf(schematic [][]string, x, y int) int {
	x1 := x + 1
	if !validateCoords(schematic, x1, y) {
		return 0
	}

	line := schematic[y]

	var n string
	for _, c := range line[x1:] {
		if !numberRegex.Match([]byte(c)) {
			break
		}

		n += c
	}

	intN, err := strconv.Atoi(n)
	if err != nil {
		return 0
	}

	return intN
}

func getNumberToTheLeftOf(schematic [][]string, x, y int) int {
	x1 := x - 1
	if !validateCoords(schematic, x1, y) {
		return 0
	}

	line := schematic[y]

	var n string
	for i := x1; i >= 0; i-- {
		c := line[i]
		if !numberRegex.Match([]byte(c)) {
			break
		}

		n = c + n
	}

	intN, err := strconv.Atoi(n)
	if err != nil {
		return 0
	}

	return intN
}

func getNumberBelow(schematic [][]string, x, y int) int {
	if !validateCoords(schematic, x, y) {
		return 0
	}

	if y+1 > len(schematic)-1 {
		return 0
	}

	line := schematic[y+1]

	if !numberRegex.Match([]byte(line[x])) {
		return 0
	}

	startingIndex := x
	for i := x; i > 0; i-- {
		c := line[i]

		if !numberRegex.Match([]byte(c)) {
			break
		}

		startingIndex = i
	}

	if startingIndex < 0 {
		return 0
	}

	var n string
	for i := startingIndex; i < len(line)-1; i++ {
		c := line[i]

		if !numberRegex.Match([]byte(c)) {
			break
		}

		n += c
	}

	intN, err := strconv.Atoi(n)
	if err != nil {
		return 0
	}

	return intN
}

func getNumberAbove(schematic [][]string, x, y int) int {
	if !validateCoords(schematic, x, y) {
		return 0
	}

	if y-1 < 0 {
		return 0
	}

	line := schematic[y-1]

	if !numberRegex.Match([]byte(line[x])) {
		return 0
	}

	startingIndex := x
	for i := x; i > 0; i-- {
		c := line[i]
		if !numberRegex.Match([]byte(c)) {
			break
		}

		startingIndex = i
	}

	if startingIndex < 0 {
		return 0
	}

	var n string
	for i := startingIndex; i < len(line)-1; i++ {
		c := line[i]

		if !numberRegex.Match([]byte(c)) {
			break
		}

		n += c
	}

	intN, err := strconv.Atoi(n)
	if err != nil {
		return 0
	}

	return intN
}

func getNumberTopLeft(schematic [][]string, x, y int) int {
	if !validateCoords(schematic, x, y) {
		return 0
	}

	x1 := x - 1
	if !validateCoords(schematic, x1, y-1) {
		return 0
	}

	line := schematic[y-1]

	if !numberRegex.Match([]byte(line[x1])) {
		return 0
	}

	startingIndex := x1
	for i := x1; i >= 0; i-- {
		c := line[i]
		if !numberRegex.Match([]byte(c)) {
			break
		}

		startingIndex = i
	}

	if startingIndex < 0 {
		return 0
	}

	var n string
	for i := startingIndex; i < len(line)-1; i++ {
		c := line[i]

		if !numberRegex.Match([]byte(c)) {
			break
		}

		n += c
	}

	intN, err := strconv.Atoi(n)
	if err != nil {
		return 0
	}

	return intN
}

func getNumberTopRight(schematic [][]string, x, y int) int {
	if !validateCoords(schematic, x, y) {
		return 0
	}

	x1 := x + 1
	if !validateCoords(schematic, x1, y-1) {
		return 0
	}

	line := schematic[y-1]

	if !numberRegex.Match([]byte(line[x1])) {
		return 0
	}

	var n string
	for i := x1; i <= len(line)-1; i++ {
		c := line[i]

		if !numberRegex.Match([]byte(c)) {
			break
		}

		n += c
	}

	intN, err := strconv.Atoi(n)
	if err != nil {
		return 0
	}

	return intN
}

func getNumberBottomLeft(schematic [][]string, x, y int) int {
	if !validateCoords(schematic, x, y) {
		return 0
	}

	x1 := x - 1
	if !validateCoords(schematic, x1, y+1) {
		return 0
	}

	line := schematic[y+1]

	if !numberRegex.Match([]byte(line[x1])) {
		return 0
	}

	startingIndex := x
	for i := x1; i >= 0; i-- {
		c := line[i]

		if !numberRegex.Match([]byte(c)) {
			break
		}

		startingIndex = i
	}

	if startingIndex < 0 {
		return 0
	}

	var n string
	for i := startingIndex; i < len(line)-1; i++ {
		c := line[i]

		if !numberRegex.Match([]byte(c)) {
			break
		}

		n += c
	}

	intN, err := strconv.Atoi(n)
	if err != nil {
		return 0
	}

	return intN
}

func getNumberBottomRight(schematic [][]string, x, y int) int {
	if !validateCoords(schematic, x, y) {
		return 0
	}

	x1 := x + 1
	if !validateCoords(schematic, x1, y+1) {
		return 0
	}

	line := schematic[y+1]

	if !numberRegex.Match([]byte(line[x1])) {
		return 0
	}

	var n string
	for i := x1; i <= len(line)-1; i++ {
		c := line[i]

		if !numberRegex.Match([]byte(c)) {
			break
		}

		n += c
	}

	intN, err := strconv.Atoi(n)
	if err != nil {
		return 0
	}

	return intN
}

func validateCoords(schematic [][]string, x, y int) bool {
	maxX := len(schematic[0])
	if x >= maxX {
		return false
	}

	if y > len(schematic)-1 {
		return false
	}

	if y < 0 {
		return false
	}

	return true
}
