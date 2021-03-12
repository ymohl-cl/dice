package loto

import "fmt"

// Prices values
var (
	PriceN5AndLucky float64 = 2000000.0
	PriceN5         float64 = 159957.0
	PriceN4AndLucky float64 = 2054.80
	PriceN4         float64 = 430.60
	PriceN3AndLucky float64 = 73.60
	PriceN3         float64 = 21.70
	PriceN2AndLucky float64 = 14.70
	PriceN2         float64 = 4.70
	PriceN1AndLucky float64 = 2.2
	PriceN1         float64 = 0.0
	PriceN0AndLucky float64 = 2.2
	PriceN0         float64 = 0.0
	CostTicket      float64 = 2.2
)

// Reward statistique
type Reward struct {
	n5AndLucky int64
	n5         int64
	n4AndLucky int64
	n4         int64
	n3AndLucky int64
	n3         int64
	n2AndLucky int64
	n2         int64
	n1AndLucky int64
	n1         int64
	n0AndLucky int64
	n0         int64
}

// AddScore on the reward database
func (r *Reward) AddScore(score int64) {
	jokerScore := int64(0)
	if score >= 100 {
		jokerScore = JokerValue
		score -= JokerValue
	}
	ballScore := score

	switch ballScore {
	case BallValue * 5:
		if jokerScore == JokerValue {
			r.n5AndLucky++
		} else {
			r.n5++
		}
	case BallValue * 4:
		if jokerScore == JokerValue {
			r.n4AndLucky++
		} else {
			r.n4++
		}
	case BallValue * 3:
		if jokerScore == JokerValue {
			r.n3AndLucky++
		} else {
			r.n3++
		}
	case BallValue * 2:
		if jokerScore == JokerValue {
			r.n2AndLucky++
		} else {
			r.n2++
		}
	case BallValue * 1:
		if jokerScore == JokerValue {
			r.n1AndLucky++
		} else {
			r.n1++
		}
	default:
		if jokerScore == JokerValue {
			r.n0AndLucky++
		} else {
			r.n0++
		}
	}
}

func (r Reward) String() string {
	return fmt.Sprintf(
		"5Balls and lucky number: %d\n"+
			"5Balls without lucky number: %d\n"+
			"4Balls and lucky number: %d\n"+
			"4Balls without lucky number: %d\n"+
			"3Balls and lucky number: %d\n"+
			"3Balls without lucky number: %d\n"+
			"2Balls and lucky number: %d\n"+
			"2Balls without lucky number: %d\n"+
			"1Balls and lucky number: %d\n"+
			"1Balls without lucky number: %d\n"+
			"0Balls and lucky number: %d\n"+
			"0Balls without lucky number: %d\n",
		r.n5AndLucky,
		r.n5,
		r.n4AndLucky,
		r.n4,
		r.n3AndLucky,
		r.n3,
		r.n2AndLucky,
		r.n2,
		r.n1AndLucky,
		r.n1,
		r.n0AndLucky,
		r.n0)
}

// Price return the cost and win value to participation
func (r Reward) Price() (float64, float64) {
	var cost float64
	var win float64

	win += float64(r.n5AndLucky) * PriceN5AndLucky
	cost += float64(r.n5AndLucky) + CostTicket
	win += float64(r.n5) * PriceN5
	cost += float64(r.n5) + CostTicket

	win += float64(r.n4AndLucky) * PriceN4AndLucky
	cost += float64(r.n4AndLucky) + CostTicket
	win += float64(r.n4) * PriceN4
	cost += float64(r.n4) + CostTicket

	win += float64(r.n3AndLucky) * PriceN3AndLucky
	cost += float64(r.n3AndLucky) + CostTicket
	win += float64(r.n3) * PriceN3
	cost += float64(r.n3) + CostTicket

	win += float64(r.n2AndLucky) * PriceN2AndLucky
	cost += float64(r.n2AndLucky) + CostTicket
	win += float64(r.n2) * PriceN2
	cost += float64(r.n2) + CostTicket

	win += float64(r.n1AndLucky) * PriceN1AndLucky
	cost += float64(r.n1AndLucky) + CostTicket
	win += float64(r.n1) * PriceN1
	cost += float64(r.n1) + CostTicket

	win += float64(r.n0AndLucky) * PriceN0AndLucky
	cost += float64(r.n0AndLucky) + CostTicket
	win += float64(r.n0) * PriceN0
	cost += float64(r.n0) + CostTicket

	return cost, win
}
