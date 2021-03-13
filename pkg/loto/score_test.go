package loto

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewScore(t *testing.T) {
	t.Run("Should be ok", func(t *testing.T) {
		d := Draw{
			B1:    21,
			B2:    42,
			B3:    10,
			B4:    11,
			B5:    31,
			Joker: 8,
		}
		balls := []int32{18, 23, 31, 42, 21}
		joker := int32(8)
		expectedScore := Score{
			Value: 130,
			Draw:  d,
			Balls: balls,
			Joker: joker,
		}

		sc := NewScore(d, balls, joker)
		assert.EqualValues(t, expectedScore, sc)
	})
}
