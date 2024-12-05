package dayfour

import (
	"testing"

	"github.com/mateusdeitos/advent-of-code-2024/embed"
	"github.com/stretchr/testify/assert"
)

func TestDayFourPartOneHandler(t *testing.T) {

	t.Run("test horizontal search", func(t *testing.T) {
		input := []byte(`XMAS______
						 ___SAMX___
						 __XMASAMX_
						 ______XMAS
						 ______SAMX
						`)

		count := DayFourPartOneHandle(input)
		assert.Equal(t, 6, count)
	})

	t.Run("test vertical search", func(t *testing.T) {
		input := []byte(`X_________
						 M___X____S
						 A___M____A
						 S___M____M
						 ____X____X
						`)

		count := DayFourPartOneHandle(input)
		assert.Equal(t, 2, count)
	})

	t.Run("test diagonal search", func(t *testing.T) {
		input := []byte(`X___________________X________X
						 _M____S_____S___S____M____S_M_
						 __A____A___A_____A____A____A__
						 ___S____M_M_______M____S__S_M_
						 _________X_________X_________X
						`)

		count := DayFourPartOneHandle(input)
		assert.Equal(t, 7, count)
	})

	t.Run("test real input", func(t *testing.T) {
		input := embed.FileInputDayFour

		count := DayFourPartOneHandle(input)
		assert.Equal(t, 2454, count)
	})
}

func TestDayTwoPartTwoHandler(t *testing.T) {

}
