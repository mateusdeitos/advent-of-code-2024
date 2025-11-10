package daythree

import (
	"testing"

	embed "github.com/mateusdeitos/advent-of-code-2024/embed/2023"
	partone "github.com/mateusdeitos/advent-of-code-2024/handlers/2023/daythree/part_one"
	parttwo "github.com/mateusdeitos/advent-of-code-2024/handlers/2023/daythree/part_two"
	"github.com/stretchr/testify/assert"
)

func TestDaythreePartOneHandler(t *testing.T) {

	t.Run("test example input", func(t *testing.T) {
		t.SkipNow()
		input := embed.FileInputDaythreeExample

		count := partone.DaythreePartOneHandle(input)
		assert.Equal(t, 4361, count)
	})

	t.Run("test corner case bottom left + bottom right", func(t *testing.T) {
		t.SkipNow()
		input := []byte("..$..\n12.12")

		count := partone.DaythreePartOneHandle(input)
		assert.Equal(t, 24, count)
	})

	t.Run("test corner case top left + top right", func(t *testing.T) {
		t.SkipNow()
		input := []byte("12.12\n..$..")

		count := partone.DaythreePartOneHandle(input)
		assert.Equal(t, 24, count)
	})

	t.Run("test real input", func(t *testing.T) {
		input := embed.FileInputDaythree

		count := partone.DaythreePartOneHandle(input)
		assert.Equal(t, 512_794, count)
	})

}

func TestDaythreePartTwoHandler(t *testing.T) {

	t.Run("test example input", func(t *testing.T) {
		input := embed.FileInputDaythreeExample

		count := parttwo.DaythreePartTwoHandle(input)
		assert.Equal(t, 467_835, count)
	})

	t.Run("test real input", func(t *testing.T) {
		input := embed.FileInputDaythree

		count := parttwo.DaythreePartTwoHandle(input)
		assert.Equal(t, 467_835, count)
	})
}
