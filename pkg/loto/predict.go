package loto

import (
	"fmt"

	"github.com/ymohl-cl/dice/pkg/dice"
)

// GeneralPredict concat the history (all balls trows on the same ball exept to the joker)
type GeneralPredict struct {
	Balls dice.Dice
	Joker dice.Dice
}

// NewGeneralPredict instance
func NewGeneralPredict() (GeneralPredict, error) {
	var err error
	var g GeneralPredict

	if g.Balls, err = dice.New(49); err != nil {
		return GeneralPredict{}, err
	}
	if g.Joker, err = dice.New(10); err != nil {
		return GeneralPredict{}, err
	}
	return g, nil
}

// Print predict method
func (p GeneralPredict) Print() {
	fmt.Printf("General prediction:\n")
	p.Balls.SetThrow(int32(8))
	p.Balls.SetThrow(int32(19))
	p.Balls.SetThrow(int32(24))
	p.Balls.SetThrow(int32(38))
	p.Balls.SetThrow(int32(40))
	faces, results := p.Balls.Order()
	fmt.Printf("grid throw: \n\n")
	for i, v := range results {
		fmt.Printf("%d :=> %d\n", faces[i], v)
	}
	fmt.Printf("Joker throw: \n\n")
	p.Joker.SetThrow(int32(10))
	faces, results = p.Joker.Order()
	for i, v := range results {
		fmt.Printf("%d :=> %d\n", faces[i], v)
	}
}

// BallsPredict concat the history (ball by ball)
type BallsPredict struct {
	Ball1 dice.Dice
	Ball2 dice.Dice
	Ball3 dice.Dice
	Ball4 dice.Dice
	Ball5 dice.Dice
	Joker dice.Dice
}

// NewBallsPredict instance
func NewBallsPredict() (BallsPredict, error) {
	var err error
	var b BallsPredict

	if b.Ball1, err = dice.New(49); err != nil {
		return BallsPredict{}, err
	}
	if b.Ball2, err = dice.New(49); err != nil {
		return BallsPredict{}, err
	}
	if b.Ball3, err = dice.New(49); err != nil {
		return BallsPredict{}, err
	}
	if b.Ball4, err = dice.New(49); err != nil {
		return BallsPredict{}, err
	}
	if b.Ball5, err = dice.New(49); err != nil {
		return BallsPredict{}, err
	}
	if b.Joker, err = dice.New(10); err != nil {
		return BallsPredict{}, err
	}

	return b, nil
}

// Print predict method
func (p BallsPredict) Print() {
	fmt.Printf("Balls prediction:\n\n")
	fmt.Printf("ball_1 has number throw: %d and weaklestface %v\n", p.Ball1.NbThrow(), p.Ball1.WeaklestFaces())
	fmt.Printf("ball_2 has number throw: %d and weaklestface %v\n", p.Ball2.NbThrow(), p.Ball2.WeaklestFaces())
	fmt.Printf("ball_3 has number throw: %d and weaklestface %v\n", p.Ball3.NbThrow(), p.Ball3.WeaklestFaces())
	fmt.Printf("ball_4 has number throw: %d and weaklestface %v\n", p.Ball4.NbThrow(), p.Ball4.WeaklestFaces())
	fmt.Printf("ball_5 has number throw: %d and weaklestface %v\n", p.Ball5.NbThrow(), p.Ball5.WeaklestFaces())
	fmt.Printf("joker has number throw: %d and weaklestface %v\n", p.Joker.NbThrow(), p.Joker.WeaklestFaces())
}
