package parttwo

import (
	"fmt"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

type GameSet struct {
	Blue  int
	Red   int
	Green int
}

func (gs *GameSet) Inc(color string, amount int) {
	switch color {
	case Blue:
		gs.Blue += amount
	case Red:
		gs.Red += amount
	case Green:
		gs.Green += amount
	}
}

type Game struct {
	ID    int
	Sets  []GameSet
	Power int
}

// enum of colors
const (
	Blue  = "blue"
	Red   = "red"
	Green = "green"
)

var colors = []string{Blue, Red, Green}

func DaytwoPartTwoHandle(file []byte) int {
	lines := strings.Split(string(file), "\n")
	sum := 0
	for i := range lines {
		g, err := parseLineToGame(lines[i])
		if err != nil {
			panic(err)
		}

		sum += g.Power
	}

	return sum
}

func parseLineToGame(line string) (*Game, error) {
	game := &Game{}

	id, err := extractGameID(line)
	if err != nil {
		return nil, err
	}

	game.ID = id

	sets, err := extractSets(line)
	if err != nil {
		return nil, err
	}
	game.Sets = sets

	var r, g, b int
	for _, s := range sets {
		if s.Red > r {
			r = s.Red
		}

		if s.Blue > b {
			b = s.Blue
		}

		if s.Green > g {
			g = s.Green
		}
	}

	game.Power = r * g * b

	return game, nil
}

func extractGameID(line string) (int, error) {
	parts := strings.Split(strings.Split(line, ":")[0], " ")
	return strconv.Atoi(parts[1])
}

var setPattern = regexp.MustCompile(`(\d*) (blue|red|green)`)

func extractSets(line string) ([]GameSet, error) {
	idx := strings.Index(line, ": ")
	line = strings.TrimSpace(line[idx+1:])
	setsStr := strings.Split(line, ";")
	sets := []GameSet{}
	for i, sStr := range setsStr {
		matches := setPattern.FindAllString(sStr, -1)
		if len(matches) == 0 {
			continue
		}

		set := GameSet{}
		for _, match := range matches {
			parts := strings.Split(match, " ")
			if len(parts) != 2 {
				return sets, fmt.Errorf("[%d] - unexpected set length: %s", i, match)
			}

			amount, err := strconv.Atoi(parts[0])
			if err != nil {
				return sets, err
			}

			color := parts[1]
			if !slices.Contains(colors, color) {
				return nil, fmt.Errorf("[%d] - invalid color: %s", i, color)
			}

			set.Inc(color, amount)
		}

		sets = append(sets, set)
	}

	return sets, nil
}
