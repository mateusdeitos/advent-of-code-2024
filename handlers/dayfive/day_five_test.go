package dayfive

import (
	"testing"

	"github.com/mateusdeitos/advent-of-code-2024/embed"
	"github.com/stretchr/testify/assert"
)

func TestDayFivePartOneHandler(t *testing.T) {

	t.Run("test example input", func(t *testing.T) {
		input := embed.FileInputDayFiveExample

		count := DayFivePartOneHandle(input)
		assert.Equal(t, 143, count)
	})

	t.Run("test real input", func(t *testing.T) {
		input := embed.FileInputDayFive

		count := DayFivePartOneHandle(input)
		assert.Equal(t, 4924, count)
	})
}

func TestDayFivePartTwoHandler(t *testing.T) {

	t.Run("test example input", func(t *testing.T) {
		input := embed.FileInputDayFiveExample

		count := DayFivePartTwoHandle(input)
		assert.Equal(t, 123, count)
	})

	t.Run("test real input", func(t *testing.T) {
		input := embed.FileInputDayFive

		count := DayFivePartTwoHandle(input)
		assert.Equal(t, 6085, count)
	})
}
