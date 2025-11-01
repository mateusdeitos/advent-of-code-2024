package dayone

import (
	"testing"

	embed "github.com/mateusdeitos/advent-of-code-2024/embed/2023"
	"github.com/stretchr/testify/assert"
)

func TestDayonePartOneHandler(t *testing.T) {

	t.Run("test example input", func(t *testing.T) {
		input := embed.FileInputDayoneExample

		count := DayonePartOneHandle(input)
		assert.Equal(t, 142, count)
	})

	t.Run("test real input", func(t *testing.T) {
		input := embed.FileInputDayone

		count := DayonePartOneHandle(input)
		assert.Equal(t, 289, count)
	})
}
