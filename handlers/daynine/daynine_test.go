package dayeight

import (
	"testing"

	"github.com/mateusdeitos/advent-of-code-2024/embed"
	"github.com/stretchr/testify/assert"
)

func TestDayNinePartOneHandler(t *testing.T) {

	t.Run("test example input", func(t *testing.T) {
		input := embed.FileInputDayNineExample

		count := DayNinePartOneHandle(input)
		assert.Equal(t, 1928, count)
	})

	t.Run("test real input", func(t *testing.T) {
		input := embed.FileInputDayNine

		count := DayNinePartOneHandle(input)
		assert.Equal(t, 289, count)
	})
}
