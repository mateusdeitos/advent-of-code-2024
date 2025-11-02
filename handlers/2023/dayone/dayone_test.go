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
		assert.Equal(t, 56465, count)
	})
}

func TestDayonePartTwoHandler(t *testing.T) {

	t.Run("test example input", func(t *testing.T) {
		input := []byte(`sevenine
	two1nine
	eightwothree
	abcone2threexyz
	xtwone3four
	4nineeightseven2
	zoneight234
	7pqrstsixteen`)

		count := DayonePartTwoHandle(input)
		assert.Equal(t, 360, count)
	})

	t.Run("text real input", func(t *testing.T) {
		input := embed.FileInputDayone

		count := DayonePartTwoHandle(input)
		assert.Equal(t, 5529, count)
	})

}
