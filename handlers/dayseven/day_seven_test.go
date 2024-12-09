package dayseven

import (
	"testing"

	"github.com/mateusdeitos/advent-of-code-2024/embed"
	"github.com/stretchr/testify/assert"
)

func TestDaySevenPartOneHandler(t *testing.T) {

	t.Run("test example input", func(t *testing.T) {
		input := embed.FileInputDaySevenExample

		count := DaySevenPartOneHandle(input)
		assert.Equal(t, int64(3749), count)
	})

	t.Run("test real input", func(t *testing.T) {
		input := embed.FileInputDaySeven

		count := DaySevenPartOneHandle(input)
		assert.Equal(t, int64(850435817339), count)
	})
}
