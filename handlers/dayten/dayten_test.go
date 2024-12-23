package dayten

import (
	"testing"

	"github.com/mateusdeitos/advent-of-code-2024/embed"
	"github.com/stretchr/testify/assert"
)

func TestDaytenPartOneHandler(t *testing.T) {

	t.Run("test example input", func(t *testing.T) {
		input := embed.FileInputDaytenExample

		count := DaytenPartOneHandle(input)
		assert.Equal(t, 36, count)
	})

	t.Run("test real input", func(t *testing.T) {
		input := embed.FileInputDayten

		count := DaytenPartOneHandle(input)
		assert.Equal(t, 733, count)
	})
}

func TestDaytenPartTwoHandler(t *testing.T) {

	t.Run("test real input", func(t *testing.T) {
		input := embed.FileInputDayten

		count := DaytenPartTwoHandle(input)
		assert.Equal(t, 1514, count)
	})
}
