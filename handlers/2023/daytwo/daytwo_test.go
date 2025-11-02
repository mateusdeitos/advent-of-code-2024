package daytwo

import (
	"testing"

	embed "github.com/mateusdeitos/advent-of-code-2024/embed/2023"
	partone "github.com/mateusdeitos/advent-of-code-2024/handlers/2023/daytwo/part_one"
	parttwo "github.com/mateusdeitos/advent-of-code-2024/handlers/2023/daytwo/part_two"
	"github.com/stretchr/testify/assert"
)

func TestDaytwoPartOneHandler(t *testing.T) {

	t.Run("test example input", func(t *testing.T) {
		input := embed.FileInputDaytwoExample

		count := partone.DaytwoPartOneHandle(input)
		assert.Equal(t, 8, count)
	})

	t.Run("test real input", func(t *testing.T) {
		input := embed.FileInputDaytwo

		count := partone.DaytwoPartOneHandle(input)
		assert.Equal(t, 289, count)
	})
}

func TestDaytwoPartTwoHandler(t *testing.T) {

	t.Run("test example input", func(t *testing.T) {
		input := embed.FileInputDaytwoExample

		count := parttwo.DaytwoPartTwoHandle(input)
		assert.Equal(t, 2_286, count)
	})

	t.Run("test real input", func(t *testing.T) {
		input := embed.FileInputDaytwo

		count := parttwo.DaytwoPartTwoHandle(input)
		assert.Equal(t, 70_768, count)
	})

}
