package main

import (
	_ "embed"
	"fmt"
	"os"

	"github.com/mateusdeitos/advent-of-code-2024/handlers/dayone"
)

//go:embed 1-day/input.txt
var file []byte

func main() {
	day := os.Args[1]
	part := os.Args[2]

	if day == "1" {
		if part == "1" {
			totalDistance, err := dayone.DayOnePartOneHandle(file)
			if err != nil {
				panic(err)
			}

			fmt.Printf("Total distance: %d\n", totalDistance)
			return
		}

		if part == "2" {
			score, err := dayone.DayOnePartTwoHandle(file)
			if err != nil {
				panic(err)
			}

			fmt.Printf("Similarity Score: %d\n", score)
			return
		}
	}

	fmt.Printf("Day %s part %s not implemented\n", day, part)
}
