package loto

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReward_AddScore(t *testing.T) {
	t.Run("SHould be ok with score == 0", func(t *testing.T) {
		r := Reward{}
		expectedReward := Reward{n0: 1}

		r.AddScore(0)

		assert.EqualValues(t, expectedReward, r)
	})
	t.Run("SHould be ok with score == 100", func(t *testing.T) {
		r := Reward{}
		expectedReward := Reward{n0AndLucky: 1}

		r.AddScore(100)

		assert.EqualValues(t, expectedReward, r)
	})
	t.Run("SHould be ok with score == 10", func(t *testing.T) {
		r := Reward{}
		expectedReward := Reward{n1: 1}

		r.AddScore(10)

		assert.EqualValues(t, expectedReward, r)
	})
	t.Run("SHould be ok with score == 110", func(t *testing.T) {
		r := Reward{}
		expectedReward := Reward{n1AndLucky: 1}

		r.AddScore(110)

		assert.EqualValues(t, expectedReward, r)
	})
	t.Run("SHould be ok with score == 20", func(t *testing.T) {
		r := Reward{}
		expectedReward := Reward{n2: 1}

		r.AddScore(20)

		assert.EqualValues(t, expectedReward, r)
	})
	t.Run("SHould be ok with score == 120", func(t *testing.T) {
		r := Reward{}
		expectedReward := Reward{n2AndLucky: 1}

		r.AddScore(120)

		assert.EqualValues(t, expectedReward, r)
	})
	t.Run("SHould be ok with score == 30", func(t *testing.T) {
		r := Reward{}
		expectedReward := Reward{n3: 1}

		r.AddScore(30)

		assert.EqualValues(t, expectedReward, r)
	})
	t.Run("SHould be ok with score == 130", func(t *testing.T) {
		r := Reward{}
		expectedReward := Reward{n3AndLucky: 1}

		r.AddScore(130)

		assert.EqualValues(t, expectedReward, r)
	})
	t.Run("SHould be ok with score == 40", func(t *testing.T) {
		r := Reward{}
		expectedReward := Reward{n4: 1}

		r.AddScore(40)

		assert.EqualValues(t, expectedReward, r)
	})
	t.Run("SHould be ok with score == 140", func(t *testing.T) {
		r := Reward{}
		expectedReward := Reward{n4AndLucky: 1}

		r.AddScore(140)

		assert.EqualValues(t, expectedReward, r)
	})
	t.Run("SHould be ok with score == 50", func(t *testing.T) {
		r := Reward{}
		expectedReward := Reward{n5: 1}

		r.AddScore(50)

		assert.EqualValues(t, expectedReward, r)
	})
	t.Run("SHould be ok with score == 150", func(t *testing.T) {
		r := Reward{}
		expectedReward := Reward{n5AndLucky: 1}

		r.AddScore(150)

		assert.EqualValues(t, expectedReward, r)
	})
}

func TestReward_String(t *testing.T) {
	t.Run("Should be ok", func(t *testing.T) {
		r := Reward{
			n5AndLucky: 42,
			n5:         43,
			n4AndLucky: 44,
			n4:         45,
			n3AndLucky: 46,
			n3:         47,
			n2AndLucky: 48,
			n2:         49,
			n1AndLucky: 50,
			n1:         51,
			n0AndLucky: 52,
			n0:         53,
		}
		expectedString := fmt.Sprintf(
			"5Balls and lucky number: 42\n" +
				"5Balls without lucky number: 43\n" +
				"4Balls and lucky number: 44\n" +
				"4Balls without lucky number: 45\n" +
				"3Balls and lucky number: 46\n" +
				"3Balls without lucky number: 47\n" +
				"2Balls and lucky number: 48\n" +
				"2Balls without lucky number: 49\n" +
				"1Balls and lucky number: 50\n" +
				"1Balls without lucky number: 51\n" +
				"0Balls and lucky number: 52\n" +
				"0Balls without lucky number: 53\n")

		assert.EqualValues(t, expectedString, r.String())
	})
}

func TestReward_Price(t *testing.T) {
	t.Run("Should be ok", func(t *testing.T) {
		r := Reward{
			n5AndLucky: 1,
			n5:         1,
			n4AndLucky: 1,
			n4:         1,
			n3AndLucky: 1,
			n3:         1,
			n2AndLucky: 1,
			n2:         1,
			n1AndLucky: 1,
			n1:         1,
			n0AndLucky: 1,
			n0:         1,
		}
		expectedCost := 12 * 2.2
		expectedWin := 2162561.5

		cost, win := r.Price()
		assert.InDelta(t, expectedCost, cost, 0.0001)
		assert.InDelta(t, expectedWin, win, 0.0001)
	})
}
