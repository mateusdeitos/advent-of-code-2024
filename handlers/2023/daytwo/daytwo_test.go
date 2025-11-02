package daytwo

import (
	"testing"

	embed "github.com/mateusdeitos/advent-of-code-2024/embed/2023"
	"github.com/stretchr/testify/assert"
)

func TestDaytwoPartOneHandler(t *testing.T) {

	t.Run("test example input", func(t *testing.T) {
		input := embed.FileInputDaytwoExample

		count := DaytwoPartOneHandle(input)
		assert.Equal(t, 8, count)
	})

	t.Run("test real input", func(t *testing.T) {
		input := embed.FileInputDaytwo

		count := DaytwoPartOneHandle(input)
		assert.Equal(t, 289, count)
	})
}
