package daytwo

import (
	"testing"

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
}
