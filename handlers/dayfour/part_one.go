package dayfour

import (
	"fmt"
	"strings"
)

const word = "XMAS"

func DayFourPartOneHandle(file []byte) int {
	lines := strings.Split(string(file), "\n")
	grid := []string{}
	for i := 0; i < len(lines); i++ {
		line := strings.TrimSpace(lines[i])
		if line == "" {
			continue
		}

		grid = append(grid, line)
	}
	result := 0

	result += searchHorizontally(grid)
	result += searchVertically(grid)
	result += searchDiagonally(grid)

	return result
}

func searchHorizontally(grid []string) int {
	x, y := 0, 0
	total := 0
	yBoundary := len(grid) - 1

	for y = 0; y <= yBoundary; y++ {
		xBoundary := len(grid[y]) - len(word)

		for x = 0; x <= xBoundary; x++ {
			w := buildWord(x, y, Neutral, Positive, word, grid)

			// left to right
			if w == word {
				total++
				continue
			}

			// right to left
			rw := reverseString(w)
			if rw == word {
				total++
			}
		}
	}

	return total
}

func searchVertically(grid []string) int {
	x, y := 0, 0
	total := 0
	yBoundary := len(grid) - len(word)
	xBoundary := len(grid[0])

	for x = 0; x < xBoundary; x++ {
		for y = 0; y <= yBoundary; y++ {
			w := buildWord(x, y, Positive, Neutral, word, grid)

			// top to bottom
			if w == word {
				total++
				continue
			}

			// bottom to top
			rw := reverseString(w)
			if rw == word {
				total++
			}
		}
	}

	return total
}

func searchDiagonally(grid []string) int {
	x, y := 0, 0
	total := 0
	yBoundary := len(grid) - len(word)
	xBoundary := len(grid[0]) - len(word)

	// top to bottom
	for x = 0; x <= xBoundary; x++ {
		for y = 0; y <= yBoundary; y++ {
			w := buildWord(x, y, Positive, Positive, word, grid)

			// top to bottom
			if w == word {
				fmt.Printf("Found word at x: %d, y: %d, dirRow: %d, dirCol: %d\n", x, y, Positive, Positive)
				total++
				continue
			}

			// bottom to top
			rw := reverseString(w)
			if rw == word {
				fmt.Printf("Found word at x: %d, y: %d, dirRow: %d, dirCol: %d\n", x, y, Positive, Positive)
				total++
			}

		}
	}

	yBoundary = len(grid) - yBoundary - 1
	// bottom to top
	for x := 0; x <= xBoundary; x++ {
		for y := len(grid) - 1; y >= yBoundary; y-- {
			w := buildWord(x, y, Negative, Positive, word, grid)

			// top to bottom
			if w == word {
				fmt.Printf("Found word at x: %d, y: %d, dirRow: %d, dirCol: %d\n", x, y, Negative, Positive)
				total++
				continue
			}

			// bottom to top
			rw := reverseString(w)
			if rw == word {
				fmt.Printf("Found word at x: %d, y: %d, dirRow: %d, dirCol: %d\n", x, y, Negative, Positive)
				total++
			}
		}
	}

	return total
}

const (
	Positive = 1 << iota
	Negative
	Neutral
)

func buildWord(
	col,
	row,
	rowDirection,
	colDirection int,
	word string,
	grid []string,
) string {
	y := row
	x := col

	if rowDirection == Neutral && colDirection == Neutral {
		panic("invalid direction")
	}

	w := ""

	for {
		if len(w) == len(word) {
			break
		}

		if len(grid) <= y {
			break
		}

		if len(grid[y]) <= x {
			break
		}

		w += string(grid[y][x])
		if rowDirection == Positive {
			y++
		} else if rowDirection == Negative {
			y--
		}

		if colDirection == Positive {
			x++
		} else if colDirection == Negative {
			x--
		}
	}

	return w
}

func reverseString(s string) string {
	r := ""
	for i := len(s) - 1; i >= 0; i-- {
		r += string(s[i])
	}

	return r
}
