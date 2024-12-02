package main

import (
	"fmt"
	"os"

	"github.com/mateusdeitos/advent-of-code-2024/embed"
	"github.com/mateusdeitos/advent-of-code-2024/handlers/dayone"
	"github.com/mateusdeitos/advent-of-code-2024/handlers/daytwo"
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

	if day == "2" {
		file = embed.FileInputDayTwo
		if part == "1" {
			count, err := daytwo.DayTwoPartOneHandle(file)
			if err != nil {
				panic(err)
			}

			fmt.Printf("Amount of safe reports: %d\n", count)
			return
		}
	}

	fmt.Printf("Day %s part %s not implemented\n", day, part)
}
