package daythree

import (
	"testing"

	"github.com/mateusdeitos/advent-of-code-2024/embed"
	"github.com/stretchr/testify/assert"
)

func TestDayTwoPartOneHandler(t *testing.T) {

	t.Run("test example input", func(t *testing.T) {
		input := []byte(`xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))`)

		count, err := DayThreePartOneHandle(input)
		assert.NoError(t, err)
		assert.Equal(t, 161, count)
	})

	t.Run("test real input", func(t *testing.T) {
		input := embed.FileInputDayThree

		count, err := DayThreePartOneHandle(input)
		assert.NoError(t, err)
		assert.Equal(t, 166905464, count)
	})
}

func TestDayTwoPartTwoHandler(t *testing.T) {

	t.Run("test example input", func(t *testing.T) {
		input := []byte(`xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))`)

		count, err := DayThreePartTwoHandle(input)
		assert.NoError(t, err)
		assert.Equal(t, 48, count)
	})

	t.Run("test real input", func(t *testing.T) {
		input := embed.FileInputDayThree

		count, err := DayThreePartTwoHandle(input)
		assert.NoError(t, err)
		assert.Equal(t, 72948684, count)
	})
}
