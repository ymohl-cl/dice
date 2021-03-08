package loto

import (
	"fmt"

	"github.com/ymohl-cl/dice/pkg/dice"
)

type Loto interface {
	Print()
}

type loto struct {
	d1    dice.Dice
	d2    dice.Dice
	d3    dice.Dice
	d4    dice.Dice
	d5    dice.Dice
	joker dice.Dice
}

// New loto instance with the loading of loto hitory
func New(historyFiles []string) (Loto, error) {
	var l loto
	var err error

	if l.d1, err = dice.New(49); err != nil {
		return nil, err
	}
	if l.d2, err = dice.New(49); err != nil {
		return nil, err
	}
	if l.d3, err = dice.New(49); err != nil {
		return nil, err
	}
	if l.d4, err = dice.New(49); err != nil {
		return nil, err
	}
	if l.d5, err = dice.New(49); err != nil {
		return nil, err
	}
	if l.joker, err = dice.New(10); err != nil {
		return nil, err
	}

	for _, f := range historyFiles {
		var history []Draw

		if history, err = NewHistory(f); err != nil {
			return nil, err
		}
		for _, draw := range history {
			l.parseDraw(draw)
		}
	}

	return &l, nil
}

func (l *loto) parseDraw(d Draw) {
	l.d1.SetThrow(d.B1)
	l.d2.SetThrow(d.B2)
	l.d3.SetThrow(d.B3)
	l.d4.SetThrow(d.B4)
	l.d5.SetThrow(d.B5)
	l.joker.SetThrow(d.Joker)
}

// Print info about the loto
func (l loto) Print() {
	fmt.Printf("D1 has number throw: %d and weaklestface %v\n", l.d1.NbThrow(), l.d1.WeaklestFaces())
	fmt.Printf("D2 has number throw: %d and weaklestface %v\n", l.d2.NbThrow(), l.d2.WeaklestFaces())
	fmt.Printf("D3 has number throw: %d and weaklestface %v\n", l.d3.NbThrow(), l.d3.WeaklestFaces())
	fmt.Printf("D4 has number throw: %d and weaklestface %v\n", l.d4.NbThrow(), l.d4.WeaklestFaces())
	fmt.Printf("D5 has number throw: %d and weaklestface %v\n", l.d5.NbThrow(), l.d5.WeaklestFaces())
	fmt.Printf("Djoker has number throw: %d and weaklestface %v\n", l.joker.NbThrow(), l.joker.WeaklestFaces())
}
