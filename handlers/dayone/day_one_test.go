package dayone

import (
	"testing"

	"github.com/mateusdeitos/advent-of-code-2024/embed"
	"github.com/stretchr/testify/assert"
)

func TestDayOnePartOneHandler(t *testing.T) {

	t.Run("test real input", func(t *testing.T) {
		input := embed.FileInputDayOne

		totalDistance, err := DayOnePartOneHandle(input)
		assert.NoError(t, err)

		assert.Equal(t, 1151792, totalDistance)
	})

	t.Run("should return total distance", func(t *testing.T) {
		input := []byte(`
			59569,19816
			85735,35587
			30874,14008
			15369,52468
			25998,79528
			37079,53944
		`)

		totalDistance, err := DayOnePartOneHandle(input)
		assert.NoError(t, err)

		assert.Equal(t, 39477, totalDistance)
	})

	t.Run("should return total distance = 0", func(t *testing.T) {
		input := []byte(`
			3,3
			1,1
			5,5	
		`)

		totalDistance, err := DayOnePartOneHandle(input)
		assert.NoError(t, err)

		assert.Equal(t, 0, totalDistance)
	})

	t.Run("should handle negative numbers", func(t *testing.T) {
		input := []byte(`
			1,-1
		`)

		totalDistance, err := DayOnePartOneHandle(input)
		assert.NoError(t, err)

		assert.Equal(t, 2, totalDistance)
	})

	t.Run("should return 0 on empty file", func(t *testing.T) {
		input := []byte(`
		`)

		totalDistance, err := DayOnePartOneHandle(input)
		assert.NoError(t, err)

		assert.Equal(t, 0, totalDistance)
	})

}

func TestDayOnePartTwoHandler(t *testing.T) {

	t.Run("test real input", func(t *testing.T) {
		input := embed.FileInputDayOne

		score, err := DayOnePartTwoHandle(input)
		assert.NoError(t, err)

		assert.Equal(t, 21790168, score)
	})
}
