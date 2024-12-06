package daysix

import (
	"testing"

	"github.com/mateusdeitos/advent-of-code-2024/embed"
	"github.com/stretchr/testify/assert"
)

func TestDaySixPartOneHandler(t *testing.T) {

	t.Run("test example input", func(t *testing.T) {
		input := embed.FileInputDaySixExample

		count := DaySixPartOneHandle(input)
		assert.Equal(t, 41, count)
	})

	t.Run("test real input", func(t *testing.T) {
		input := embed.FileInputDaySix

		count := DaySixPartOneHandle(input)
		assert.Equal(t, 4973, count)
	})
}

// func TestDaySixPartTwoHandler(t *testing.T) {

// 	t.Run("test example input", func(t *testing.T) {
// 		input := embed.FileInputDaySixExample

// 		count := DayFivePartTwoHandle(input)
// 		assert.Equal(t, 123, count)
// 	})

// 	t.Run("test real input", func(t *testing.T) {
// 		input := embed.FileInputDaySix

// 		count := DayFivePartTwoHandle(input)
// 		assert.Equal(t, 6085, count)
// 	})
// }
