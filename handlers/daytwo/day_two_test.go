package daytwo

import (
	"testing"

	"github.com/mateusdeitos/advent-of-code-2024/embed"
	"github.com/stretchr/testify/assert"
)

func TestDayTwoPartOneHandler(t *testing.T) {

	t.Run("test example input", func(t *testing.T) {
		input := []byte(`
			7 6 4 2 1
			1 2 7 8 9
			9 7 6 2 1
			1 3 2 4 5
			8 6 4 4 1
			1 3 6 7 9
		`)

		count, err := DayTwoPartOneHandle(input)
		assert.NoError(t, err)
		assert.Equal(t, 2, count)
	})

	t.Run("test real input", func(t *testing.T) {
		input := embed.FileInputDayTwo

		count, err := DayTwoPartOneHandle(input)
		assert.NoError(t, err)
		assert.Equal(t, 631, count)
	})
}

func TestDayTwoPartTwoHandler(t *testing.T) {

	t.Run("test example input", func(t *testing.T) {
		input := []byte(`
			48 45 47 50 53 56 57
			7 6 4 2 1
			1 2 7 8 9
			9 7 6 2 1
			1 3 2 4 5
			8 6 4 4 1
			1 3 6 7 9
		`)

		count, err := DayTwoPartTwoHandle(input)
		assert.NoError(t, err)
		assert.Equal(t, 5, count)
	})

	t.Run("test real input", func(t *testing.T) {
		input := embed.FileInputDayTwo

		count, err := DayTwoPartTwoHandle(input)
		assert.NoError(t, err)
		assert.Equal(t, 665, count)
	})
}
