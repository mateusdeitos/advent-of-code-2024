package dayone

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDayOnePartOneHandler(t *testing.T) {

	t.Run("should return total distance", func(t *testing.T) {
		input := []byte(`
			1,2
			3,4
			5,6	
		`)

		totalDistance, err := DayOnePartOneHandle(input)
		assert.NoError(t, err)

		assert.Equal(t, 3, totalDistance)
	})

	t.Run("should return total distance = 0", func(t *testing.T) {
		input := []byte(`
			1,1
			3,3
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
		input := embed.FileInput

		score, err := DayOnePartTwoHandle(input)
		assert.NoError(t, err)

		assert.Equal(t, 21790168, score)
	})
}
