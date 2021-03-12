package main

import (
	"fmt"
	"math/rand"

	"github.com/ymohl-cl/dice/pkg/loto"
)

func main() {
	var l loto.Loto
	var err error
	var r loto.Reward

	if l, err = loto.New(); err != nil {
		panic(err)
	}

	func() {
		if r, err = playGeneralPrediction(l); err != nil {
			panic(err)
		}
		PrintAnalyse(r, "general prediction with get more probabilities draw after each trow")
	}()

	func() {
		balls := []int32{25, 7, 47, 23, 14}
		joker := int32(5)
		if r, err = playStaticSolution(l, balls, joker); err != nil {
			panic(err)
		}
		PrintAnalyse(r, fmt.Sprintf("static solution with following balls %v and joker number %d for each trows", balls, joker))
	}()

	func() {
		balls := []int32{48, 15, 37, 3, 5}
		joker := int32(4)
		if r, err = playStaticSolution(l, balls, joker); err != nil {
			panic(err)
		}
		PrintAnalyse(r, fmt.Sprintf("static solution with following balls %v and joker number %d for each trows", balls, joker))
	}()

	func() {
		balls := []int32{20, 19, 37, 26, 41}
		joker := int32(2)
		if r, err = playStaticSolution(l, balls, joker); err != nil {
			panic(err)
		}
		PrintAnalyse(r, fmt.Sprintf("static solution with following balls %v and joker number %d for each trows", balls, joker))
	}()

	func() {
		seed := int64(42)
		nBalls := int64(15)
		nJoker := int64(4)
		if r, err = playGeneralPredictionWithRandomSelection(l, nBalls, nJoker, seed); err != nil {
			panic(err)
		}
		PrintAnalyse(r, fmt.Sprintf("general prediction with get more %d balls proba and %d joker proba to random selection (seed: %d)", nBalls, nJoker, seed))
	}()

	func() {
		seed := int64(1680339887)
		nBalls := int64(15)
		nJoker := int64(4)
		if r, err = playGeneralPredictionWithRandomSelection(l, nBalls, nJoker, seed); err != nil {
			panic(err)
		}
		PrintAnalyse(r, fmt.Sprintf("general prediction with get more %d balls proba and %d joker proba to random selection (seed: %d)", nBalls, nJoker, seed))
	}()

	return
}

// playGeneralPrediction predict the top 5 probability to find
func playGeneralPrediction(l loto.Loto) (loto.Reward, error) {
	var err error
	var gPredict loto.GeneralPredict
	var scores []loto.Score

	r := loto.Reward{}
	hist := l.History()

	for i := hist.Len() - 100; i >= 0; i-- {
		if gPredict, err = l.GeneralPrediction(loto.Option{
			Index: i,
		}); err != nil {
			return loto.Reward{}, err
		}
		faces, _ := gPredict.Balls.Order()
		faces = faces[:5]
		jokers, _ := gPredict.Joker.Order()
		jokers = jokers[:1]

		draw := loto.Draw{}
		if i-1 < 0 {
			fmt.Printf("faces %v && joker %d\n", faces, jokers[0])
			continue
		}
		if draw, err = hist.DrawByIndex(i - 1); err != nil {
			return loto.Reward{}, err
		}

		s := loto.NewScore(draw, faces, jokers[0])
		scores = append(scores, s)
	}

	for _, s := range scores {
		r.AddScore(s.Value)
	}

	return r, nil
}

// playStaticSolution to test with the same balls and joker on the all history
func playStaticSolution(l loto.Loto, balls []int32, joker int32) (loto.Reward, error) {
	var err error
	var scores []loto.Score

	r := loto.Reward{}
	hist := l.History()

	for i := hist.Len() - 100; i >= 0; i-- {
		draw := loto.Draw{}
		if i-1 < 0 {
			continue
		}
		if draw, err = hist.DrawByIndex(i - 1); err != nil {
			return loto.Reward{}, err
		}
		s := loto.NewScore(draw, balls, joker)
		scores = append(scores, s)
	}
	for _, s := range scores {
		r.AddScore(s.Value)
	}

	return r, nil
}

// playGeneralPredictionWithRandomSelection predict the top n probability est select them randomly with the specific seed
func playGeneralPredictionWithRandomSelection(l loto.Loto, nBalls, nJokers, seed int64) (loto.Reward, error) {
	var err error
	var gPredict loto.GeneralPredict
	var scores []loto.Score

	r := loto.Reward{}
	hist := l.History()

	for i := hist.Len() - 100; i >= 0; i-- {
		if gPredict, err = l.GeneralPrediction(loto.Option{
			Index: i,
		}); err != nil {
			return loto.Reward{}, err
		}
		faces, _ := gPredict.Balls.Order()
		faces = faces[:nBalls]
		jokers, _ := gPredict.Joker.Order()
		jokers = jokers[:nJokers]
		faces = Shuffle(faces, int64(seed))[:5]
		jokers = Shuffle(jokers, int64(seed))[:1]

		draw := loto.Draw{}
		if i-1 < 0 {
			fmt.Printf("faces %v && joker %d\n", faces, jokers[0])
			continue
		}
		if draw, err = hist.DrawByIndex(i - 1); err != nil {
			return loto.Reward{}, err
		}
		s := loto.NewScore(draw, faces, jokers[0])
		scores = append(scores, s)
	}
	for _, s := range scores {
		r.AddScore(s.Value)
	}

	return r, nil
}

// PrintAnalyse the way to find the great number
func PrintAnalyse(r loto.Reward, wayDescription string) {
	fmt.Printf("\n-----\n%s\n\n%s\n", wayDescription, r.String())
	cost, win := r.Price()
	fmt.Printf("total cost %f to total win %f (€) to total benefit: %f\n", cost, win, win-cost)
}

// Shuffle randomise the vals parameter with the salt random
func Shuffle(vals []int32, seed int64) []int32 {
	r := rand.New(rand.NewSource(seed))
	ret := make([]int32, len(vals))
	perm := r.Perm(len(vals))
	for i, randIndex := range perm {
		ret[i] = vals[randIndex]
	}
	return ret
}
