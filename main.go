package main

import (
	"fmt"
	"os"

	"github.com/mateusdeitos/advent-of-code-2024/embed"
	"github.com/mateusdeitos/advent-of-code-2024/handlers/dayone"
)

func main() {
	day := os.Args[1]
	part := os.Args[2]
	file := embed.FileInputDayOne

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
