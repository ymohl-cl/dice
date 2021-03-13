package loto

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDefaultLoto_Prediction(t *testing.T) {
	t.Run("Should return an error because index is out of range", func(t *testing.T) {
		l := &defaultLoto{}
		opt := Option{Index: 1}
		expectedRest := int64(0)
		expectedErr := "invalid index to start exploration draw"
		expectedBallsNBThrow := int64(0)
		expectedJokerNBThrow := int64(0)

		rest, err := l.Prediction(opt)
		if assert.Error(t, err, expectedErr) {
			assert.EqualValues(t, expectedRest, rest)
			assert.EqualValues(t, expectedBallsNBThrow, l.balls.NbThrow())
			assert.EqualValues(t, expectedJokerNBThrow, l.joker.NbThrow())
		}
	})
	t.Run("Should be ok without history", func(t *testing.T) {
		l := &defaultLoto{}
		opt := Option{Index: 0}
		expectedRest := int64(0)
		expectedBallsNBThrow := int64(0)
		expectedJokerNBThrow := int64(0)

		rest, err := l.Prediction(opt)
		if assert.NoError(t, err) {
			assert.EqualValues(t, expectedRest, rest)
			assert.EqualValues(t, expectedBallsNBThrow, l.balls.NbThrow())
			assert.EqualValues(t, expectedJokerNBThrow, l.joker.NbThrow())
		}
	})
	t.Run("Should be ok with one draw history and non option set", func(t *testing.T) {
		l := &defaultLoto{history: History{RecentDraws: []Draw{{ID: "hello"}}}}
		opt := Option{Index: 0}
		expectedRest := int64(0)
		expectedBallsNBThrow := int64(5)
		expectedJokerNBThrow := int64(1)

		rest, err := l.Prediction(opt)
		if assert.NoError(t, err) {
			assert.EqualValues(t, expectedRest, rest)
			assert.EqualValues(t, expectedBallsNBThrow, l.balls.NbThrow())
			assert.EqualValues(t, expectedJokerNBThrow, l.joker.NbThrow())
		}
	})
	t.Run("Should be ok with a portion draw history", func(t *testing.T) {
		l := &defaultLoto{history: History{RecentDraws: []Draw{
			{ID: "1"}, {ID: "2"}, {ID: "3"}, {ID: "4"}}}}
		opt := Option{Index: 1, NDraws: 2}
		expectedRest := int64(1)
		expectedBallsNBThrow := int64(10)
		expectedJokerNBThrow := int64(2)

		rest, err := l.Prediction(opt)
		if assert.NoError(t, err) {
			assert.EqualValues(t, expectedRest, rest)
			assert.EqualValues(t, expectedBallsNBThrow, l.balls.NbThrow())
			assert.EqualValues(t, expectedJokerNBThrow, l.joker.NbThrow())
		}
	})
	t.Run("Should be ok with a the rest and option Ndraw out of range", func(t *testing.T) {
		l := &defaultLoto{history: History{RecentDraws: []Draw{
			{ID: "1"}, {ID: "2"}, {ID: "3"}, {ID: "4"}}}}
		opt := Option{Index: 0, NDraws: 10}
		expectedRest := int64(0)
		expectedBallsNBThrow := int64(20)
		expectedJokerNBThrow := int64(4)

		rest, err := l.Prediction(opt)
		if assert.NoError(t, err) {
			assert.EqualValues(t, expectedRest, rest)
			assert.EqualValues(t, expectedBallsNBThrow, l.balls.NbThrow())
			assert.EqualValues(t, expectedJokerNBThrow, l.joker.NbThrow())
		}
	})

	t.Run("Should be ok with the old history", func(t *testing.T) {
		l := &defaultLoto{history: History{OldDraws: []OldDraw{
			{ID: "1"}, {ID: "2"}, {ID: "3"}, {ID: "4"}}}}
		opt := Option{Index: 0, NDraws: 10, Old: true}
		expectedRest := int64(0)
		expectedBallsNBThrow := int64(20)
		expectedJokerNBThrow := int64(4)

		rest, err := l.Prediction(opt)
		if assert.NoError(t, err) {
			assert.EqualValues(t, expectedRest, rest)
			assert.EqualValues(t, expectedBallsNBThrow, l.balls.NbThrow())
			assert.EqualValues(t, expectedJokerNBThrow, l.joker.NbThrow())
		}
	})
	t.Run("Should be ok with the old history and the sisteeth ball", func(t *testing.T) {
		l := &defaultLoto{history: History{OldDraws: []OldDraw{
			{ID: "1"}, {ID: "2"}, {ID: "3"}, {ID: "4"}}}}
		opt := Option{Index: 0, NDraws: 10, Old: true, Old6thBall: true}
		expectedRest := int64(0)
		expectedBallsNBThrow := int64(24)
		expectedJokerNBThrow := int64(4)

		rest, err := l.Prediction(opt)
		if assert.NoError(t, err) {
			assert.EqualValues(t, expectedRest, rest)
			assert.EqualValues(t, expectedBallsNBThrow, l.balls.NbThrow())
			assert.EqualValues(t, expectedJokerNBThrow, l.joker.NbThrow())
		}
	})
	t.Run("Should be ok with the old history and the lucky ball", func(t *testing.T) {
		l := &defaultLoto{history: History{OldDraws: []OldDraw{
			{ID: "1"}, {ID: "2"}, {ID: "3"}, {ID: "4"}}}}
		opt := Option{Index: 0, NDraws: 10, Old: true, OldLuckyBall: true}
		expectedRest := int64(0)
		expectedBallsNBThrow := int64(24)
		expectedJokerNBThrow := int64(4)

		rest, err := l.Prediction(opt)
		if assert.NoError(t, err) {
			assert.EqualValues(t, expectedRest, rest)
			assert.EqualValues(t, expectedBallsNBThrow, l.balls.NbThrow())
			assert.EqualValues(t, expectedJokerNBThrow, l.joker.NbThrow())
		}
	})
}
func TestDefaultLoto_History(t *testing.T) {
	t.Run("Should be ok", func(t *testing.T) {
		l := &defaultLoto{history: History{RecentDraws: []Draw{
			{ID: "1"}, {ID: "2"}, {ID: "3"}, {ID: "4"}}}}
		expectedHist := History{RecentDraws: []Draw{
			{ID: "1"}, {ID: "2"}, {ID: "3"}, {ID: "4"}}}

		hist := l.History()
		assert.EqualValues(t, expectedHist, hist)
	})
}
func TestDefaultLoto_String(t *testing.T) {
	t.Run("Should be ok", func(t *testing.T) {
		l := &defaultLoto{}
		l.resetPrediction()
		l.balls.SetThrow(42)
		l.balls.SetThrow(42)
		l.balls.SetThrow(7)
		l.balls.SetThrow(2)
		l.joker.SetThrow(18)
		// add unit test on dice before fix this test. because dice.Order should not return the non called faces
		expectedSTR := fmt.Sprintf("the default order prediction in follow about the picks balls:\n\tface number 2 pick 1 time(s)\n\tface number 7 pick 1 time(s)\n\tface number 42 pick 2 time(s)\n")
		expectedSTR += fmt.Sprintf("the default order prediction in follow about the pick joker:\n\tface number 18 pick 1 time(s)\n")

		str := l.String()
		assert.EqualValues(t, expectedSTR, str)
	})
}

func TestVariant1Loto_Prediction(t *testing.T) {
	t.Run("Should return an error because index is out of range", func(t *testing.T) {
		l := &variant1Loto{}
		opt := Option{Index: 1}
		expectedRest := int64(0)
		expectedErr := "invalid index to start exploration draw"
		expectedBallsNBThrow := int64(0)
		expectedJokerNBThrow := int64(0)

		rest, err := l.Prediction(opt)
		if assert.Error(t, err, expectedErr) {
			assert.EqualValues(t, expectedRest, rest)
			assert.EqualValues(t, expectedBallsNBThrow, l.ball1.NbThrow())
			assert.EqualValues(t, expectedBallsNBThrow, l.ball2.NbThrow())
			assert.EqualValues(t, expectedBallsNBThrow, l.ball3.NbThrow())
			assert.EqualValues(t, expectedBallsNBThrow, l.ball4.NbThrow())
			assert.EqualValues(t, expectedBallsNBThrow, l.ball5.NbThrow())
			assert.EqualValues(t, expectedJokerNBThrow, l.joker.NbThrow())
		}
	})
	t.Run("Should be ok without history", func(t *testing.T) {
		l := &variant1Loto{}
		opt := Option{Index: 0}
		expectedRest := int64(0)
		expectedBallsNBThrow := int64(0)
		expectedJokerNBThrow := int64(0)

		rest, err := l.Prediction(opt)
		if assert.NoError(t, err) {
			assert.EqualValues(t, expectedRest, rest)
			assert.EqualValues(t, expectedBallsNBThrow, l.ball1.NbThrow())
			assert.EqualValues(t, expectedBallsNBThrow, l.ball2.NbThrow())
			assert.EqualValues(t, expectedBallsNBThrow, l.ball3.NbThrow())
			assert.EqualValues(t, expectedBallsNBThrow, l.ball4.NbThrow())
			assert.EqualValues(t, expectedBallsNBThrow, l.ball5.NbThrow())
			assert.EqualValues(t, expectedJokerNBThrow, l.joker.NbThrow())
		}
	})
	t.Run("Should be ok with one draw history and non option set", func(t *testing.T) {
		l := &variant1Loto{history: History{RecentDraws: []Draw{{ID: "hello"}}}}
		opt := Option{Index: 0}
		expectedRest := int64(0)
		expectedBallsNBThrow := int64(1)
		expectedJokerNBThrow := int64(1)

		rest, err := l.Prediction(opt)
		if assert.NoError(t, err) {
			assert.EqualValues(t, expectedRest, rest)
			assert.EqualValues(t, expectedBallsNBThrow, l.ball1.NbThrow())
			assert.EqualValues(t, expectedBallsNBThrow, l.ball2.NbThrow())
			assert.EqualValues(t, expectedBallsNBThrow, l.ball3.NbThrow())
			assert.EqualValues(t, expectedBallsNBThrow, l.ball4.NbThrow())
			assert.EqualValues(t, expectedBallsNBThrow, l.ball5.NbThrow())
			assert.EqualValues(t, expectedJokerNBThrow, l.joker.NbThrow())
		}
	})
	t.Run("Should be ok with a portion draw history", func(t *testing.T) {
		l := &variant1Loto{history: History{RecentDraws: []Draw{
			{ID: "1"}, {ID: "2"}, {ID: "3"}, {ID: "4"}}}}
		opt := Option{Index: 1, NDraws: 2}
		expectedRest := int64(1)
		expectedBallsNBThrow := int64(2)
		expectedJokerNBThrow := int64(2)

		rest, err := l.Prediction(opt)
		if assert.NoError(t, err) {
			assert.EqualValues(t, expectedRest, rest)
			assert.EqualValues(t, expectedBallsNBThrow, l.ball1.NbThrow())
			assert.EqualValues(t, expectedBallsNBThrow, l.ball2.NbThrow())
			assert.EqualValues(t, expectedBallsNBThrow, l.ball3.NbThrow())
			assert.EqualValues(t, expectedBallsNBThrow, l.ball4.NbThrow())
			assert.EqualValues(t, expectedBallsNBThrow, l.ball5.NbThrow())
			assert.EqualValues(t, expectedJokerNBThrow, l.joker.NbThrow())
		}
	})
	t.Run("Should be ok with a the rest and option Ndraw out of range", func(t *testing.T) {
		l := &variant1Loto{history: History{RecentDraws: []Draw{
			{ID: "1"}, {ID: "2"}, {ID: "3"}, {ID: "4"}}}}
		opt := Option{Index: 0, NDraws: 10}
		expectedRest := int64(0)
		expectedBallsNBThrow := int64(4)
		expectedJokerNBThrow := int64(4)

		rest, err := l.Prediction(opt)
		if assert.NoError(t, err) {
			assert.EqualValues(t, expectedRest, rest)
			assert.EqualValues(t, expectedBallsNBThrow, l.ball1.NbThrow())
			assert.EqualValues(t, expectedBallsNBThrow, l.ball2.NbThrow())
			assert.EqualValues(t, expectedBallsNBThrow, l.ball3.NbThrow())
			assert.EqualValues(t, expectedBallsNBThrow, l.ball4.NbThrow())
			assert.EqualValues(t, expectedBallsNBThrow, l.ball5.NbThrow())
			assert.EqualValues(t, expectedJokerNBThrow, l.joker.NbThrow())
		}
	})

	t.Run("Should be ok with the old history", func(t *testing.T) {
		l := &variant1Loto{history: History{OldDraws: []OldDraw{
			{ID: "1"}, {ID: "2"}, {ID: "3"}, {ID: "4"}}}}
		opt := Option{Index: 0, NDraws: 10, Old: true}
		expectedRest := int64(0)
		expectedBallsNBThrow := int64(4)
		expectedJokerNBThrow := int64(4)

		rest, err := l.Prediction(opt)
		if assert.NoError(t, err) {
			assert.EqualValues(t, expectedRest, rest)
			assert.EqualValues(t, expectedBallsNBThrow, l.ball1.NbThrow())
			assert.EqualValues(t, expectedBallsNBThrow, l.ball2.NbThrow())
			assert.EqualValues(t, expectedBallsNBThrow, l.ball3.NbThrow())
			assert.EqualValues(t, expectedBallsNBThrow, l.ball4.NbThrow())
			assert.EqualValues(t, expectedBallsNBThrow, l.ball5.NbThrow())
			assert.EqualValues(t, expectedJokerNBThrow, l.joker.NbThrow())
		}
	})
	t.Run("Should be ok with the old history and the sisteeth ball", func(t *testing.T) {
		l := &variant1Loto{history: History{OldDraws: []OldDraw{
			{ID: "1"}, {ID: "2"}, {ID: "3"}, {ID: "4"}}}}
		opt := Option{Index: 0, NDraws: 10, Old: true, Old6thBall: true}
		expectedRest := int64(0)
		expectedBallsNBThrow := int64(8)
		expectedJokerNBThrow := int64(4)

		rest, err := l.Prediction(opt)
		if assert.NoError(t, err) {
			assert.EqualValues(t, expectedRest, rest)
			assert.EqualValues(t, expectedBallsNBThrow, l.ball1.NbThrow())
			assert.EqualValues(t, expectedBallsNBThrow, l.ball2.NbThrow())
			assert.EqualValues(t, expectedBallsNBThrow, l.ball3.NbThrow())
			assert.EqualValues(t, expectedBallsNBThrow, l.ball4.NbThrow())
			assert.EqualValues(t, expectedBallsNBThrow, l.ball5.NbThrow())
			assert.EqualValues(t, expectedJokerNBThrow, l.joker.NbThrow())
		}
	})
	t.Run("Should be ok with the old history and the lucky ball", func(t *testing.T) {
		l := &variant1Loto{history: History{OldDraws: []OldDraw{
			{ID: "1"}, {ID: "2"}, {ID: "3"}, {ID: "4"}}}}
		opt := Option{Index: 0, NDraws: 10, Old: true, OldLuckyBall: true}
		expectedRest := int64(0)
		expectedBallsNBThrow := int64(8)
		expectedJokerNBThrow := int64(4)

		rest, err := l.Prediction(opt)
		if assert.NoError(t, err) {
			assert.EqualValues(t, expectedRest, rest)
			assert.EqualValues(t, expectedBallsNBThrow, l.ball1.NbThrow())
			assert.EqualValues(t, expectedBallsNBThrow, l.ball2.NbThrow())
			assert.EqualValues(t, expectedBallsNBThrow, l.ball3.NbThrow())
			assert.EqualValues(t, expectedBallsNBThrow, l.ball4.NbThrow())
			assert.EqualValues(t, expectedBallsNBThrow, l.ball5.NbThrow())
			assert.EqualValues(t, expectedJokerNBThrow, l.joker.NbThrow())
		}
	})
}
func TestVariant1Loto_History(t *testing.T) {
	t.Run("Should be ok", func(t *testing.T) {
		l := &variant1Loto{history: History{RecentDraws: []Draw{
			{ID: "1"}, {ID: "2"}, {ID: "3"}, {ID: "4"}}}}
		expectedHist := History{RecentDraws: []Draw{
			{ID: "1"}, {ID: "2"}, {ID: "3"}, {ID: "4"}}}

		hist := l.History()
		assert.EqualValues(t, expectedHist, hist)
	})
}

func TestVariant1Loto_String(t *testing.T) {
	t.Run("Should be ok", func(t *testing.T) {

	})
}
