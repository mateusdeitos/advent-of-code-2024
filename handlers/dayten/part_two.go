package dayten

import (
	"fmt"
	"strings"
)

const trailStart string = "0"
const trailEnd string = "9"

func DaytenPartOneHandle(file []byte) int {
	grid := strings.Split(string(file), "\n")

	sum := 0
	for y := range grid {
		for x := range grid[y] {
			c := string(grid[y][x])
			if c == trailStart {
				p := newPoint(x, y, grid)
				trail := Trail{}
				trail.append(*p)
				trails := getTrailsStartingAt(nil, p, grid, trail)
				trails = keepUniqueTrails(trails)
				score := len(trails)
				sum += score
				fmt.Printf("Trail ending at (%d, %d) has score %d\n", x, y, score)
			}
		}
	}

	return sum
}

func keepUniqueTrails(trails []Trail) []Trail {
	record := map[string]bool{}
	var filteredTrails []Trail
	for _, t := range trails {
		trailStr := t.Points[len(t.Points)-1].String()
		_, ok := record[trailStr]
		if ok {
			continue
		}

		record[trailStr] = true
		filteredTrails = append(filteredTrails, t)
	}

	return filteredTrails
}
