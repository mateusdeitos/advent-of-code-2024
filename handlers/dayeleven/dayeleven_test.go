package dayeleven

import (
	"testing"

	"github.com/mateusdeitos/advent-of-code-2024/embed"
	"github.com/stretchr/testify/assert"
)

func TestDayelevenPartOneHandler(t *testing.T) {

	t.Run("test example input", func(t *testing.T) {
		input := embed.FileInputDayelevenExample

		count := DayelevenPartOneHandle(input)
		assert.Equal(t, 55312, count)
	})

	t.Run("test real input", func(t *testing.T) {
		input := embed.FileInputDayeleven

		count := DayelevenPartOneHandle(input)
		assert.Equal(t, 213625, count)
	})
}
