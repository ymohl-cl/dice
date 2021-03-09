package main

import (
	"github.com/ymohl-cl/dice/pkg/loto"
)

func main() {
	var l loto.Loto
	var err error
	var gPredict loto.GeneralPredict
	var bPredict loto.BallsPredict

	if l, err = loto.New(); err != nil {
		panic(err)
	}

	if gPredict, err = l.GeneralPrediction(loto.Option{}); err != nil {
		panic(err)
	}
	if bPredict, err = l.BallsPrediction(loto.Option{}); err != nil {
		panic(err)
	}
	gPredict.Print()
	bPredict.Print()
	return
}
