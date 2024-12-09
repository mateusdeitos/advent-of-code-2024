package dayeight

import (
	"testing"

	"github.com/mateusdeitos/advent-of-code-2024/embed"
	"github.com/stretchr/testify/assert"
)

func TestDayEightPartOneHandler(t *testing.T) {

	t.Run("test example input", func(t *testing.T) {
		input := embed.FileInputDayEightExample

		count := DayEightPartOneHandle(input)
		assert.Equal(t, 14, count)
	})

	t.Run("test real input", func(t *testing.T) {
		input := embed.FileInputDayEight

		count := DayEightPartOneHandle(input)
		assert.Equal(t, 289, count)
	})
}
