package dayeight

import (
	"strings"
)

type AntinodesTracker map[int]map[int]bool

func (ant AntinodesTracker) Set(x, y int) {
	if _, ok := ant[x]; !ok {
		ant[x] = map[int]bool{}
	}

	if _, ok := ant[x][y]; !ok {
		ant[x][y] = true
	}
}

func (ant AntinodesTracker) Len() int {
	total := 0
	for k := range ant {
		total += len(ant[k])
	}

	return total
}

func DayEightPartOneHandle(file []byte) int {
	lines := strings.Split(string(file), "\n")
	grid := make([][]string, len(lines))
	for i := range lines {
		cells := strings.Split(lines[i], "")
		grid[i] = append(grid[i], cells...)
	}

	antennas := map[string][][2]int{}

	for y, cells := range grid {
		for x := range cells {
			s := cells[x]
			if s == "." {
				continue
			}

			if _, ok := antennas[s]; !ok {
				antennas[s] = [][2]int{}
			}

			antennas[s] = append(antennas[s], [2]int{x, y})
		}
	}

	antinodes := make(AntinodesTracker)
	for _, coord := range antennas {
		if len(coord) <= 1 {
			continue
		}

		for i := 0; i < len(coord); i++ {
			for j := 0; j < len(coord); j++ {
				if j == i {
					continue
				}

				evaluateAntinodes(antinodes, grid, coord[i], coord[j])
			}
		}
	}

	return antinodes.Len()
}

func evaluateAntinodes(antinodes AntinodesTracker, grid [][]string, a, b [2]int) {
	aDiffX := a[0] - b[0]
	aDiffY := a[1] - b[1]

	x0 := a[0] + aDiffX
	y0 := a[1] + aDiffY

	if isWithinBounds(grid, x0, y0) {
		antinodes.Set(x0, y0)
	}

	bDiffX := b[0] - a[0]
	bDiffY := b[1] - a[1]

	x1 := b[0] + bDiffX
	y1 := b[1] + bDiffY

	if isWithinBounds(grid, x1, y1) {
		antinodes.Set(x1, y1)
	}
}

func isWithinBounds(grid [][]string, x, y int) bool {
	yBound := len(grid) - 1
	xBound := len(grid[0]) - 1

	if y < 0 || y > yBound {
		return false
	}

	if x < 0 || x > xBound {
		return false
	}

	return true
}
