package dayten

import (
	"fmt"
	"strconv"
	"strings"
)

func DaytenPartTwoHandle(file []byte) int {
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
				trails = removeRepeatedTrais(trails)
				score := len(trails)
				sum += score
				fmt.Printf("Trail ending at (%d, %d) has score %d\n", x, y, score)
			}
		}
	}

	return sum
}

func removeRepeatedTrais(trails []Trail) []Trail {
	record := map[string]bool{}
	var filteredTrails []Trail
	for _, t := range trails {
		trailStr := t.String()
		_, ok := record[trailStr]
		if ok {
			continue
		}

		record[trailStr] = true
		filteredTrails = append(filteredTrails, t)
	}

	return filteredTrails
}

type Trail struct {
	Points []Point
}

func Copy(trail Trail) Trail {
	return Trail{
		Points: append([]Point{}, trail.Points...),
	}
}

func (t *Trail) append(p Point) {
	t.Points = append(t.Points, p)
}

func (t *Trail) IsEqual(other *Trail) bool {
	if len(t.Points) != len(other.Points) {
		return false
	}

	for i, p := range t.Points {
		if !p.Equals(&other.Points[i]) {
			return false
		}
	}

	return true
}

func (t *Trail) String() string {
	var s string
	for _, p := range t.Points {
		s += p.String()
	}

	return s
}

type Point struct {
	Value int
	X     int
	Y     int
}

func newPoint(x, y int, grid []string) *Point {
	v := getValueAtPos(x, y, grid)
	if v == -1 {
		return nil
	}

	return &Point{
		Value: v,
		X:     x,
		Y:     y,
	}
}

func (p *Point) String() string {
	return fmt.Sprintf("(%d, %d)", p.X, p.Y)
}

func (p *Point) Equals(other *Point) bool {
	return p.X == other.X && p.Y == other.Y
}

func (p *Point) Left(grid []string) *Point {
	return newPoint(p.X-1, p.Y, grid)
}

func (p *Point) Right(grid []string) *Point {
	return newPoint(p.X+1, p.Y, grid)
}

func (p *Point) Up(grid []string) *Point {
	return newPoint(p.X, p.Y-1, grid)
}

func (p *Point) Down(grid []string) *Point {
	return newPoint(p.X, p.Y+1, grid)
}

func getTrailsStartingAt(previous, p *Point, grid []string, trail Trail) []Trail {
	points := []*Point{
		p.Left(grid),
		p.Down(grid),
		p.Right(grid),
		p.Up(grid),
	}

	var trails []Trail

	for _, point := range points {
		if point == nil {
			continue
		}

		if previous != nil && previous.Equals(point) {
			continue
		}

		if point.Value-p.Value != 1 {
			continue
		}

		prevTrail := Copy(trail)

		trail.append(*point)

		if point.Value == toInt(trailEnd) {
			trails = append(trails, trail)
			trail = prevTrail
			continue
		}

		ts := getTrailsStartingAt(p, point, grid, trail)
		trails = append(trails, ts...)
		trail = prevTrail
	}

	return trails
}

func getValueAtPos(x, y int, grid []string) int {
	if y < 0 || y >= len(grid) {
		return -1
	}

	line := grid[y]
	if x < 0 || x >= len(line) {
		return -1
	}

	c := grid[y][x]
	if string(c) == "A" {
		return -1
	}

	return toInt(string(c))
}

func toInt(s string) int {
	v, _ := strconv.Atoi(s)
	return v
}
