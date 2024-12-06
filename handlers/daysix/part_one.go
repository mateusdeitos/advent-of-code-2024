package daysix

import (
	"errors"
	"strings"
)

const obstacle = "#"

type GuardFacingDirection string

const (
	UP    GuardFacingDirection = "^"
	DOWN  GuardFacingDirection = "v"
	LEFT  GuardFacingDirection = "<"
	RIGHT GuardFacingDirection = ">"
)

type Guard struct {
	facingDirection GuardFacingDirection
	coordinates     [2]int
}

type GuardMovementTracker struct {
	coordinates map[int]map[int]bool
}

func (gmt *GuardMovementTracker) Track(c [2]int) {
	x, y := c[0], c[1]
	if gmt.coordinates == nil {
		gmt.coordinates = make(map[int]map[int]bool)
	}

	if _, ok := gmt.coordinates[x]; !ok {
		gmt.coordinates[x] = make(map[int]bool)
	}

	gmt.coordinates[x][y] = true
}

func (gmt *GuardMovementTracker) Len() int {
	l := 0
	for k := range gmt.coordinates {
		l += len(gmt.coordinates[k])
	}

	return l
}

func newGuard(m [][]string) (*Guard, error) {
	for y, line := range m {
		for x, cell := range line {
			coordinates := [2]int{x, y}
			switch GuardFacingDirection(cell) {
			case UP:
				return &Guard{UP, coordinates}, nil
			case DOWN:
				return &Guard{DOWN, coordinates}, nil
			case LEFT:
				return &Guard{LEFT, coordinates}, nil
			case RIGHT:
				return &Guard{RIGHT, coordinates}, nil
			default:
				continue
			}
		}
	}

	return nil, errors.New("Guard not found")
}

func (g *Guard) Move(m [][]string) bool {
	c := g.coordinates
	switch g.facingDirection {
	case UP:
		g.coordinates[1]--
	case DOWN:
		g.coordinates[1]++
	case LEFT:
		g.coordinates[0]--
	case RIGHT:
		g.coordinates[0]++
	}

	if g.hasObstacle(m) {
		g.coordinates = c
		g.changeDirection()
		return true
	}

	return !g.isOutOfBound(m)
}

func (g *Guard) changeDirection() {
	switch g.facingDirection {
	case UP:
		g.facingDirection = RIGHT
	case RIGHT:
		g.facingDirection = DOWN
	case DOWN:
		g.facingDirection = LEFT
	case LEFT:
		g.facingDirection = UP
	}
}

func (g *Guard) hasObstacle(m [][]string) bool {
	if g.isOutOfBound(m) {
		return false
	}

	return m[g.coordinates[1]][g.coordinates[0]] == obstacle
}

func (g *Guard) isOutOfBound(m [][]string) bool {
	if g.coordinates[0] < 0 || g.coordinates[1] < 0 {
		return true
	}

	xBoundary := len(m[0]) - 1
	if g.coordinates[0] > xBoundary {
		return true
	}

	yBoundary := len(m) - 1
	if g.coordinates[1] > yBoundary {
		return true
	}

	return false
}

func DaySixPartOneHandle(file []byte) int {
	paths := strings.Split(string(file), "\n")
	m := make([][]string, len(paths))
	for i := range paths {
		x := strings.Split(paths[i], "")
		m[i] = x
	}

	shouldKeepMoving := true
	g, err := newGuard(m)
	if err != nil {
		panic(err)
	}

	tracker := &GuardMovementTracker{}

	for shouldKeepMoving {
		c := g.coordinates
		shouldKeepMoving = g.Move(m)

		tracker.Track(c)
	}

	return tracker.Len()
}
